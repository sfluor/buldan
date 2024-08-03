import { primaryColorTxt, secondaryColorTxt } from "./constants";

export default function BuldanText() {
    return <div className="text-5xl subpixel-antialiased font-semibold">
        <span className={secondaryColorTxt}>Bul</span>
        <span className={primaryColorTxt}>dan</span>
    </div>
}
