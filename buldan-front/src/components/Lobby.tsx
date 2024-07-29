import { useEffect, useRef, useState } from "react";
import { LOBBY_URL, primaryColorTxt, secondaryColorTxt } from "./constants";
import Button from "./Button";
import Toast from "./Toast";
import { useLocation } from "wouter";


function copyToClipboard(data: string) {
    navigator.clipboard.writeText(data).then(
        () => console.log("Successfully copied data to clipboard"),
        err => {
            console.error("Failed to copy to clipboard", err);
            alert("Copy to clipboard failed");
        }
    )
}


interface Player {
    Name: string
    Admin: boolean
}

interface Notification {
    message: string
    error?: boolean
}

export default function Lobby({ id, user }: { id: string, user: string }) {
    const [players, setPlayers] = useState<Player[]>([]);
    const ws = useRef<WebSocket | null>(null);
    const [notif, setNotif] = useState<Notification | null>(null);

    const setLocation = useLocation()[1];
    useEffect(
        () => {
            const url = `${LOBBY_URL}/${id}/${user}`;

            const notifAndRedirect = (notif: Notification) => {
                setNotif(notif);
                setTimeout(() => {
                    setNotif(null);
                    setLocation("/");
                }, 5000);
            }

            // Technically this is only ran once but in debug mode it can be run
            // twice that's why we check that it's not already initialized.
            if (ws.current == null || (ws.current.readyState != WebSocket.OPEN && ws.current.readyState != WebSocket.CONNECTING)) {
                ws.current = new WebSocket(url);
                console.log("Creating new websocket to: ", url, ws);
            }
            ws.current.onopen = (e) => {
                console.log("Opened WebSocket connection to: ", url);
                (e.target as WebSocket).send(JSON.stringify({ "hello": "hi" }));
            };

            ws.current.onmessage = function(event) {
                const json = JSON.parse(event.data);
                if (json.Type === "players") {
                    setPlayers(json.Players);
                } else {
                    notifAndRedirect({ message: `Unknown payload received: ${event.data}`, error: true });
                }
            };

            ws.current.onclose = function(event) {
                console.warn("Closed websocket", event);
                if (notif === null) {
                    notifAndRedirect({ message: `An error occurred: ${event.reason}. Redirecting back home...`, error: true });
                }
            }


            return () => {
                // Cleanup web socket on unmount;
                if (ws.current) {
                    console.log("Closing websocket connection", ws);
                    ws.current.close(1000, "Going away");
                }
            }
        }
        , [notif])

    return <div>
        Hello from lobby: {id}

        {players.map(({ Name, Admin }, idx) => <div key={idx} className={Name === user ? `${primaryColorTxt} font-semibold` : secondaryColorTxt}>{`${Name == user ? "> " : ""}${Name}${Admin ? " (admin)" : ""}`}</div>)}

        <Button onClick={() => alert("Not implemented")}>Start game !</Button>
        <Button secondary onClick={() => copyToClipboard(`${window.location.host}/lobby/${id}`)}>Share lobby !</Button>
        {notif && <Toast message={notif.message} error={notif.error} />}
    </div>
}
