import { useState } from "react";
import Input from "./Input";
import { RoundState } from "./Lobby";
import Button from "./Button";
import PlayerBoxes from "./PlayerBoxes";

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
        <div>Guessing for <b>{Letter}</b>, {Remaining} remaining to guess, {CurrentPlayerRemainingGuesses} attempts remaining </div>

        {isPlaying && <>
            <Input value={guess} onChange={e => setGuess(e.target.value)} />
            <Button onClick={() => sendGuess(guess)}>Guess !</Button>
        </>}
        <div className="mt-8">Players:</div>
        <PlayerBoxes players={Players} current={currentPlayer} user={user} />

        <div className="mt-8"> Guesses</div>
        {Guesses.map(({ Guess, Player, Correct, Flag }, idx) => <div key={Guess + Player + idx}><i className={`capitalize text-2xl ${Correct ? "text-green-700" : "text-red-700"}`}>{Flag && `${Flag} `}{Guess}</i> from <b>{Player}</b></div>)}
    </>
}
