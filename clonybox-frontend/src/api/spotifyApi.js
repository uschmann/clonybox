import axios from 'axios'

export default {
    search(query) {
        return axios.post('/api/spotify/search', query).then(result => result.data);
    }
}