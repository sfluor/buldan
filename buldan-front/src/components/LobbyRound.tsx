import { useState } from "react";
import Input from "./Input";
import { RoundState } from "./Lobby";
import Button from "./Button";
import PlayerBoxes from "./PlayerBoxes";
import GuessBoxes from "./GuessBoxes";

export default function LobbyRound({ user, round, sendGuess }: { user: string, round: RoundState | null, sendGuess: (guess: string) => void }) {

    // Must to that before the early return.
    const [guess, setGuess] = useState(round === null ? "" : round.Letter);

    // TODO: handle failure with notifs
    if (round === null) {
        return <> Unexpected round not being initialized </>
    }

    const { Players, Guesses, Remaining, CurrentPlayerRemainingGuesses, CurrentPlayerIndex, Letter } = round;

    const currentPlayer = Players[CurrentPlayerIndex].Name;
    const isPlaying = currentPlayer === user;

    return <>
        <div>Countries starting with <b className="text-2xl capitalize">{Letter}</b>, {Remaining} remaining to guess, {CurrentPlayerRemainingGuesses} attempts remaining </div>

        {isPlaying && <>
            <Input value={guess} onChange={e => setGuess(e.target.value)} />
            <Button className="m-4" onClick={() => sendGuess(guess)}>Guess !</Button>
        </>}

        <div className="grid grid-cols-2">
            <PlayerBoxes players={Players} current={currentPlayer} user={user} />
            <GuessBoxes guesses={Guesses} />
        </div>
    </>
}
