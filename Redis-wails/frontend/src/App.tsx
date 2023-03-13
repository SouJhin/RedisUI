import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import {ConnectionList} from "../wailsjs/go/main/App";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    function greet() {
        // Greet(name).then(updateResultText);
        ConnectionList().then(res => {
            setResultText(res.data)
            setName(res.code)
        })

    }

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo"/>
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <div>{name}</div>
                <div>{resultText}</div>
                <button className="btn" onClick={greet}>Greet</button>
            </div>
        </div>
    )
}

export default App
