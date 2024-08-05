import { useState } from "react";
import { useLocation } from "wouter";
import generateName from "./name_gen";
import Input from "./Input";
import Button from "./Button";

export default function JoinLobby({ id }: { id: string }) {
  const [user, setUser] = useState(generateName());
  const setLocation = useLocation()[1];

  return (
    <div className="flex flex-col items-center gap-y-4">
      <div className="font-semibold text-xl">Join lobby {id} !</div>
      <Input
        label="Nickname"
        value={user}
        onChange={(e) => setUser(e.target.value)}
      />
      <Button onClick={() => setLocation(`/lobby/${id}/${user}`)}>
        Join !
      </Button>
    </div>
  );
}
