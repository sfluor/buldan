import { useState } from "react";
import { LOBBY_URL } from "./constants";
import Button from "./Button";


let ws: WebSocket | undefined;

function copyToClipboard(data: string) {
    navigator.clipboard.writeText(data).then(
        () => console.log("Successfully copied data to clipboard"),
        err => {
            console.error("Failed to copy to clipboard", err);
            alert("Copy to clipboard failed");
        }
    )
}


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
            event.target.setMessages(msgs => {
                const newMessages = [...msgs, event.data];
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

        <Button onClick={() => alert("Not implemented")}>Start game !</Button>
        <Button secondary onClick={() => copyToClipboard(`${window.location.host}/lobby/${id}`)}>Share lobby !</Button>
    </div>
}
