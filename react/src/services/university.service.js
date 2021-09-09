/* Helpers. */
import { publicAxios } from 'src/helpers/axios-helper'
/* Constants. */
import { UNIVERSITY } from 'src/constants/endpoints'
import { UNIVERSITY_ERRORS } from 'src/constants/errors'

export {
    getUniversities
}

const getUniversities = async () => {
    try{
        const universities = await publicAxios.get(`${UNIVERSITY}`)
        return universities.data.results
    }catch(error){
        throw new Error(UNIVERSITY_ERRORS[error.response.data.code]['esp'])
    }
}