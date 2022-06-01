import axios from 'axios'

export default {
    index() {
        return axios.get('/api/info').then(result => result.data);
    }
}