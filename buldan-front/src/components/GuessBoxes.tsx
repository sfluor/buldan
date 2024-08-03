import GuessBox from "./GuessBox";
import { Guess } from "./Lobby";

export default function GuessBoxes({ guesses }: { guesses: Guess[] }) {
  const reversed: Guess[] = guesses.toReversed();

  return (
    <div>
      <div className="mt-8"> Guesses</div>
      {reversed.map((guess, idx) => (
        <GuessBox key={idx} guess={guess} />
      ))}
    </div>
  );
}
