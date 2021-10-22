import { useState, useEffect } from 'react'
import { getInstalledApps, createInstalledApp } from 'src/services/app.service'

export const useInstalledApps = () => {

    const [ installedApps, setInstalledApps ] = useState([])

    useEffect(() => {
        handleInstalledApps()
    },  [])

    const handleInstalledApps = async () => {
        try{
            const response = await getInstalledApps()
            setInstalledApps(response.results)
        }catch(error){
        }
    }

    const handleCreateInstalledApp = async (installedApp) => {
        try{
            await createInstalledApp(installedApp)
        }catch(error){
        }
    }

    return {
        installedApps,
        handleCreateInstalledApp
    }
} 