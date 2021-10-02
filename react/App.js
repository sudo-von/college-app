import React from 'react'
import Router from './src/router'
import { Provider as PaperProvider } from 'react-native-paper'
import { AuthProvider } from './src/providers/auth.provider'
import { theme } from 'src/constants/themes'
import { navigationRef } from 'src/refs/navigation.ref'
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