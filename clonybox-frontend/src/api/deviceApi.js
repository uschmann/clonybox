import axios from "axios"


export default {

    index() {
        return axios.get('/api/device').then(result => result.data);
    }

}