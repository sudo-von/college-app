import { get, post, patch, destroy } from 'src/helpers/protected-axios-helper'

const getDepartmentByID = async (departmentID) => {
    try{
        const advice = await get(`/departments/${departmentID}`)
        return advice.data
    }catch(error){
        throw error
    }
}

const getDepartments = async () => {
    try{
        const advices = await get('/departments')
        return advices.data
    }catch(error){
        throw error
    }
}

const createDepartment = async (departmentPayload) => {
    try{
        await post('/departments', departmentPayload)
    }catch(error){
        throw error
    }
}

const deleteDepartmentByID = async (departmentID) => {
    try{
        await destroy(`/departments/${departmentID}`)
        return '¡Has elimiando el departamento con éxito!'
    }catch(error){
        throw error
    }
}

const updateDepartmentByID = async (departmentID, departmentPayload) => {
    try{
        await patch(`/departments/${departmentID}`, departmentPayload)
    }catch(error){
        throw error
    }
}

export {
    getDepartmentByID,
    getDepartments,
    createDepartment,
    deleteDepartmentByID,
    updateDepartmentByID
}