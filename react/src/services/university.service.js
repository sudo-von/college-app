/* Helpers. */
import { publicAxios, protectedAxios } from 'src/helpers/axios-helper'
/* Constants. */
import { UNIVERSITY } from 'src/constants/endpoints'

export {
    getUniversityByID,
    getUniversities
}

const getUniversityByID = async (universityID) => {
    try {
        const university = await protectedAxios.get(`${UNIVERSITY}/${universityID}`)
        return university.data
    }catch(error){
        throw error
    }
}

const getUniversities = async () => {
    try{
        const universities = await publicAxios.get(`${UNIVERSITY}`)
        return universities.data.results
    }catch(error){
        throw error
    }
}