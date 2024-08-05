import {
  primaryOutlineLight,
  primaryColorLight,
  secondaryOutlineLight,
  secondaryColorLight,
  errorColorLight,
  errorOutlineLight,
  inactiveColorLight,
  inactiveOutlineLight,
} from "./constants";
import { Player } from "./Lobby";

export default function PlayerBox({
  isUser,
  isPlaying,
  isFirst,
  hasLost,
  player: { Name, Admin, Points, Connected },
}: {
  isUser: boolean;
  isPlaying: boolean;
  isFirst: boolean;
  hasLost: boolean;
  player: Player;
}) {
  let prefix = "";
  const color = isUser
    ? `${primaryColorLight} ${primaryOutlineLight}`
    : `${secondaryColorLight} ${secondaryOutlineLight}`;
  let className = color;
  if (!Connected) {
    prefix = "📶 ";
    className = `text-white ${inactiveColorLight} ${inactiveOutlineLight}`;
  } else if (isPlaying) {
    prefix = "⏳ ";
    className = `${color} font-semibold`;
  } else if (hasLost) {
    prefix = "❌ ";
    className = `${errorColorLight} ${errorOutlineLight} line-through`;
  }

  if (Admin) {
    prefix = `👑 ${prefix}`;
  }

  className = `p-2 my-6 outline outline-offset-2 ${
    isPlaying ? "animate-bounce outline-4" : "outline-1"
  } rounded max-w-64 ${className}`;

  if (isUser) {
    Name = Name + " (you)";
  }

  return (
    <div className={className}>
      <b>{`${prefix}${Name}`}</b>
      {" - "}
      <span
        className={isFirst ? "underline decoration-4 decoration-green-800" : ""}
      >
        ({Points}
        {" points)"}
      </span>
    </div>
  );
}
