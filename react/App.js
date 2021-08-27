import React from 'react'
import { Provider as PaperProvider } from 'react-native-paper'
import { Text, View } from 'react-native'

export default function App() {
  return (
    <PaperProvider>
      <View>
        <Text>Open up App.js to start working on your app!</Text>
      </View>
    </PaperProvider>
  )
}