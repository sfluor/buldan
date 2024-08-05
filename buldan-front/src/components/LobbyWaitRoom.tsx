import { useState } from "react";
import Button from "./Button";
import { GameOptions, Player } from "./Lobby";
import PlayerBoxes from "./PlayerBoxes";
import Input from "./Input";
import { mainViewCols } from "./constants";

function copyToClipboard(data: string) {
  navigator.clipboard.writeText(data).then(
    () => console.log("Successfully copied data to clipboard"),
    (err) => {
      console.error("Failed to copy to clipboard", err);
      alert("Copy to clipboard failed");
    },
  );
}

export default function LobbyWaitRoom({
  players,
  user,
  shareUrl,
  startGame,
}: {
  players: Player[];
  user: string;
  shareUrl: string;
  startGame: (options: GameOptions) => void;
}) {
  const isAdmin = players.find((player) => player.Name === user)?.Admin;

  const [rounds, setRounds] = useState(5);
  const [guessTime, setGuessTime] = useState(30);

  // No current player since we are in the waiting room.
  const currentPlayer = "";
  return (
    <>
      <div className={`${mainViewCols} mb-8 items-center`}>
        {isAdmin ? (
          <div>
            <div className="mt-8 mb-4"> Options</div>
            <div className="flex flex-col gap-y-4">
              <Input
                type="number"
                value={rounds}
                min={3}
                max={10}
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
            startGame({ Rounds: rounds, GuessTimeSeconds: guessTime })
          }
        >
          Start game !
        </Button>
      )}
      <Button secondary onClick={() => copyToClipboard(shareUrl)}>
        Share lobby !
      </Button>
    </>
  );
}
