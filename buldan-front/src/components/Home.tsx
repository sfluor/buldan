import { useState } from "react"
import { primaryColorTxt, secondaryColorTxt } from "./constants"
import Input from "./Input"
import Button from "./Button";



export default function Home() {

    const [user, setUser] = useState("hello");

    return <div className="flex flex-col items-center gap-y-16">
        <div className="text-5xl subpixel-antialiased font-semibold">
            <span className={secondaryColorTxt}>Bul</span>
            <span className={primaryColorTxt}>dan</span>
        </div>
        <div className="text-2xl items-center">
            <span className={primaryColorTxt}>Welcome to</span>
            <br />
            <br />
            <Input value={user} onChange={e => setUser(e.target.value)} />
        </div>
        <div className="text-2xl">
            <span className={secondaryColorTxt}>Start </span>
            <span className={primaryColorTxt}>guessing !</span>
        </div>
        <div className="flex flex-row self-center gap-x-12">
            <Button>New game</Button>
            <Button secondary>Join game</Button>
        </div>
    </div>
}
