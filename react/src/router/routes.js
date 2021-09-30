/* Routes. */
import Login from 'src/pages/Login'
import Signup from 'src/pages/Signup'
import Logout from 'src/pages/Logout'
import NestedHomeNavigator from 'src/router/Components/NestedHomeNavigator'
import NestedConfigurationNavigator from 'src/router/Components/NestedConfigurationNavigator'

const publicRoutes = [
  {
    name: '/login',
    component: Login,
    title: '',
    icon: '',
    options: {
      headerShown: false,
      swipeEnabled: false
    }
  },
  {
    name: '/signup',
    component: Signup,
    title: '',
    icon: '',
    options: {
      headerShown: false,
      swipeEnabled: false
    }
  }
]

const protectedRoutes = [
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
  }
]

export {
  publicRoutes,
  protectedRoutes
}