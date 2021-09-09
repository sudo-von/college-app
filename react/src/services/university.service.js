/* Helpers. */
import { publicAxios } from 'src/helpers/axios-helper'
/* Constants. */
import { UNIVERSITY } from 'src/constants/endpoints'

export {
    getUniversities
}

const getUniversities = async () => {
    try{
        const universities = await publicAxios.get(`${UNIVERSITY}`)
        return universities.data.results
    }catch(error){
        throw new Error(error.message)
    }
}