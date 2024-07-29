export default function ErrrorToast({ error }: { error: string | null }) {
    if (error === null) {
        return <></>
    }

    return <div
        id="error-toast"
        className="transition duration-1000 animate-bounce rounded fixed bottom-8 right-20 text-white bg-red-500 px-4 py-2 font-semibold border-b-4 border-r-4 border-solide border-red-800 max-w-80 text-xl"
    >
        {error}
    </div>
}
