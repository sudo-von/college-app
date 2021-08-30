import axios from 'axios'
/* Constants. */
import { API_URL } from 'src/constants'
/* Helpers. */
import { getToken } from './auth-helper'

/* Instance defaults. */
const instance = axios.create({
    baseURL: API_URL,
})

/* Request interceptors. */
instance.interceptors.request.use(async config => {
    if (!config.headers.Authorization) {
        // const token = await getToken()
        // if(token){
        //     config.headers.Authorization = `Bearer ${token}`;
        // }
    }
    return config
  },
  error => Promise.reject(error)
)

export default instance