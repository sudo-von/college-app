import React from 'react'
/* React navigation. */
import { createNativeStackNavigator } from '@react-navigation/native-stack'
/* Routes. */
import Home from 'src/pages/Home'
import Suggestions from 'src/pages/Suggestions'

const Stack = createNativeStackNavigator()

const routes = [
    {
        title: 'Inicio',
        name: '/',
        component: Home,
        options: {
            header: () => null
        }
    },
    {
        title: 'Sugerencias',
        name: '/suggestions',
        component: Suggestions,
        options: {}
    }
]

const NestedHomeNavigator = () =>
    <Stack.Navigator>
        { routes && routes.map(({ title, name, component, options }, index) =>
            <Stack.Screen 
                key={`${name}-${index}`}
                name={name}
                component={component}
                options={{
                    title,
                    ...options   
                }}
            />
        )}
    </Stack.Navigator>

export default NestedHomeNavigator