import axios from 'axios'

export default {
    reset() {
        return axios.post('/api/settings/reset').then(result => result.data)
    }
}