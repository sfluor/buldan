import { useState } from "react";
import { primaryColorTxt, secondaryColorTxt, SERVER_URL } from "./constants";
import Input from "./Input";
import Button from "./Button";
import generateName from "./name_gen";
import axios from "axios";
import { useLocation } from "wouter";
import BuldanText from "./BuldanText";

export default function Home() {
  const [user, setUser] = useState(generateName());
  const setLocation = useLocation()[1];

  const createLobby = async () => {
    try {
      const res = await axios.post(`${SERVER_URL}/new-lobby`);
      setLocation(`/lobby/${res.data.id}/${user}`);
    } catch (error) {
      console.error(error);
      alert(`An error occurred ${error}`);
    }
  };

  return (
    <div className="flex flex-col items-center gap-y-16">
      <BuldanText />
      <div className="text-2xl items-center">
        <span className={primaryColorTxt}>Welcome to</span>
        <br />
        <br />
        <Input value={user} onChange={(e) => setUser(e.target.value)} />
      </div>
      <div className="text-2xl">
        <span className={secondaryColorTxt}>Start </span>
        <span className={primaryColorTxt}>guessing !</span>
      </div>
      <div className="flex flex-row self-center gap-x-12">
        <Button onClick={createLobby}>New game</Button>
        <Button secondary>Join game</Button>
      </div>
    </div>
  );
}
