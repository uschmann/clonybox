import axios from 'axios'

export default {

    index() {
        return axios.get('/api/playback-config').then(result => result.data);
    }

}