const websocket = new WebSocket("ws://localhost:8080/api/ws")

websocket.onmessage = function (event) {
    broadcastEvent = JSON.parse(event.data);

    switch (broadcastEvent.type) {
        case "playback_config.scanned":
            console.log("Scanned!", broadcastEvent.data.playback_config)
            break;
    }
}