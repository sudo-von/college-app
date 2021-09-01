import axios from 'axios'
/* Constants. */
import { API_URL } from 'src/constants'
/* Helpers. */
import { getToken } from './auth-helper'
/* Refs. */
import { navigate } from 'src/refs/navigation.ref'

/* Instance defaults. */
const instance = axios.create({
    baseURL: API_URL,
})

/* Request interceptors. */
instance.interceptors.request.use(async config => {
    try{
        const token = await getToken()
        if (!config.headers.Authorization && token) {
            config.headers.Authorization = token
        }
        return config
    }catch(error){
        throw error
    }
  },
  error => Promise.reject(error)
)

/* Response interceptors. */
instance.interceptors.response.use(response => response,
    (error) => {
        console.log(error.response.status)
        if (error.response.status === 401) {
            navigate('/signup')
        }
        return  Promise.reject(error)
    }
)

export default instance