import React, { Component } from "react";
import Message from "../Message";
import "./ChatHistory.scss";

class ChatHistory extends Component {
    render() {
        console.log(this.props.chatHistory);
        const messages = this.props.chatHistory.map((msg, index) => <Message key={index} message={msg.data} />);

        return (
            <div className="ChatHistory">
                {messages}

            </div>
        );
    }
}

export default ChatHistory;