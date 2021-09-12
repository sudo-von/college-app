import React from 'react'
/* React navigation. */
import { createDrawerNavigator } from '@react-navigation/drawer'
import { useNavigation } from '@react-navigation/native'
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
/* Helpers. */
import { getHeaderButton, getHeaderTitle } from 'src/helpers/react-navigation-helper'

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

    /* Handles drawer theme. */
    const { colors } = useTheme()
    const screenOptions = {
        headerStyle:{
            backgroundColor: colors.primary,
        },
        headerTintColor: colors.background,
        drawerStyle: {
            backgroundColor: colors.background
        },
        drawerActiveBackgroundColor: colors.transparency,
        drawerActiveTintColor: colors.primary,
        drawerInactiveTintColor: colors.accent,
    }
    const navigations = useNavigation()

    return(
        <Drawer.Navigator 
            screenOptions={screenOptions}
            drawerContent={(props) => <DrawerContent {...props}/>}
        >
            { routes && routes.map(({ name, component, title, icon, options }, index) =>
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
                                color={focused ? colors.primary : colors.accent }
                            />
                        ),
                        headerTitle: getHeaderTitle(route),
                        headerLeft: () => getHeaderButton(navigations, route),
                        ...options,
                    })}
                />
            )}
        </Drawer.Navigator>
    )
}

export default ProtectedRoutes