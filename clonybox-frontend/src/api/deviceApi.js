import axios from "axios"


export default {

    index() {
        return axios.get('/api/device').then(result => result.data);
    },

    setDefault(id) {
        return axios.post('/api/device/default/' + id).then(result => result.data);
    }

}