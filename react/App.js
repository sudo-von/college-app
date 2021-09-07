import React from 'react'
/* React native paper. */
import { Provider as PaperProvider } from 'react-native-paper'
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