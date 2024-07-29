import { useEffect, useRef, useState } from "react";
import { LOBBY_URL } from "./constants";
import Button from "./Button";
import ErrrorToast from "./ErrorToast";



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
    const ws = useRef<WebSocket | null>(null);
    const [error, setError] = useState<string | null>(null);

    useEffect(
        () => {
            console.log("Creating new websocket to: ", url, ws);
            // Technically this is only ran once but in debug mode it can be run
            // twice that's why we check that it's not already initialized.
            if (ws.current == null || (ws.current.readyState != WebSocket.OPEN && ws.current.readyState != WebSocket.CONNECTING)) {
                ws.current = new WebSocket(url);
            }
            ws.current.onopen = (e) => {
                console.log("Opened WebSocket connection to: ", url);
                (e.target as WebSocket).send(JSON.stringify({ "hello": "hi" }));
            };

            ws.current.onmessage = function(event) {
                setMessages(msgs => {
                    const newMessages = [...msgs, event.data];
                    return newMessages;
                });
            };

            ws.current.onerror = function(event) {
                setError("An error occurred while connecting to the websocket !");
                setTimeout(() => setError(null), 5000);
                console.error(event);
            }


            return () => {
                // Cleanup web socket on unmount;
                console.log("ws.current", ws);
                if (ws.current) {
                    ws.current.close(1000, "Going away");
                }
            }
        }
        , [])

    console.log("messages", messages);

    return <div>
        Hello from lobby: {id}
        {messages.map((msg, idx) => <div key={idx}>{msg}</div>)}

        <Button onClick={() => alert("Not implemented")}>Start game !</Button>
        <Button secondary onClick={() => copyToClipboard(`${window.location.host}/lobby/${id}`)}>Share lobby !</Button>
        <ErrrorToast error={error} />
    </div>
}
