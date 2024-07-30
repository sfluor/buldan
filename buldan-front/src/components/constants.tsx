export const primaryColor = "bg-blue-500";
export const secondaryColor = "bg-orange-500";
export const primaryColorTxt = "text-blue-500";
export const secondaryColorTxt = "text-orange-500";

// TODO
const env = process.env.NODE_ENV;
export const HOST = env === "development" ? "localhost:8080" : window.location.host
export const SERVER_URL = `http://${HOST}/api`;
// TODO: wss
export const LOBBY_URL = `ws://${HOST}/api/lobby`;
