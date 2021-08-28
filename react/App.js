import React from 'react'
/* React native paper. */
import { Provider as PaperProvider } from 'react-native-paper'
/* React router. */
import { NativeRouter, Route } from 'react-router-native'
/* Routes. */
import Login from './src/pages/Login'
/* Contexts. */
import { AuthProvider } from './src/providers/auth.provider'

export default function App() {
  return (
    <AuthProvider>
      <PaperProvider>
        <NativeRouter>
          <Route exact path="/" component={Login}/>
        </NativeRouter>
      </PaperProvider>
    </AuthProvider>
  )
}