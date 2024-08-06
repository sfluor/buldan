import { EndRound } from "./Lobby";

export default function LobbyRoundEnd({
  endRound,
  remainingSec,
}: {
  endRound: EndRound;
  remainingSec: number | null;
}) {
  return (
    <div>
      End of round ! {remainingSec} seconds before next round...
      {endRound.Countries.map(({ Name, Flag, GuessedBy }, idx) => (
        <div
          className={GuessedBy ? "text-green-300" : "text-red-300"}
          key={idx}
        >{`${Flag} ${Name} (${GuessedBy})`}</div>
      ))}
    </div>
  );
}
