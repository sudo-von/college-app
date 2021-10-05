import React from 'react'
/* React navigation. */
import { createNativeStackNavigator } from '@react-navigation/native-stack'
/* Routes. */
import HomePage from 'src/pages/Home/HomePage/HomePage'
import PanicButtonPage from 'src/pages/PanicButton/PanicButtonPage/PanicButtonPage'
import AdvicesPage from 'src/pages/Advices/AdvicesPage/AdvicesPage'
import CreateAdvicePage from 'src/pages/Advices/CreateAdvicePage/CreateAdvicePage'
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
        component: PanicButtonPage,
        options: {
            header: () => null
        }
    },
    {
        title: 'Asesorías',
        name: '/advices',
        component: AdvicesPage,
        options: {
            header: () => null
        }
    },
    {
        title: 'Asesorías',
        name: '/create-advice',
        component: CreateAdvicePage,
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