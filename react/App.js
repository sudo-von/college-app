import React from 'react'
/* React native paper. */
import { Provider as PaperProvider } from 'react-native-paper'
/* React navigation. */
import { NavigationContainer } from '@react-navigation/native'
import { createNativeStackNavigator } from '@react-navigation/native-stack'
/* Routes. */
import Login from './src/pages/Login'
import Signup from './src/pages/Signup'
/* Contexts. */
import { AuthProvider } from './src/providers/auth.provider'
/* Refs. */
import { navigationRef } from './src/refs/navigation.ref'

const Stack = createNativeStackNavigator()

export default function App() {
  return (
    <AuthProvider>
      <PaperProvider>
        <NavigationContainer ref={navigationRef}>
          <Stack.Navigator>
            <Stack.Screen name="/" component={Login}/>
            <Stack.Screen name="/signup" component={Signup}/>
          </Stack.Navigator>
        </NavigationContainer>
      </PaperProvider>
    </AuthProvider>
  )
}