let msg = {
    name: 'channel add',
    data: {
        name: "hardware support"
    }
}

let subMsg = {
    name: 'channel subscribe'
}

let ws = new WebSocket("ws://localhost:4000/")
ws.onopen = () => {
    ws.send(JSON.stringify(msg))
    ws.send(JSON.stringify(subMsg))
}

ws.onmessage = (e) => {
    console.log(e.data)
}