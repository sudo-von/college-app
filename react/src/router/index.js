import React from 'react'
import DrawerContent from './Components/DrawerContent/DrawerContent'
import MaterialCommunityIcons from 'react-native-vector-icons/MaterialCommunityIcons'
import { createDrawerNavigator } from '@react-navigation/drawer'
import { useNavigation } from '@react-navigation/native'
import { useUser } from 'src/hooks/useUser'
import { useTheme } from 'react-native-paper'
import { publicRoutes, protectedRoutes } from './routes'
import { getHeaderLeft, getHeaderTitle } from 'src/helpers/react-navigation-helper'

const Drawer = createDrawerNavigator()

const Router = () => {

  const { isLoggedIn } = useUser()
  const { colors } = useTheme()

  const routes = isLoggedIn ? protectedRoutes : publicRoutes
  const navigation = useNavigation()

  const handleDrawerContent = (props) => isLoggedIn && <DrawerContent {...props}/>

  return (
    <Drawer.Navigator 
      drawerContent={handleDrawerContent}
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
  headerTintColor: colors.background,
  headerStyle:{
    backgroundColor: colors.primary,
  }
})

export default Router