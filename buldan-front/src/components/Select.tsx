import { useId } from "react";
import { primaryColor } from "./constants";



const blockClass = "min-w-36 min-h-12 font-bold py-2 px-4 border-l-2 border-t-2 border-b-4 border-r-2 text-white rounded transition duration-500";

const primaryClass = `${primaryColor} hover:bg-blue-300 border-blue-700 hover:border-blue-500`;

const optionClass = `${primaryColor} text-white min-w-36 min-h-12 p-2`

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
    // TODO: quite ugly
    return (
        <label htmlFor={id}>
            {label}
            <select id={id} value={value} onChange={onChange} className={`${blockClass} ${primaryClass}`}>
                {choices.map((choice, idx) => (
                    <option key={idx} value={choice} className={optionClass}>
                        {choice}
                    </option>
                ))}
            </select>
        </label>

    );
}
