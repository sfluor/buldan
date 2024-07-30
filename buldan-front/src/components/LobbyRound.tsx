import { Guess, Player } from "./Lobby";

export default function LobbyRound({ players, guesses, letter, remaining }: { players: Player[], guesses: Guess[], letter: string, remaining: number }) {
    return <>

        <div>Guessing for <b>{letter}</b>, {remaining} remaining to guess </div>

        <div>Players:</div>
        {players.map(({ Name }) => <div key={Name}>{Name}</div>)}

        <div> Guesses</div>
        {guesses.map(({ Guess, Player, Correct }) => <div key={Guess + Player}>{Guess} from {Player} {Correct ? "correct" : "wrong"}</div>)}
    </>
}
