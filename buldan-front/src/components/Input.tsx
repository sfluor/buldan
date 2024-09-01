export default function Input({
  type,
  label,
  value,
  onChange,
  onEnter,
  min,
  max,
}: {
  type?: string;
  label?: string;
  min?: number;
  max?: number;
  value?: string | number;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
  onEnter?: () => void;
}) {
  const valid = (() => {
    if (value === undefined) {
      return false;
    }

    if (type !== "number") {
      return true;
    }

    if (typeof value !== "number") {
      value = parseInt(value);
    }

    if (min !== undefined && value < min) {
      return false;
    }
    if (max !== undefined && value > max) {
      return false;
    }
    return true;
  })();

  let borderClasses =
    "focus:border-indigo-500 focus:text-blue-500 border border-solid border-blue-500 text-indigo-500 font-semibold";
  if (!valid) {
    borderClasses =
      "border border-solid border-red-500 text-red-500 font-semibold";
  }

  const input = (
    <input
      type={type ? type : "text"}
      min={min}
      max={max}
      value={value}
      onChange={onChange}
      onKeyUp={
        onEnter
          ? (event) => {
              if (event.key === "Enter") onEnter();
            }
          : undefined
      }
      className={`${borderClasses} border-r-8 border-b-8 px-4 py-2 outline-none box transition duration-1000 max-w-80`}
    />
  );

  if (label) {
    return (
      <div className="flex flex-col">
        <span
          className={`${borderClasses} font-semibold max-w-28 border-1 border-b-0 px-2`}
        >
          {label}
        </span>
        {input}
      </div>
    );
  }

  return input;
}
