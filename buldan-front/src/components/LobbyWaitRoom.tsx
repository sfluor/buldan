import Button from "./Button";
import { Player } from "./Lobby";
import PlayerBoxes from "./PlayerBoxes";


function copyToClipboard(data: string) {
    navigator.clipboard.writeText(data).then(
        () => console.log("Successfully copied data to clipboard"),
        err => {
            console.error("Failed to copy to clipboard", err);
            alert("Copy to clipboard failed");
        }
    )
}



export default function LobbyWaitRoom({ players, user, shareUrl, startGame }: { players: Player[], user: string, shareUrl: string, startGame: () => void }) {
    const isAdmin = players.find(player => player.Name === user)?.Admin;

    // No current player since we are in the waiting room.
    const currentPlayer = "";
    return <>
        <PlayerBoxes players={players} current={currentPlayer} user={user} />

        {isAdmin && <Button onClick={startGame}>Start game !</Button>}
        <Button secondary onClick={() => copyToClipboard(shareUrl)}>Share lobby !</Button>
    </>
}
