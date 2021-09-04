import React from 'react'
/* React navigation. */
import { createDrawerNavigator } from '@react-navigation/drawer'
/* Routes. */
import Home from 'src/pages/Home'
import Logout from 'src/pages/Logout'
/* React native paper. */
import { useTheme } from '@react-navigation/native'

const Drawer = createDrawerNavigator()

const ProtectedRoutes = () => {
    const { colors } = useTheme()
    return(
        <Drawer.Navigator 
            screenOptions={{
                drawerStyle: {
                backgroundColor: '#FFFFFF',
                width: 250,
                },
                drawerActiveBackgroundColor: colors.transparency,
                headerStyle:{
                    backgroundColor: colors.primary,
                },
                headerTintColor: 'white',
            }}
            initialRouteName='/'
        >
            <Drawer.Screen 
                name="/" 
                component={Home} 
                options={{ 
                    title: 'Inicio',
                }}
            />
            <Drawer.Screen 
                name="/logout" 
                component={Logout} 
                options={{ 
                        title: 'Cerrar sesión',
                        header: () => null
                }}
            />
        </Drawer.Navigator>
    )
}

export { ProtectedRoutes }