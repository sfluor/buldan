import { useEffect, useRef, useState } from "react";
import { LOBBY_URL, SERVER_URL } from "./constants";
import Toast from "./Toast";
import { useLocation } from "wouter";
import LobbyWaitRoom from "./LobbyWaitRoom";
import axios from "axios";
import LobbyRound from "./LobbyRound";
import LobbyEndGame from "./LobbyEndGame";
import BuldanText from "./BuldanText";
import LobbyRoundEnd from "./LobbyRoundEnd";

export enum Language {
  French = "french",
  English = "english",
}

export interface Player {
  Name: string;
  Admin: boolean;
  Points: number;
  Connected: boolean;
}

export interface PlayerStatus {
  RemainingGuesses: number;
}

interface Notification {
  message: string;
  error?: boolean;
}

export interface Guess {
  Guess: string;
  Player: string;
  Correct: boolean;
  Flag?: string;
}

export interface GameOptions {
  Language: Language;
  Rounds: number;
  GuessTimeSeconds: number;
  MaxGuessesPerRound: number;
}

export interface RoundState {
  Round: number;
  MaxRounds: number;
  Guesses: Guess[];
  Letter: string;
  PlayersStatuses: Record<string, PlayerStatus>;
  Remaining: number;
  CurrentPlayerIndex: number;
}

export interface Country {
  Name: string;
  Flag: string;
  GuessedBy?: string;
}

export interface EndRound {
  IsLastRound: boolean;
  Round: number;
  MaxRounds: number;
  Letter: string;
  Guesses: Guess[];
  Countries: Country[];
}

export interface EndGame {
  Players: Player[];
}

export const createLobby = async (
  user: string,
  setLocation: (url: string) => void,
) => {
  try {
    const res = await axios.post(`${SERVER_URL}/new-lobby`);
    setLocation(`/lobby/${res.data.id}/${user}`);
  } catch (error) {
    console.error(error);
    alert(`An error occurred ${error}`);
  }
};

export default function Lobby({ id, user }: { id: string; user: string }) {
  // TODO: dedupe players and round
  const [players, setPlayers] = useState<Player[]>([]);
  const ws = useRef<WebSocket | null>(null);
  const [notif, setNotif] = useState<Notification | null>(null);
  const [round, setRound] = useState<RoundState | null>(null);
  const [endRound, setEndRound] = useState<EndRound | null>(null);
  const [endGame, setEndGame] = useState<EndGame | null>(null);
  const [remainingSec, setRemainingSec] = useState<number | null>(null);

  // (otherwise not happy about the any)
  // eslint-disable-next-line
  const send = (content: any) => {
    if (ws.current) {
      ws.current.send(JSON.stringify(content));
    } else {
      // TODO
      console.error("web socket isn't initialized yet");
    }
  };

  const startGame = (options: GameOptions) =>
    send({ Type: "start-game", Options: options });

  const sendGuess = (guess: string) => send({ Type: "guess", Guess: guess });

  const url = `${LOBBY_URL}/${id}/${user}`;
  const shareUrl = `${window.location.protocol}//${window.location.host}/lobby/${id}`;
  const setLocation = useLocation()[1];

  useEffect(() => {
    // Technically this is only ran once but in debug mode it can be run
    // twice that's why we check that it's not already initialized.
    if (
      ws.current == null ||
      (ws.current.readyState != WebSocket.OPEN &&
        ws.current.readyState != WebSocket.CONNECTING)
    ) {
      ws.current = new WebSocket(url);
      console.log("Creating new websocket to: ", url, ws);
    }

    ws.current.onopen = () => {
      console.log("Opened WebSocket connection to: ", url);
    };

    return () => {
      // Cleanup web socket on unmount;
      if (ws.current) {
        console.log("Closing websocket connection", ws);
        ws.current.close(1000, "Going away");
      }
    };
  }, [url]);

  const notify = (notif: Notification, redirect: boolean) => {
    setNotif(notif);
    setTimeout(() => {
      setNotif(null);
      if (redirect) {
        setLocation("/");
      }
    }, 5000);
  };

  useEffect(() => {
    if (ws.current === null) {
      return;
    }

    ws.current.onmessage = function (event) {
      const json = JSON.parse(event.data);
      if (json.Type === "tick") {
        setRemainingSec(json.RemainingSec);
      } else if (json.Type === "players") {
        setPlayers([...json.Players]);
      } else if (json.Type === "new-round" || json.Type === "round-update") {
        setRemainingSec(null);
        setEndRound(null);
        setRound(json.Round);
        setPlayers(json.Players);
      } else if (json.Type === "end-round") {
        setRemainingSec(null);
        setRound(null);
        setEndRound(json);
      } else if (json.Type === "end") {
        setRemainingSec(null);
        setEndRound(null);
        setRound(null);
        setEndGame(json);
      } else {
        notify(
          {
            message: `Unknown payload received ${json.Type}: ${event.data}`,
            error: true,
          },
          true,
        );
      }
    };

    ws.current.onclose = function (event) {
      console.warn("Closed websocket", event);
      if (notif === null) {
        notify(
          {
            message: `An error occurred: ${event.reason}. Redirecting back home...`,
            error: true,
          },
          true,
        );
      }
    };
  }, [notify, setLocation, setRound, setPlayers, setRemainingSec]);

  let component;
  let currentRound;
  let maxRounds;
  if (endGame != null) {
    component = <LobbyEndGame user={user} players={endGame.Players} />;
  } else if (round === null && endRound === null) {
    component = (
      <LobbyWaitRoom
        players={players}
        user={user}
        shareUrl={shareUrl}
        onCopyLink={() =>
          notify(
            {
              message: "Successfully copied lobby link to the clipboard !",
            },
            false,
          )
        }
        startGame={startGame}
      />
    );
  } else if (round != null) {
    component = (
      <LobbyRound
        user={user}
        round={round}
        players={players}
        sendGuess={sendGuess}
        remainingSec={remainingSec}
      />
    );
    currentRound = round.Round;
    maxRounds = round.MaxRounds;
  } else if (endRound != null) {
    component = (
      <LobbyRoundEnd remainingSec={remainingSec} endRound={endRound} />
    );
    currentRound = endRound.Round;
    maxRounds = endRound.MaxRounds;
  } else {
    component = <div> Unexpected lobby state ! </div>;
  }

  return (
    <div className="p-4">
      <BuldanText />
      <LobbyHeader
        shareUrl={shareUrl}
        id={id}
        currentRound={currentRound}
        maxRounds={maxRounds}
      />

      {component}

      {notif && <Toast message={notif.message} error={notif.error} />}
    </div>
  );
}

function LobbyHeader({
  id,
  shareUrl,
  currentRound,
  maxRounds,
}: {
  id: string;
  shareUrl: string;
  currentRound?: number;
  maxRounds?: number;
}) {
  return (
    <div className="p-4 my-4 text-lg bg-indigo-100 rounded-md">
      {" "}
      <a
        className="font-medium text-blue-600 dark:text-blue-500 hover:underline"
        href={shareUrl}
        target="_blank"
      >
        {" "}
        Lobby {id}
        {maxRounds ? ` | Round ${currentRound}/${maxRounds}` : ""}
      </a>
    </div>
  );
}
