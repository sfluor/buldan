import { Guess } from "./Lobby";



export default function GuessBox({ guess: { Player, Guess, Correct, Flag } }: { guess: Guess }) {
    let prefix = '❌ ';
    if (Correct) {
        prefix = `${Flag} `;
    }

    return <div className="my-4 capitalize text-2xl">
        {prefix}{' '}
        <i className={Correct ? "text-green-700" : "text-red-700"}>{Guess}{' '}</i>
        from <b>{Player}</b>
    </div>
}
