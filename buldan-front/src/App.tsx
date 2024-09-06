import "./App.css";
import Home from "./components/Home";
import Lobby from "./components/Lobby";
import { Route, Switch } from "wouter";
import JoinLobby from "./components/JoinLobby";
import { VERSION } from "./components/constants";

function App() {
  return (
    <div className="flex flex-col min-h-full">
      <Switch>
        <Route path="/">
          {" "}
          <Home />{" "}
        </Route>
        <Route path="/lobby/:id">{({ id }) => <JoinLobby id={id} />}</Route>
        <Route path="/lobby/:id/:user">
          {({ id, user }) => <Lobby id={id} user={user} />}
        </Route>
      </Switch>
      <a className="p-2 mt-auto justify-self-end self-end text-blue-500 underline hover:text-blue-700 transition-colors duration-300" href={`https://github.com/sfluor/buldan/commit/${VERSION.split("-")[0]}`}>Version: {VERSION}</a>
    </div>
  );
}

export default App;
