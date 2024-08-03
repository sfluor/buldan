export const primaryColor = "bg-blue-500";
export const secondaryColor = "bg-orange-500";

export const primaryColorLight = "bg-blue-200";
export const secondaryColorLight = "bg-orange-200";
export const errorColorLight = "bg-red-200";

export const primaryOutlineLight = "outline-blue-700";
export const secondaryOutlineLight = "outline-orange-700";
export const errorOutlineLight = "outline-red-900";

export const primaryColorTxt = "text-blue-500";
export const secondaryColorTxt = "text-orange-500";

export const errorColorTxt = "text-red-500";

// TODO
const env = process.env.NODE_ENV;
export const HOST =
  env === "development" ? "localhost:8080" : window.location.host;
export const SERVER_URL = `http://${HOST}/api`;
// TODO: wss
export const LOBBY_URL = `ws://${HOST}/api/lobby`;
