import React from 'react'
import { createNativeStackNavigator } from '@react-navigation/native-stack'
import { nestedConfigurationRoutes } from 'src/router/routes' 
import { useTheme } from 'react-native-paper'

const Stack = createNativeStackNavigator()

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
            { nestedConfigurationRoutes && nestedConfigurationRoutes.map(({ title, name, component, options }, index) =>
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