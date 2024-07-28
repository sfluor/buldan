import { useState } from "react";
import { LOBBY_URL } from "./constants";


let ws: WebSocket | undefined;


export default function Lobby({ id, user }: { id: string, user: string }) {
    const url = `${LOBBY_URL}/${id}/${user}`;

    const [messages, setMessages] = useState<string[]>([]);

    if (ws === undefined || ws.readyState === WebSocket.CLOSING || ws.readyState === WebSocket.CLOSED) {
        console.log("Recreating new websocket to: ", url, ws);
        ws = new WebSocket(url);

        ws.onopen = (e) => {
            console.log("Opened WebSocket connection to: ", url);
            e.target.send(JSON.stringify({ "hello": "hi" }));
        };

        ws.onmessage = function(event) {
            console.log("On message", event.data);
            event.target.setMessages(msgs => {
                const newMessages = [...msgs, event.data];
                console.log("callback", newMessages);
                return newMessages;
            });
        };
    } else {
        ws.setMessages = setMessages;
    }

    console.log("messages", messages);

    return <div>
        Hello from lobby: {id}
        {messages.map((msg, idx) => <div key={idx}>{msg}</div>)}
    </div>
}
