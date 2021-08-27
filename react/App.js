import React from 'react'
/* React-native-paper. */
import { Provider as PaperProvider } from 'react-native-paper'
/* React router. */
import { NativeRouter, Route, Link } from 'react-router-native'
/* Routes. */
import Login from './src/pages/Login'

export default function App() {
  return (
    <PaperProvider>
      <NativeRouter>
        <Route exact path="/" component={Login}/>
      </NativeRouter>
    </PaperProvider>
  )
}