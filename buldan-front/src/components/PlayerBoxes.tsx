import { Player, PlayerStatus } from "./Lobby";
import PlayerBox from "./PlayerBox";

export default function PlayerBoxes({
  players,
  current,
  playersStatuses,
  user,
}: {
  players: Player[];
  current?: string;
  playersStatuses?: Record<string, PlayerStatus>;
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
          hasLost={
            playersStatuses
              ? playersStatuses[player.Name].RemainingGuesses <= 0
              : false
          }
          isUser={player.Name === user}
        />
      ))}
    </div>
  );
}
