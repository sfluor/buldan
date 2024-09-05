import "./App.css";
import Home from "./components/Home";
import Lobby from "./components/Lobby";
import { Route, Switch } from "wouter";
import JoinLobby from "./components/JoinLobby";
import { VERSION } from "./components/constants";

function App() {
  return (
    <>
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
      <div>Version: {VERSION}</div>
    </>
  );
}

export default App;
