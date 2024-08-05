import { EndRound } from "./Lobby";

export default function LobbyRoundEnd({ endRound }: { endRound: EndRound }) {
  return (
    <div>
      End of round !
      {endRound.Countries.map(({ Name, Flag, GuessedBy }, idx) => (
        <div
          className={GuessedBy ? "text-green-300" : "text-red-300"}
          key={idx}
        >{`${Flag} ${Name} (${GuessedBy})`}</div>
      ))}
    </div>
  );
}
