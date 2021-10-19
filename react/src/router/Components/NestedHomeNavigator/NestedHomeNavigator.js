import React from 'react'
/* React navigation. */
import { createNativeStackNavigator } from '@react-navigation/native-stack'
/* Routes. */
import HomePage from 'src/pages/Home/HomePage/HomePage'
import PanicButtonPage from 'src/pages/PanicButton/PanicButtonPage/PanicButtonPage'
import AdvicesPage from 'src/pages/Advices/AdvicesPage/AdvicesPage'
import DepartmentsPage from 'src/pages/Departments/DepartmentsPage/DepartmentsPage'
import CreateDepartmentPage from 'src/pages/Departments/CreateDepartmentPage/CreateDepartmentPage'
import CreateAdvicePage from 'src/pages/Advices/CreateAdvicePage/CreateAdvicePage'
import UpdateAdvicePage from 'src/pages/Advices/UpdateAdvicePage/UpdateAdvicePage'
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
        title: 'Agregar asesoría',
        name: '/create-advice',
        component: CreateAdvicePage,
        options: {
            header: () => null
        }
    },
    {
        title: 'Actualizar asesoría',
        name: '/update-advice',
        component: UpdateAdvicePage,
        options: {
            header: () => null
        }
    },
    {
        title: 'Departamentos',
        name: '/departments',
        component: DepartmentsPage,
        options: {
            header: () => null
        }
    },
    {
        title: 'Agregar departamento',
        name: '/create-department',
        component: CreateDepartmentPage,
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