import { Player } from "./Lobby";
import PlayerBox from "./PlayerBox";

export default function PlayerBoxes({
  players,
  current,
  user,
}: {
  players: Player[];
  current: string;
  user: string;
}) {
  return (
    <div>
      <div className="mt-8"> Players</div>
      {players.map((player, idx) => (
        <PlayerBox
          player={player}
          key={idx}
          isPlaying={player.Name === current}
          isUser={player.Name === user}
        />
      ))}
    </div>
  );
}
