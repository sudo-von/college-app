import { get, post } from 'src/helpers/protected-axios-helper'

const getInstalledApps = async () => {
    try{
        const installedApps = await get('/installed-app')
        return installedApps.data
    }catch(error){
        throw error
    }
}

const createInstalledApp = async (installedApp) => {
    try{
        await post('/installed-apps', installedApp)
    }catch(error){
        throw error
    }
}

export {
    getInstalledApps,
    createInstalledApp
}