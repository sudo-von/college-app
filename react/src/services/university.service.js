import { get } from 'src/helpers/protected-axios-helper'
import { get as publicGet } from 'src/helpers/protected-axios-helper'

const getUniversityByID = async (universityID) => {
    try {
        const university = await get(`/universities/${universityID}`)
        return university.data
    }catch(error){
        throw error
    }
}

const getUniversities = async () => {
    try{
        const universities = await publicGet('/universities')
        return universities.data
    }catch(error){
        throw error
    }
}

export {
    getUniversityByID,
    getUniversities
}
