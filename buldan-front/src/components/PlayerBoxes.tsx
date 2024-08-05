import { Player } from "./Lobby";
import PlayerBox from "./PlayerBox";

export default function PlayerBoxes({
  players,
  current,
  playersOut,
  user,
}: {
  players: Player[];
  current: string;
  playersOut?: Record<string, boolean>;
  user: string;
}) {
  const bestScore = Math.max(...players.map(({ Points }) => Points));
  return (
    <div>
      <div className="mt-8"> Players</div>
      {players.map((player, idx) => (
        <PlayerBox
          player={player}
          key={idx}
          isFirst={bestScore === player.Points && bestScore > 0}
          isPlaying={player.Name === current}
          hasLost={playersOut ? playersOut[player.Name] : false}
          isUser={player.Name === user}
        />
      ))}
    </div>
  );
}
