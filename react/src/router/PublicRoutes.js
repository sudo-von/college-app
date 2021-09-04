import React from 'react'
/* React navigation. */
import { createNativeStackNavigator } from '@react-navigation/native-stack'
/* Routes. */
import Login from 'src/pages/Login'
import Signup from 'src/pages/Signup'

const Stack = createNativeStackNavigator()

const PublicRoutes = () =>
    <Stack.Navigator>
        <Stack.Screen 
            name="/login" 
            component={Login} 
            options={{header: () => null}}
        />
        <Stack.Screen 
            name="/signup" 
            component={Signup}
            options={{header: () => null}}
        />
    </Stack.Navigator>

export { PublicRoutes }