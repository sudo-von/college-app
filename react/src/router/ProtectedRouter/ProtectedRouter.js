import React from 'react'
/* React navigation. */
import { createDrawerNavigator } from '@react-navigation/drawer'
import { useNavigation } from '@react-navigation/native'
/* Custom components. */
import DrawerContent from './Components/DrawerContent'
/* Routes. */
import Logout from 'src/pages/Logout'
/* React native paper. */
import { useTheme } from 'react-native-paper'
/* React native icons. */
import MaterialCommunityIcons from 'react-native-vector-icons/MaterialCommunityIcons'
/* Nested routers. */
import NestedHomeNavigator from './Components/NestedHomeNavigator'
import NestedConfigurationNavigator from './Components/NestedConfigurationNavigator'
/* Helpers. */
import { getHeaderLeft, getHeaderTitle } from 'src/helpers/react-navigation-helper'

const Drawer = createDrawerNavigator()

const routes = [
    {
        name: '/home',
        component: NestedHomeNavigator,
        title: 'Inicio',
        icon: 'home'
    },
    {
        name: '/configuration',
        component: NestedConfigurationNavigator,
        title: 'Configuración',
        icon: 'cog'
    },
    {
        name: '/logout',
        component: Logout,
        title: 'Cerrar sesión',
        icon: 'logout'
    }
]

const ProtectedRouter = () => {

    const { colors } = useTheme()
    const navigation = useNavigation()

    return(
        <Drawer.Navigator 
            screenOptions={screenOptions(colors)}
            drawerContent={(props) => <DrawerContent {...props}/>}
        >
            { routes && routes.map(({ name, component, title, icon }, index) =>
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
                        headerLeft: () => getHeaderLeft(navigation, route)
                    })}
                />
            )}
        </Drawer.Navigator>
    )
}

const screenOptions = (colors) => ({
    drawerStyle: {
        backgroundColor: colors.background
    },
    drawerActiveBackgroundColor: colors.transparency,
    drawerActiveTintColor: colors.primary,
    drawerInactiveTintColor: colors.accent,
    headerStyle:{
        backgroundColor: colors.primary,
    },
    headerTintColor: colors.background
})

export default ProtectedRouter