import { mainViewCols, strongGreenTxt, strongRedTxt } from "./constants";
import GuessBoxes from "./GuessBoxes";
import { Country, EndRound } from "./Lobby";

function CountryLine({
  country: { Name, Flag, GuessedBy },
}: {
  country: Country;
}) {
  const color = GuessedBy ? strongGreenTxt : strongRedTxt;
  const className = `${color} my-4 capitalize text-2xl`;

  let suffix = "";
  if (GuessedBy) {
    suffix = ` (${GuessedBy})`;
  }

  return <div className={className}>{`${Flag} ${Name}${suffix}`}</div>;
}

export default function LobbyRoundEnd({
  endRound,
  remainingSec,
}: {
  endRound: EndRound;
  remainingSec: number | null;
}) {
  return (
    <div>
      End of round ! {remainingSec} seconds before{" "}
      {endRound.IsLastRound ? "leaderboard" : "next round"}...
      <div className={mainViewCols}>
        <div>
          <div className="mt-8"> Countries</div>
          {endRound.Countries.map((country, idx) => (
            <CountryLine key={idx} country={country} />
          ))}
        </div>
        <GuessBoxes guesses={endRound.Guesses} />
      </div>
    </div>
  );
}
