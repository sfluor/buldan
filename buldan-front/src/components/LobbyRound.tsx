import {  RoundState } from "./Lobby";

export default function LobbyRound({ round }: { round: RoundState }) {

// TODO: handle failure with notifs
    if (round === null) {
        return <> Unexpected round not being initialized </>
    }

    const {Players, Guesses, Remaining, Letter} = round;

    return <>

        <div>Guessing for <b>{Letter}</b>, {Remaining} remaining to guess </div>

        <div>Players:</div>
        {Players.map(({ Name }) => <div key={Name}>{Name}</div>)}

        <div> Guesses</div>
        {Guesses.map(({ Guess, Player, Correct }) => <div key={Guess + Player}>{Guess} from {Player} {Correct ? "correct" : "wrong"}</div>)}
    </>
}
