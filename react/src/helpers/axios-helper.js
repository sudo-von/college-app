import axios from 'axios'
/* Constants. */
import { API_URL } from 'src/constants/endpoints'

/* Instance defaults. */
const instance = axios.create({
    baseURL: API_URL,
})

export default instance