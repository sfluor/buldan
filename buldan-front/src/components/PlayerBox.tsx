import {
  primaryOutlineLight,
  primaryColorLight,
  secondaryOutlineLight,
  secondaryColorLight,
  errorColorLight,
  errorOutlineLight,
} from "./constants";
import { Player } from "./Lobby";

export default function PlayerBox({
  isUser,
  isPlaying,
  isFirst,
  player: { Name, Admin, Lost, Points },
}: {
  isUser: boolean;
  isPlaying: boolean;
  isFirst: boolean;
  player: Player;
}) {
  let prefix = "";
  const color = isUser
    ? `${primaryColorLight} ${primaryOutlineLight}`
    : `${secondaryColorLight} ${secondaryOutlineLight}`;
  let className = color;
  if (isPlaying) {
    prefix = "⏳ ";
    className = `${color} font-semibold`;
  } else if (Lost) {
    prefix = "❌ ";
    className = `${errorColorLight} ${errorOutlineLight} line-through`;
  }

  if (Admin) {
    prefix = `👑 ${prefix}`;
  }

  className = `p-2 my-6 outline outline-offset-2 ${
    isPlaying ? "animate-bounce outline-4" : "outline-1"
  } rounded max-w-64 ${className}`;

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
