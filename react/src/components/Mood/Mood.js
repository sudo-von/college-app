import React, { useState } from 'react'
import { StyleSheet, View } from 'react-native'
/* React native paper. */
import {Dialog, Portal } from 'react-native-paper'
/* Custom components. */
import Small from 'src/components/Small'
import Center from 'src/components/Center'
import Button from 'src/components/Button'
import SliderInput from 'src/components/SliderInput'
import { FAB } from 'react-native-paper';

const Mood = () => {

  const initialMoodValue = 2.5
  const [visible, setVisible] = useState(true)
  const hideDialog = () => setVisible(false)
  const [mood, setMood] = useState(initialMoodValue)

  return (
      <View>
        <Portal>
          <Dialog visible={visible}>
            <Dialog.Content>
              <Dialog.Title style={styles.title}>¿Cómo te sientes el día de hoy?</Dialog.Title>
              <Small style={styles.small}>Nos importas mucho y nos gustaría saber cómo te sientes el día de hoy.</Small>
              <Small style={styles.small}>Usa el slider de abajo para contárnoslo.</Small>
              <SliderInput
                initialValue={initialMoodValue}
                minimumValue={0}
                maximumValue={5}
                changeValue={setMood}
                minimumText='Triste'
                maximumText='Feliz'
              />
              <Button onPress={hideDialog}>Guardar estado de ánimo</Button>
            </Dialog.Content>
          </Dialog>
        </Portal>
      </View>
  )
}

const styles = StyleSheet.create({
  title: {
    fontSize: 20,
    textAlign: 'center',
    fontWeight: 'bold'
  },
  small: {
    textAlign: 'center',
  }
})

export default Mood