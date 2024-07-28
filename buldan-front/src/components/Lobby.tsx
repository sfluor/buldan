import { LOBBY_URL } from "./constants";

export default function Lobby({ id }: { id: string }) {
    const url = `${LOBBY_URL}/${id}`;

    console.log("Creating new websocket to: ", url);
    const ws = new WebSocket(url);

    ws.onopen = () => {
        console.log("Opened WebSocket connection to: ", url)
        ws.send(JSON.stringify({ "hello": "hi" }));
    };

    ws.onmessage = function(event) {
        const json = JSON.parse(event.data);
        console.log("On message", json);
    };

    return <div>
        Hello from lobby: {id}
    </div>
}
