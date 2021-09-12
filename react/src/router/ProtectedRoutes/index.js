import React from 'react'
import { StyleSheet, Text } from 'react-native'
/* React navigation. */
import { createDrawerNavigator } from '@react-navigation/drawer'
import { getFocusedRouteNameFromRoute } from '@react-navigation/native'
/* Custom components. */
import DrawerContent from './DrawerContent'
/* Routes. */
import Logout from 'src/pages/Logout'
/* React native paper. */
import { useTheme } from 'react-native-paper'
/* React native icons. */
import MaterialCommunityIcons from 'react-native-vector-icons/MaterialCommunityIcons'
/* Nested routers. */
import NestedHomeNavigator from './NestedHomeNavigator'
import NestedConfigurationNavigator from './NestedConfigurationNavigator'

const getHeaderTitle = (route) => {
  const routeName = getFocusedRouteNameFromRoute(route)
  switch (routeName) {
    case '/':
        return 'Inicio'
    case '/panic-button':
        return 'Botón de pánico'
    case '/suggestions':
        return 'Sugerencias'
    case '/configuration-home':
        return 'Configuración'
    case '/contact-configuration':
        return 'Configuración de contacto'
    case '/account-configuration':
        return 'Configuración de la cuenta'
  }
}

const Drawer = createDrawerNavigator()

const routes = [
    {
        name: '/home',
        component: NestedHomeNavigator,
        title: 'Inicio',
        icon: 'home',
        options: {} 
    },
    {
        name: '/configuration',
        component: NestedConfigurationNavigator,
        title: 'Configuración',
        icon: 'cog',
        options: {} 
    },
    {
        name: '/logout',
        component: Logout,
        title: 'Cerrar sesión',
        icon: 'logout',
        options: {} 
    },
]

const ProtectedRoutes = () => {
    const { colors } = useTheme()
    return(
        <Drawer.Navigator 
            screenOptions={{
                drawerStyle: {
                    backgroundColor: colors.background,
                },
                drawerActiveBackgroundColor: colors.transparency,
                headerStyle:{
                    backgroundColor: colors.primary,
                },
                headerTintColor: colors.background,
            }}
            drawerContent={(props) => <DrawerContent {...props}/>}
        >
            { routes && routes.map(({ name, component, title, icon, options}, index) =>
                <Drawer.Screen 
                    key={`${name}-${index}`}
                    name={name}
                    component={component} 
                    options={({route}) => ({ 
                        title,
                        drawerIcon: ({focused, size}) => (
                            <MaterialCommunityIcons
                            name={icon}
                            size={size}
                            color={focused ? '#7cc' : '#ccc'}
                            />
                        ),
                        headerTitle: getHeaderTitle(route),
                        ...options
                    })}
                />
            )}
        </Drawer.Navigator>
    )
}

export default ProtectedRoutes