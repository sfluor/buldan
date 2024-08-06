import { strongGreenTxt, strongRedTxt } from "./constants";
import { Guess } from "./Lobby";

export default function GuessBox({
  guess: { Player, Guess, Correct, Flag },
}: {
  guess: Guess;
}) {
  let prefix = "❌ ";
  if (Correct) {
    prefix = `${Flag} `;
  }

  return (
    <div className="my-4 capitalize text-2xl">
      {prefix}{" "}
      <i className={Correct ? strongGreenTxt : strongRedTxt}>{Guess} </i>
      from <b>{Player}</b>
    </div>
  );
}
