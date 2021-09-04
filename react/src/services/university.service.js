/* Helpers. */
import axios from 'src/helpers/axios-helper'
/* Constants. */
import { UNIVERSITY } from 'src/constants/endpoints'

export {
    getUniversities
}

const getUniversities = async () => {
    try{
        const universities = await axios.get(`${UNIVERSITY}`)
        return universities.data
    }catch(error){
        throw new Error(error.response.data.message)
    }
}