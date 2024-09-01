import { useLocation } from "wouter";
import Button from "./Button";
import { mainViewCols } from "./constants";
import { createLobby, Player } from "./Lobby";
import PlayerBoxes from "./PlayerBoxes";

export default function LobbyEndGame({
  user,
  players,
}: {
  user: string;
  players: Player[];
}) {
  const setLocation = useLocation()[1];

  const sorted = [...players].sort((p1, p2) => {
    if (p1.Points > p2.Points) return -1;
    if (p1.Points < p2.Points) return 1;
    return 0;
  });
  return (
    <>
      <h3 className="text-2xl mb-4">🏆 Final leaderboard ! 🏆</h3>
      <Button onClick={() => createLobby(user, setLocation)}>New game</Button>
      <PlayerBoxes players={sorted} user={user} />
    </>
  );
}
