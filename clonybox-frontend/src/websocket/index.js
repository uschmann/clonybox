import router from "@/router";
import notification from "@/utils/notification";


const url = `ws://${location.host}/api/ws`
const websocket = new WebSocket(url)

websocket.onmessage = function (event) {
    const broadcastEvent = JSON.parse(event.data);

    switch (broadcastEvent.type) {
        case "playback_config.scanned":
            notification.info(broadcastEvent.data.playback_config.rfid + ' was scanned');
            router.push({name: 'playbackConfig.detail', params: {id: broadcastEvent.data.playback_config.id}})
            break;
    }
}