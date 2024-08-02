import { useState } from "react";
import Input from "./Input";
import { RoundState } from "./Lobby";
import Button from "./Button";

export default function LobbyRound({ user, round, sendGuess }: { user: string, round: RoundState | null, sendGuess: (guess: string) => void }) {

    // Must to that before the early return.
    const [guess, setGuess] = useState(round === null ? "" : round.Letter);

    // TODO: handle failure with notifs
    if (round === null) {
        return <> Unexpected round not being initialized </>
    }

    const { Players, Guesses, Remaining, CurrentPlayerIndex, Letter } = round;

    const isPlaying = Players[CurrentPlayerIndex].Name === user;

    return <>
        <div>Guessing for <b>{Letter}</b>, {Remaining} remaining to guess </div>

        {isPlaying && <>
            <Input value={guess} onChange={e => setGuess(e.target.value)} />
            <Button onClick={() => sendGuess(guess)}>Guess !</Button>
        </>}

        <br />

        <div>Players:</div>
        {Players.map(({ Name }) => <div key={Name}>{Name}</div>)}

        <div> Guesses</div>
        {Guesses.map(({ Guess, Player, Correct }, idx) => <div key={Guess + Player + idx}>{Guess} from {Player} {Correct ? "correct" : "wrong"}</div>)}
    </>
}
