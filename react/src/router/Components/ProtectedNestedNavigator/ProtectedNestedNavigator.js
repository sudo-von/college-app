import React from 'react'
/* React navigation. */
import { createNativeStackNavigator } from '@react-navigation/native-stack'
/* Routes. */
import Home from 'src/pages/Home'

const Stack = createNativeStackNavigator()

const ProtectedNestedNavigator = () =>
    <Stack.Navigator>
        <Stack.Screen 
            component={Home}
            name="/home"
            options={{header: () => null}}
        />
        <Stack.Screen 
            component={Home} 
            name="/suggestions"
        />
    </Stack.Navigator>

export default ProtectedNestedNavigator