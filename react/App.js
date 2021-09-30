import React from 'react'
/* React native paper. */
import { Provider as PaperProvider } from 'react-native-paper'
/* Custom components. */
import Router from './src/router'
/* Contexts. */
import { AuthProvider } from './src/providers/auth.provider'
/* Constants. */
import { theme } from 'src/constants/themes'
/* Refs. */
import { navigationRef } from 'src/refs/navigation.ref'
/* React navigation. */
import { NavigationContainer } from '@react-navigation/native'

const App = () =>
  <AuthProvider>
    <PaperProvider theme={theme}>
      <NavigationContainer theme={theme} ref={navigationRef}>
        <Router/>
      </NavigationContainer>
    </PaperProvider>
  </AuthProvider>

export default App