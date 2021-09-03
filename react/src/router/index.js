import React from 'react'
/* Custom components. */
import { PublicRoutes } from './PublicRoutes'
import { ProtectedRoutes } from './ProtectedRoutes'
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
  return (
    <NavigationContainer theme={theme} ref={navigationRef}>
        { authState.isLoggedIn ?
          <ProtectedRoutes/> :
          <PublicRoutes/>
        }
    </NavigationContainer>
  )
}

export default Router