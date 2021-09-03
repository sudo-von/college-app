import React from 'react'
/* React navigation. */
import { createDrawerNavigator } from '@react-navigation/drawer'
/* Routes. */
import Home from 'src/pages/Home'

const Drawer = createDrawerNavigator()

const ProtectedRoutes = () =>
    <Drawer.Navigator initialRouteName='/'>
        <Drawer.Screen 
            name="/" 
            component={Home} 
            options={{ title: 'Inicio' }}
        />
    </Drawer.Navigator>

export { ProtectedRoutes }