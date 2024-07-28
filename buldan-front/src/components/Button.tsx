import { primaryColor, secondaryColor} from "./constants"

const buttonBaseClass = "min-w-36 min-h-12 font-bold py-2 px-4 border-b-4 border-r-4 text-white rounded transition duration-500 focus:border-b-0 focus:border-r-0"

const primaryClass = `${primaryColor} hover:bg-blue-300 border-blue-700 hover:border-blue-500`
const secondaryClass = `${secondaryColor} hover:bg-orange-300 border-orange-700 hover:border-orange-500`

const primaryButton = `${buttonBaseClass} ${primaryClass}`
const secondaryButton = `${buttonBaseClass} ${secondaryClass}`

function Button({ children, secondary, onClick }: { children?: React.ReactNode, secondary?: boolean | undefined, onClick?: () => void }) {
    return <button
        onClick={onClick}
        className={secondary ? secondaryButton : primaryButton}
    >
        {children}
    </button>
}

export default Button;
