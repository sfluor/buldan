import { primaryOutlineLight, primaryColorLight, secondaryOutlineLight, secondaryColorLight, errorColorLight, errorOutlineLight } from "./constants";
import { Player } from "./Lobby";

export default function PlayerBox({ isUser, isPlaying, player: { Name, Admin, Lost } }: { isUser: boolean, isPlaying: boolean, player: Player }) {

    let prefix = "";
    const color = isUser ? `${primaryColorLight} ${primaryOutlineLight}` : `${secondaryColorLight} ${secondaryOutlineLight}`;
    let className = color;
    if (isPlaying) {
        prefix = "⏳ ";
        className = `${color} font-semibold`;
    } else if (Lost) {
        prefix = '❌ ';
        className = `${errorColorLight} ${errorOutlineLight} line-through`;
    }

    if (Admin) {
        prefix = `👑 ${prefix}`;
    }

    className = `p-2 m-4 outline outline-offset-2 rounded max-w-64 ${className}`

    return <div className={className}>{`${prefix}${Name}`}</div>
}
