import React from 'react'
/* React navigation. */
import { createNativeStackNavigator } from '@react-navigation/native-stack'
/* Routes. */
import Login from 'src/pages/Login'
import Signup from 'src/pages/Signup'

const Stack = createNativeStackNavigator()

const routes = [
    {
        name: '/login',
        component: Login,
        options: {
            header: () => null
        }
    },
    {
        name: '/signup',
        component: Signup,
        options: {
            header: () => null
        }
    }
]

const PublicRouter = () =>
    <Stack.Navigator>
        { routes && routes.map(({ name, component, options}, index) => 
            <Stack.Screen
                key={`${name}-${index}`}
                name={name} 
                component={component} 
                options={options}
            />
        )}
    </Stack.Navigator>

export default PublicRouter