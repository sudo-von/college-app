import React from 'react'
/* React navigation. */
import { createDrawerNavigator } from '@react-navigation/drawer'
/* Routes. */
import Home from 'src/pages/Home'
import Logout from 'src/pages/Logout'

const Drawer = createDrawerNavigator()

const ProtectedRoutes = () =>
    <Drawer.Navigator initialRouteName='/'>
        <Drawer.Screen 
            name="/" 
            component={Home} 
            options={{ title: 'Inicio' }}
        />
        <Drawer.Screen 
            name="/logout" 
            component={Logout} 
            options={
                { 
                    title: 'Cerrar sesión',
                    header: () => null
                }
            }
        />
    </Drawer.Navigator>

export { ProtectedRoutes }