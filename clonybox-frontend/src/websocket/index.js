import router from "@/router";

const websocket = new WebSocket("ws://localhost:8080/api/ws")

websocket.onmessage = function (event) {
    const broadcastEvent = JSON.parse(event.data);

    switch (broadcastEvent.type) {
        case "playback_config.scanned":
            router.push({name: 'playbackConfig.detail', params: {id: broadcastEvent.data.playback_config.id}})
            break;
    }
}