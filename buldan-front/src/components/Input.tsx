export default function Input({
  type,
  label,
  value,
  onChange,
}: {
  type?: string;
  label?: string;
  value?: string | number;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
}) {
  const input = (
    <input
      type={type ? type : "text"}
      value={value}
      onChange={onChange}
      className="focus:border-indigo-500 focus:text-blue-500 outline-none box border border-solid border-blue-500 border-b-8 px-4 py-2 text-indigo-500 font-semibold border-r-8 transition duration-1000 max-w-80"
    />
  );

  if (label) {
    return (
      <div className="flex flex-col">
        <span className="font-semibold border border-solid max-w-28 border-blue-500 border-1 border-b-0 px-2">
          {label}
        </span>
        {input}
      </div>
    );
  }

  return input;
}
