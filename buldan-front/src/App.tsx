import './App.css'
import Home from './components/Home'
import Lobby from './components/Lobby'
import { Route, Switch } from "wouter";


function App() {
    return (
        <>
            <Switch>
                <Route path="/"> <Home /> </Route>
                <Route path="/lobby/:id">
                    {({ id }) => <Lobby id={id} />}
                </Route>
            </Switch>
        </>
    )
}

export default App

