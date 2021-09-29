import { getFocusedRouteNameFromRoute } from '@react-navigation/native'

const getHeaderTitle = (route) => {
  const routeName = getFocusedRouteNameFromRoute(route)
  switch (routeName) {
    case '/configuration-home':
      return 'Configuración'
    case '/contact-configuration':
      return 'Configuración de contacto'
    case '/account-configuration':
      return 'Configuración de la cuenta'
  }
}

export default getHeaderTitle