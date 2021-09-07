import React from 'react'
/* React navigation. */
import { createNativeStackNavigator } from '@react-navigation/native-stack'
/* Routes. */
import Configuration from 'src/pages/Configuration'
import AccountConfiguration from 'src/pages/AccountConfiguration'

const Stack = createNativeStackNavigator()

const routes = [
    {
        title: 'Configuración',
        name: '/configuration-home',
        component: Configuration,
        options: {
            header: () => null
        }
    },
    {
        title: 'Configuración de la cuenta',
        name: '/account-configuration',
        component: AccountConfiguration,
        options: {}
    }
]

const NestedConfigurationNavigator = () =>
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

export default NestedConfigurationNavigator