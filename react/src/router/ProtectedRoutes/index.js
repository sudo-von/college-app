import React from 'react'
/* React navigation. */
import { createDrawerNavigator } from '@react-navigation/drawer'
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

const Drawer = createDrawerNavigator()

const routes = [
    {
        name: '/',
        component: NestedHomeNavigator,
        title: 'Inicio',
        icon: 'home',
        options: {
            header: () => null
        } 
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
                    options={{ 
                        title,
                        drawerIcon: ({focused, size}) => (
                            <MaterialCommunityIcons
                            name={icon}
                            size={size}
                            color={focused ? '#7cc' : '#ccc'}
                            />
                        ),
                        ...options
                    }}
                />
            )}
        </Drawer.Navigator>
    )
}

export default ProtectedRoutes