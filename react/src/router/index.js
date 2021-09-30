import React from 'react'
/* React navigation. */
import { createDrawerNavigator } from '@react-navigation/drawer'
import { useNavigation } from '@react-navigation/native'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'
/* Custom components. */
import DrawerContent from './Components/DrawerContent'
/* React native paper. */
import { useTheme } from 'react-native-paper'
/* Routes. */
import { publicRoutes, protectedRoutes } from './routes'
/* React native icons. */
import MaterialCommunityIcons from 'react-native-vector-icons/MaterialCommunityIcons'
/* Helpers. */
import { getHeaderLeft, getHeaderTitle } from 'src/helpers/react-navigation-helper'

const Drawer = createDrawerNavigator()

const Router = () => {

  const { authState } = useAuth()
  const { colors } = useTheme()
  const routes = authState.isLoggedIn ? protectedRoutes : publicRoutes
  const navigation = useNavigation()

  return (
    <Drawer.Navigator 
      drawerContent={(props) => authState.isLoggedIn && <DrawerContent {...props}/>}
      screenOptions={screenOptions(colors)}
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
            headerLeft: () => getHeaderLeft(navigation, route),
            ...options
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

export default Router