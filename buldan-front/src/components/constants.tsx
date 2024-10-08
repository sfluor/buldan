export const primaryColor = "bg-blue-500";
export const secondaryColor = "bg-indigo-500";

export const primaryColorLight = "bg-blue-200";
export const secondaryColorLight = "bg-indigo-200";
export const inactiveColorLight = "bg-slate-500";
export const errorColorLight = "bg-red-200";

export const primaryOutlineLight = "outline-blue-700";
export const secondaryOutlineLight = "outline-indigo-700";
export const inactiveOutlineLight = "outline-slate-700";
export const errorOutlineLight = "outline-red-900";

export const primaryColorTxt = "text-blue-500";
export const secondaryColorTxt = "text-indigo-500";

export const errorColorTxt = "text-red-500";

export const strongGreenTxt = "text-green-700";
export const strongRedTxt = "text-red-700";

export const mainViewCols = "grid md:grid-cols-2 grid-cols-1";

const env = process.env.NODE_ENV;
export const HOST =
  env === "development" ? "localhost:8080" : window.location.host;


export const VERSION = import.meta.env.VITE_BULDAN_VERSION;
const PROTOCOL = window.location.protocol;
export const SERVER_URL = `${PROTOCOL}//${HOST}/api`;

const WS_PROTOCOL = PROTOCOL.replace("http", "ws");
export const LOBBY_URL = `${WS_PROTOCOL}//${HOST}/api/lobby`;
