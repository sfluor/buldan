export default function Toast({
  message,
  error,
}: {
  message: string;
  error?: boolean;
}) {
  return (
    <div
      className={`transition duration-1000 animate-bounce rounded fixed bottom-8 right-20 text-white px-4 py-2 font-semibold border-b-4 border-r-4 max-w-80 text-xl ${
        error ? "bg-red-500 border-red-800" : "bg-teal-500 border-teal-800"
      }`}
    >
      {message}
    </div>
  );
}
