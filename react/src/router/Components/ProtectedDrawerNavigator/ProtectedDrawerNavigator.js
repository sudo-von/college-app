import React from 'react'
/* React navigation. */
import { createDrawerNavigator } from '@react-navigation/drawer'
import CustomDrawer from '../CustomDrawer'
/* Routes. */
import Home from 'src/pages/Home'
import Logout from 'src/pages/Logout'
/* React native paper. */
import { useTheme } from 'react-native-paper'
/* React native icons. */
import MaterialCommunityIcons from 'react-native-vector-icons/MaterialCommunityIcons'

const Drawer = createDrawerNavigator()

const ProtectedDrawerNavigator = () => {
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
            drawerContent={(props) => <CustomDrawer {...props} />}
            initialRouteName='/'
        >
            <Drawer.Screen 
                name="/" 
                component={Home} 
                options={{ 
                    title: 'Inicio',
                    drawerIcon: ({focused, size}) => (
                        <MaterialCommunityIcons
                           name="home"
                           size={size}
                           color={focused ? '#7cc' : '#ccc'}
                        />
                     ),
                }}
            />
            <Drawer.Screen 
                name="/logout" 
                component={Logout} 
                options={{ 
                        title: 'Cerrar sesión',
                        header: () => null,
                        drawerIcon: ({focused, size}) => (
                            <MaterialCommunityIcons
                               name="logout"
                               size={size}
                               color={focused ? '#7cc' : '#ccc'}
                            />
                         ),
                }}
            />
        </Drawer.Navigator>
    )
}

export default ProtectedDrawerNavigator