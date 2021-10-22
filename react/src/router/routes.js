import LoginPage from 'src/pages/Login/LoginPage/LoginPage'
import SignupPage from 'src/pages/Signup/SignupPage/SignupPage'
import LogoutPage from 'src/pages/Logout/LogoutPage/LogoutPage'
import NestedHomeNavigator from 'src/router/Components/NestedHomeNavigator/NestedHomeNavigator'
import NestedConfigurationNavigator from 'src/router/Components/NestedConfigurationNavigator/NestedConfigurationNavigator'
import HomePage from 'src/pages/Home/HomePage/HomePage'
import PanicButtonPage from 'src/pages/PanicButton/PanicButtonPage/PanicButtonPage'
import AdvicesPage from 'src/pages/Advices/AdvicesPage/AdvicesPage'
import DepartmentsPage from 'src/pages/Departments/DepartmentsPage/DepartmentsPage'
import CreateDepartmentPage from 'src/pages/Departments/CreateDepartmentPage/CreateDepartmentPage'
import UpdateDepartmentPage from 'src/pages/Departments/UpdateDepartmentPage/UpdateDepartmentPage'
import CreateAdvicePage from 'src/pages/Advices/CreateAdvicePage/CreateAdvicePage'
import UpdateAdvicePage from 'src/pages/Advices/UpdateAdvicePage/UpdateAdvicePage'
import CreateSuggestionPage from 'src/pages/Suggestions/CreateSuggestionPage/CreateSuggestionPage'
import ConfigurationPage from 'src/pages/Configuration/ConfigurationPage/ConfigurationPage'
import UpdateAccountConfigurationPage from 'src/pages/Configuration/UpdateAccountConfigurationPage/UpdateAccountConfigurationPage'
import ContactConfigurationPage from 'src/pages/Configuration/ContactConfigurationPage/ContactConfigurationPage'

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

const nestedHomeRoutes = [
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
    title: 'Actualizar departamento',
    name: '/update-department',
    component: UpdateDepartmentPage,
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

const nestedConfigurationRoutes = [
  {
    title: 'Configuración',
    name: '/configuration-home',
    component: ConfigurationPage,
    options: {
      header: () => null
    }
  },
  {
    title: 'Configuración de la cuenta',
    name: '/account-configuration',
    component: UpdateAccountConfigurationPage,
    options: {
      header: () => null
    }
  },
  {
    title: 'Configuración de contacto',
    name: '/contact-configuration',
    component: ContactConfigurationPage,
    options: {
      header: () => null
    }
  }
]

export {
  publicRoutes,
  protectedRoutes,
  nestedHomeRoutes,
  nestedConfigurationRoutes
}