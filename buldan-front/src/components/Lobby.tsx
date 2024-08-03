import { useEffect, useRef, useState } from "react";
import { LOBBY_URL } from "./constants";
import Toast from "./Toast";
import { useLocation } from "wouter";
import LobbyWaitRoom from "./LobbyWaitRoom";
import LobbyRound from "./LobbyRound";
import BuldanText from "./BuldanText";

export interface Player {
    Name: string
    Admin: boolean
    Lost?: boolean
}

interface Notification {
    message: string
    error?: boolean
}

export interface Guess {
    Guess: string
    Player: string
    Correct: boolean
    Flag?: string
}

enum LobbyState {
    WaitingRoom,
    Round,
    BetweenRound,
}

export interface RoundState {
    Players: Player[]
    Guesses: Guess[]
    Letter: string
    Remaining: number
    CurrentPlayerIndex: number
    CurrentPlayerRemainingGuesses: number
}

export default function Lobby({ id, user }: { id: string, user: string }) {
    const [players, setPlayers] = useState<Player[]>([]);
    const ws = useRef<WebSocket | null>(null);
    const [notif, setNotif] = useState<Notification | null>(null);
    const [state, setState] = useState(LobbyState.WaitingRoom);
    const [round, setRound] = useState<RoundState | null>(null);

    // (otherwise not happy about the any)
    // eslint-disable-next-line
    const send = (content: any) => {
        if (ws.current) {
            ws.current.send(JSON.stringify(content))
        } else {
            // TODO
            console.error("web socket isn't initialized yet");
        }
    }

    const startGame = () => send({ Type: "start-game" });

    const sendGuess = (guess: string) => send({ Type: "guess", Guess: guess });

    const url = `${LOBBY_URL}/${id}/${user}`;
    const shareUrl = `${window.location.protocol}//${window.location.host}/lobby/${id}`;
    const setLocation = useLocation()[1];

    useEffect(
        () => {
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

            return () => {
                // Cleanup web socket on unmount;
                if (ws.current) {
                    console.log("Closing websocket connection", ws);
                    ws.current.close(1000, "Going away");
                }
            }
        }, [url]
    );

    useEffect(
        () => {
            if (ws.current === null) {
                return;
            }

            const notifAndRedirect = (notif: Notification) => {
                setNotif(notif);
                setTimeout(() => {
                    setNotif(null);
                    setLocation("/");
                }, 5000);
            }

            ws.current.onmessage = function(event) {
                const json = JSON.parse(event.data);
                if (json.Type === "players") {
                    setPlayers(json.Players);
                } else if (json.Type === "new-round") {
                    setState(LobbyState.Round);
                    setRound(json.Round);
                    console.log("New round", json);
                } else if (json.Type === "guess") {
                    setRound(json.Round);
                    console.log("New round", json);
                } else {
                    notifAndRedirect({ message: `Unknown payload received ${json.Type}: ${event.data}`, error: true });
                }
            };

            ws.current.onclose = function(event) {
                console.warn("Closed websocket", event);
                if (notif === null) {
                    notifAndRedirect({ message: `An error occurred: ${event.reason}. Redirecting back home...`, error: true });
                }
            }
        }
        , [notif, setLocation])

    let component;
    if (state === LobbyState.WaitingRoom) {
        component = <LobbyWaitRoom players={players} user={user} shareUrl={shareUrl} startGame={startGame} />
    } else if (state === LobbyState.Round) {
        component = <LobbyRound user={user} round={round} sendGuess={sendGuess} />
    } else if (state === LobbyState.BetweenRound) {
        component = <div> Loading... </div>
    } else {
        component = <div> Unexpected lobby state ! </div>
    }

    return <div className="p-4">
    <BuldanText />
        <LobbyHeader shareUrl={shareUrl} id={id} />

        {component}

        {notif && <Toast message={notif.message} error={notif.error} />}
    </div>
}

function LobbyHeader({ id, shareUrl }: { id: string, shareUrl: string }) {
    return <div className="p-4 my-4 text-lg bg-orange-100 rounded-md"> <a className="font-medium text-blue-600 dark:text-blue-500 hover:underline" href={shareUrl} target="_blank"> Lobby {id}</a></div>
}
