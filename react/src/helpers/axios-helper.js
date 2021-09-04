import axios from 'axios'
/* Constants. */
import { API_URL } from 'src/constants/endpoints'
/* Helpers. */
import { getToken, deleteToken } from './auth-helper'
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
    async (error) => {
        if (error.response?.status === 401) {
            await deleteToken()
            navigate('/login')
        }
        return  Promise.reject(error)
    }
)

export default instance