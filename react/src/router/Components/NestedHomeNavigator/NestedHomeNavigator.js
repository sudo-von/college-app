import React from 'react'
/* React navigation. */
import { createNativeStackNavigator } from '@react-navigation/native-stack'
/* Routes. */
import HomePage from 'src/pages/Home/HomePage/HomePage'
import PanicButton from 'src/pages/PanicButton'
import Advices from 'src/pages/Advices'
import CreateAdvice from 'src/pages/CreateAdvice'
import CreateSuggestionPage from 'src/pages/Suggestions/CreateSuggestionPage/CreateSuggestionPage'
/* React native paper. */
import { useTheme } from 'react-native-paper'

const Stack = createNativeStackNavigator()

const routes = [
    {
        title: 'Inicio',
        name: '/',
        component: HomePage,
        options: {
            header: () => null
        }
    },
    {
        title: 'Botón de pánico',
        name: '/panic-button',
        component: PanicButton,
        options: {
            header: () => null
        }
    },
    {
        title: 'Asesorías',
        name: '/advices',
        component: Advices,
        options: {
            header: () => null
        }
    },
    {
        title: 'Asesorías',
        name: '/create-advice',
        component: CreateAdvice,
        options: {
            header: () => null
        }
    },
    {
        title: 'Sugerencias',
        name: '/suggestions',
        component: CreateSuggestionPage,
        options: {
            header: () => null
        }
    }
]

const NestedHomeNavigator = () => {
    const { colors } = useTheme()
    return(
        <Stack.Navigator
            screenOptions={{
                headerMode: 'screen',
                headerTintColor: colors.background,
                headerStyle: { backgroundColor: colors.primary },
            }}
            initialRouteName='/'
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

export default NestedHomeNavigator