import './App.css'
import Home from './components/Home'
import Lobby from './components/Lobby'
import { Route, Switch } from "wouter";


function App() {
    return (
        <>
            <Switch>
                <Route path="/"> <Home /> </Route>
                <Route path="/lobby/:id/:user">
                    {({ id, user }) => <Lobby id={id} user={user} />}
                </Route>
            </Switch>
        </>
    )
}

export default App

