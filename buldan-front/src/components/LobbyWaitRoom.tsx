import { useState } from "react";
import Button from "./Button";
import { GameOptions, Player, Language } from "./Lobby";
import PlayerBoxes from "./PlayerBoxes";
import Input from "./Input";
import { mainViewCols } from "./constants";
import Select from "./Select";

function share(url: string, title: string, onCopyLink: () => void) {
  if (navigator.share) {
    navigator.share({ url, title }).then(
      () => console.log("Successfully shared"),
      (err) => {
        console.error("Failed to share", err);
        alert("Share failed");
      },
    );
  } else if (navigator.clipboard) {
    navigator.clipboard.writeText(url).then(
      () => {
        onCopyLink();
        console.log("Successfully copied data to clipboard");
      },
      (err) => {
        console.error("Failed to copy to clipboard", err);
        alert("Copy to clipboard failed");
      },
    );
  } else {
    alert("Copy share URL: " + url);
  }
}

function toInt(v: string | number) {
  if (typeof v === "string") {
    return parseInt(v);
  }

  return v;
}

export default function LobbyWaitRoom({
  players,
  user,
  shareUrl,
  onCopyLink,
  startGame,
}: {
  players: Player[];
  user: string;
  shareUrl: string;
  onCopyLink: () => void;
  startGame: (options: GameOptions) => void;
}) {
  const isAdmin = players.find((player) => player.Name === user)?.Admin;

  const [rounds, setRounds] = useState<string | number>(5);
  const [guessesPerRound, setGuessesPerRound] = useState<string | number>(3);
  const [guessTime, setGuessTime] = useState<string | number>(30);
  const [lang, setLang] = useState<Language>(Language.English);

  // No current player since we are in the waiting room.
  const currentPlayer = "";
  return (
    <>
      <div className={`${mainViewCols} mb-8 items-center`}>
        {isAdmin ? (
          <div>
            <div className="mt-8 mb-4"> Options</div>
            <div className="flex flex-col gap-y-4">
              <Select
                value={lang}
                onChange={(e) => setLang(e.target.value as Language)}
                choices={Object.values(Language)}
                label="Language"
              />
              <Input
                type="number"
                value={guessesPerRound}
                min={1}
                max={10}
                onChange={(e) => setGuessesPerRound(e.target.value)}
                label="Guesses per round"
              />
              <Input
                type="number"
                value={rounds}
                min={3}
                max={20}
                onChange={(e) => setRounds(e.target.value)}
                label="Max rounds"
              />
              <Input
                type="number"
                min={10}
                max={90}
                value={guessTime}
                onChange={(e) => setGuessTime(e.target.value)}
                label="Guess time (sec)"
              />
            </div>
          </div>
        ) : (
          <div className="text-2xl">
            {" "}
            Waiting for the admin to launch the game...
          </div>
        )}
        <PlayerBoxes players={players} current={currentPlayer} user={user} />
      </div>
      {isAdmin && (
        <Button
          className="mr-4"
          onClick={() =>
            startGame({
              Language: lang,
              Rounds: toInt(rounds),
              GuessTimeSeconds: toInt(guessTime),
              MaxGuessesPerRound: toInt(guessesPerRound),
            })
          }
        >
          Start game !
        </Button>
      )}
      <Button
        secondary
        onClick={() => share(shareUrl, "Join Buldan game !", onCopyLink)}
      >
        Share lobby !
      </Button>
    </>
  );
}
