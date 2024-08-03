import { Player } from "./Lobby";
import PlayerBox from "./PlayerBox";

export default function PlayerBoxes({ players, current, user }: { players: Player[], current: string, user: string }) {
    return <>
        {players.map((player, idx) => <PlayerBox player={player} key={idx} isPlaying={player.Name === current} isUser={player.Name === user} />)}
    </>
}
