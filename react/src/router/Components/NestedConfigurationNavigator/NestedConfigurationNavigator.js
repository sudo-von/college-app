import React from 'react'
/* React navigation. */
import { createNativeStackNavigator } from '@react-navigation/native-stack'
/* Routes. */
import ConfigurationPage from 'src/pages/Configuration/ConfigurationPage/ConfigurationPage'
import AccountConfiguration from 'src/pages/AccountConfiguration'
import ContactConfigurationPage from 'src/pages/Configuration/ContactConfigurationPage/ContactConfigurationPage'
/* React native paper. */
import { useTheme } from 'react-native-paper'

const Stack = createNativeStackNavigator()

const routes = [
    {
        title: 'Configuración',
        name: '/configuration-home',
        component: ConfigurationPage,
        options: {
            header: () => null
        }
    },
    {
        title: 'Configuración de la cuenta',
        name: '/account-configuration',
        component: AccountConfiguration,
        options: {
            header: () => null
        }
    },
    {
        title: 'Configuración de contacto',
        name: '/contact-configuration',
        component: ContactConfigurationPage,
        options: {
            header: () => null
        }
    }
]

const NestedConfigurationNavigator = () => {
    const { colors } = useTheme()
    console.log('entré al nested configuraiton nav')
    return(
        <Stack.Navigator
            screenOptions={{
                headerMode: 'screen',
                headerTintColor: 'black',
                headerStyle: { 
                    backgroundColor: 'transparent',
                },
                headerShadowVisible: false
            }}
            >
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
    )
}

export default NestedConfigurationNavigator