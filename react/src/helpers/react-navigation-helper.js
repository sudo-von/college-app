import React from 'react'
/* React navigation. */
import { DrawerActions, getFocusedRouteNameFromRoute } from '@react-navigation/native'
/* React native paper. */
import { IconButton } from 'react-native-paper'

const getHeaderTitle = (route) => {
    const routeName = getFocusedRouteNameFromRoute(route)
    switch (routeName) {
        case '/':
            return 'Inicio'
        case '/account-configuration':
            return 'Configuración de la cuenta'
        case '/advices':
            return 'Asesorías'
        case '/create-advice':
            return 'Agregar asesoría'
        case '/departments':
            return 'Departamentos'
        case '/create-department':
            return 'Agregar departamento'
        case '/update-advice':
            return 'Actualizar asesoría'
        case '/configuration-home':
            return 'Configuración'
        case '/contact-configuration':
            return 'Configuración de contacto'
        case '/panic-button':
            return 'Botón de pánico'
        case '/suggestions':
            return 'Sugerencias'
    }
}
  
const getHeaderLeft = (navigation, route) => {

    const routeName = getFocusedRouteNameFromRoute(route) ?? '/'
    const BackButton = () => <IconButton onPress={() => navigation.goBack()} icon='arrow-left' color='black'/>
    const MenuButton = () => <IconButton onPress={() => navigation.dispatch(DrawerActions.openDrawer())} icon='menu' color='black'/>
    /* Home screens. */
    switch (routeName) {
        case '/':
            return <MenuButton/>
        case '/configuration-home':
            return <MenuButton/>
        default:
            return <BackButton/>
    }
}

export {
    getHeaderTitle,
    getHeaderLeft
}