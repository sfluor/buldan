import { useId } from "react";

export default function Select({
  choices,
  value,
  onChange,
  label,
}: {
  choices: string[];
  value?: string;
  label: string;
  onChange?: (e: React.ChangeEvent<HTMLSelectElement>) => void;
}) {
  const id = useId();
  return (
    <label htmlFor={id}>
      {label}
      <select id={id} value={value} onChange={onChange}>
        {choices.map((choice, idx) => (
          <option key={idx} value={choice}>
            {choice}
          </option>
        ))}
      </select>
    </label>
  );
}
