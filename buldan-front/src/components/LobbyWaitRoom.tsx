import Button from "./Button";
import { primaryColorTxt, secondaryColorTxt } from "./constants";
import { Player } from "./Lobby";


function copyToClipboard(data: string) {
    navigator.clipboard.writeText(data).then(
        () => console.log("Successfully copied data to clipboard"),
        err => {
            console.error("Failed to copy to clipboard", err);
            alert("Copy to clipboard failed");
        }
    )
}



export default function LobbyWaitRoom({ players, user, id, startGame }: { players: Player[], user: string, id: string, startGame: () => void }) {
    const isAdmin = players.find(player => player.Name === user)?.Admin;

    return <>
        {players.map(({ Name, Admin }, idx) => <div key={idx} className={Name === user ? `${primaryColorTxt} font-semibold` : secondaryColorTxt}>{`${Name == user ? "> " : ""}${Name}${Admin ? " (admin)" : ""}`}</div>)}

        {isAdmin && <Button onClick={startGame}>Start game !</Button>}
        <Button secondary onClick={() => copyToClipboard(`${window.location.host}/lobby/${id}`)}>Share lobby !</Button>
    </>
}
