import * as Location from 'expo-location'

const checkPermissions = async () => {
    try{
        let { status } = await Location.requestForegroundPermissionsAsync()
        if (status !== 'granted') {
            throw new Error('Considera habilitar el permiso de ubicación, sin este permiso el botón de pánico no podrá acceder a tu ubicación.')
        }
    }catch(error){
        throw error
    }
}

const getCoords = async () => {
    try{
        let { coords } = await Location.getCurrentPositionAsync({ accuracy: 1 })
        return coords
    }catch(error){
        throw new Error('No ha sido posible acceder a tu ubicación.')
    }
}

export {
    checkPermissions,
    getCoords
}