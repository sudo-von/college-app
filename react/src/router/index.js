import React from 'react'
/* Routers. */
import PublicRouter from './PublicRouter'
import ProtectedRoutes from './ProtectedRoutes'
/* React navigation. */
import { NavigationContainer } from '@react-navigation/native'
/* Refs. */
import { navigationRef } from 'src/refs/navigation.ref'
/* Constants. */
import { theme } from 'src/constants/themes'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'

const Router = () => {

  const { authState } = useAuth()
  console.log('entré')
  console.log(authState)

  return (
    <NavigationContainer theme={theme} ref={navigationRef}>
      { authState.isLoggedIn ?
        <ProtectedRoutes/> :
        <PublicRouter/>
      }
    </NavigationContainer>
  )
}

export default Router