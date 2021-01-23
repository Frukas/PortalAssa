var socket = new WebSocket("ws://localhost:9090/ws")

let connect = cb => {
    console.log("Attempting Connection...")

    socket.onopen = () => {
        console.log("Succefully Connected");
    };

    socket.onmessage = msg => {
        console.log(msg);
        cb(msg)
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
};

let sendMsg = msg => {
    console.log("sending msg: ", msg);
    socket.send(msg);
};

export { connect, sendMsg };