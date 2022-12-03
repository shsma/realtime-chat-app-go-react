import React, { Component } from "react";

import "./App.css";
import { connect, sendMsg } from "./api";
import Header from "./components/Header";
import ChatHistory from "./components/ChatHistory";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
        chatHistory: []
    }
  }

  componentDidMount() {
    connect((msg) => {
        console.log("New Message")
        this.setState(prevState => ({
            chatHistory: [...this.state.chatHistory, msg]
        }))
        console.log(this.state);
    });
  }

  send() {
    console.log("Sending hello");
    sendMsg("Hello Shady!");
  }

  render() {
    return (
        <div className="App">
            <Header />
            <ChatHistory chatHistory={this.state.chatHistory} />
            <button onClick={this.send}>Send Message</button>
        </div>
    );
  }
}

export default App;