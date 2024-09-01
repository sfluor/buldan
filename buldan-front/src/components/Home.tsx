import { useState } from "react";
import { primaryColorTxt, secondaryColorTxt } from "./constants";
import Input from "./Input";
import Button from "./Button";
import generateName from "./name_gen";
import { useLocation } from "wouter";
import BuldanText from "./BuldanText";
import { createLobby } from "./Lobby";

export default function Home() {
  // TODO: save pseudo
  const [user, setUser] = useState(generateName());
  const setLocation = useLocation()[1];

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
        <Button onClick={() => createLobby(user, setLocation)}>New game</Button>
        {/* <Button secondary>Join game</Button>*/}
      </div>
    </div>
  );
}
