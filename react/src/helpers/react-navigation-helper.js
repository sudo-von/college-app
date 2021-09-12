import React from 'react'
/* React navigation. */
import { DrawerActions, getFocusedRouteNameFromRoute } from '@react-navigation/native'
/* React native paper. */
import { IconButton, Menu } from 'react-native-paper'

const getHeaderTitle = (route) => {
    const routeName = getFocusedRouteNameFromRoute(route)
    switch (routeName) {
        case '/':
            return 'Inicio'
        case '/panic-button':
            return 'Botón de pánico'
        case '/suggestions':
            return 'Sugerencias'
        case '/configuration-home':
            return 'Configuración'
        case '/contact-configuration':
            return 'Configuración de contacto'
        case '/account-configuration':
            return 'Configuración de la cuenta'
    }
}
  
const getHeaderButton = (navigation, route) => {

    const routeName = getFocusedRouteNameFromRoute(route) ?? '/'
    const BackButton = () => <IconButton onPress={() => navigation.goBack()} icon='arrow-left' color='white'/>
    const MenuButton = () => <IconButton onPress={() => navigation.dispatch(DrawerActions.openDrawer())} icon='menu' color='white'/>
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
    getHeaderButton
}