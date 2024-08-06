import { useState } from "react";
import Input from "./Input";
import { Player, RoundState } from "./Lobby";
import Button from "./Button";
import PlayerBoxes from "./PlayerBoxes";
import GuessBoxes from "./GuessBoxes";
import { mainViewCols } from "./constants";

export default function LobbyRound({
  user,
  round,
  players,
  sendGuess,
}: {
  user: string;
  round: RoundState | null;
  players: Player[];
  sendGuess: (guess: string) => void;
}) {
  // Must to that before the early return.
  const [guess, setGuess] = useState(round === null ? "" : round.Letter);

  // TODO: handle failure with notifs
  if (round === null) {
    return <> Unexpected round not being initialized </>;
  }

  const { Guesses, Remaining, CurrentPlayerIndex, Letter } = round;

  const currentPlayer = players[CurrentPlayerIndex].Name;
  const isPlaying = currentPlayer === user;

  const remainingGuesses =
    round.PlayersStatuses[currentPlayer].RemainingGuesses;

  return (
    <>
      <div>
        Countries starting with <b className="text-2xl capitalize">{Letter}</b>,{" "}
        {Remaining} remaining to guess, {remainingGuesses} attempts remaining{" "}
      </div>

      {isPlaying && (
        <>
          <Input value={guess} onChange={(e) => setGuess(e.target.value)} />
          <Button className="m-4" onClick={() => sendGuess(guess)}>
            Guess !
          </Button>
        </>
      )}

      <div className={mainViewCols}>
        <PlayerBoxes
          players={players}
          playersStatuses={round.PlayersStatuses}
          current={currentPlayer}
          user={user}
        />
        <GuessBoxes guesses={Guesses} />
      </div>
    </>
  );
}
