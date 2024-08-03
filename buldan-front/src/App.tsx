import "./App.css";
import Home from "./components/Home";
import Lobby from "./components/Lobby";
import { Redirect, Route, Switch } from "wouter";
import generateName from "./components/name_gen";

function App() {
  return (
    <>
      <Switch>
        <Route path="/">
          {" "}
          <Home />{" "}
        </Route>
        <Route path="/lobby/:id">
          {({ id }) => <Redirect to={`/lobby/${id}/${generateName()}`} />}
        </Route>
        <Route path="/lobby/:id/:user">
          {({ id, user }) => <Lobby id={id} user={user} />}
        </Route>
      </Switch>
    </>
  );
}

export default App;
