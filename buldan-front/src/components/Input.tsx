export default function Input({
  value,
  onChange,
}: {
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
}) {
  return (
    <input
      type="text"
      value={value}
      onChange={onChange}
      className="focus:border-indigo-500 focus:text-blue-500 outline-none box border border-solid border-blue-500 border-b-8 px-4 py-2 text-indigo-500 font-semibold border-r-8 transition duration-1000 max-w-80"
    />
  );
}
