import React from 'react'
/* React navigation. */
import { createNativeStackNavigator } from '@react-navigation/native-stack'
/* Routes. */
import Home from 'src/pages/Home'
import Suggestions from 'src/pages/Suggestions'

const Stack = createNativeStackNavigator()

const NestedHomeNavigator = () =>
    <Stack.Navigator>
        <Stack.Screen 
            component={Home}
            name="/home"
            options={{header: () => null}}
        />
        <Stack.Screen 
            component={Suggestions}
            options={{ title: 'Sugerencias'}}
            name="/suggestions"
        />
    </Stack.Navigator>

export default NestedHomeNavigator