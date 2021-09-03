import React from 'react'
/* React native paper. */
import { Provider as PaperProvider } from 'react-native-paper'
/* React navigation. */
import { NavigationContainer } from '@react-navigation/native'
import { createNativeStackNavigator } from '@react-navigation/native-stack'
/* Custom components. */
import Router from './src/router'
/* Contexts. */
import { AuthProvider } from './src/providers/auth.provider'
/* Constants. */
import { theme } from 'src/constants/themes'

const App = () =>
  <AuthProvider>
    <PaperProvider theme={theme}>
      <Router/>
    </PaperProvider>
  </AuthProvider>

export default App