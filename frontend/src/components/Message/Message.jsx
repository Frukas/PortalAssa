import React, { Component } from "react";
import "./Message.scss";

class Message extends Component {
    constructor(props) {
        super(props);
        let temp = JSON.parse(this.props.message)
        this.state = {
            message: temp
        };
    }

    render() {
        return <div className="Message">{this.state.message.body}
            <input placeholder="Nome" />
            <input placeholder="Horário" />
            <button>Start</button>
            <button>End</button>
        </div>
    }
}

export default Message;