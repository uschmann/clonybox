import axios from 'axios'

export default {

    index() {
        return axios.get('/api/playback-config').then(result => result.data);
    },

    show(id) {
        return axios.get(`/api/playback-config/${id}`).then(result => result.data);
    },

    update(playbackConfig) {
        return axios.put(`/api/playback-config/${playbackConfig.id}`, playbackConfig).then(result => result.data);
    },

    delete(id) {
        return axios.delete(`/api/playback-config/${id}`).then(result => result.data);
    }

}