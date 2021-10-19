/* Routes. */
import LoginPage from 'src/pages/Login/LoginPage/LoginPage'
import SignupPage from 'src/pages/Signup/SignupPage/SignupPage'
import LogoutPage from 'src/pages/Logout/LogoutPage/LogoutPage'
import NestedHomeNavigator from 'src/router/Components/NestedHomeNavigator'
import NestedConfigurationNavigator from 'src/router/Components/NestedConfigurationNavigator'

const publicRoutes = [
  {
    name: '/login',
    component: LoginPage,
    title: '',
    icon: '',
    options: {
      headerShown: false,
      swipeEnabled: false
    }
  },
  {
    name: '/signup',
    component: SignupPage,
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
    component: LogoutPage,
    title: 'Cerrar sesión',
    icon: 'logout',
    options: {
      header: () => null
    }
  }
]

export {
  publicRoutes,
  protectedRoutes
}