import React, { useState } from 'react'
import { StyleSheet, View, Alert } from 'react-native'
/* React native paper. */
import {Dialog, Portal } from 'react-native-paper'
/* Custom components. */
import Small from 'src/components/Small'
import Button from 'src/components/Button'
import SliderInput from 'src/components/SliderInput'
/* Services. */
import { sendMood } from 'src/services/mood.service'

const Mood = ({ initialMoodValue, minimumValue, maximumValue, minimumText, maximumText }) => {

  const [visible, setVisible] = useState(true)
  const [mood, setMood] = useState(initialMoodValue)
  const handleMood = async () => {
    try{
      await sendMood({mood})
      setVisible(false)
    }catch(error){
      Alert.alert('',error.message)
    }
  }

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
                minimumValue={minimumValue}
                maximumValue={maximumValue}
                changeValue={setMood}
                minimumText={minimumText}
                maximumText={maximumText}
                style={styles.sliderInput}
              />
              <Button onPress={() => handleMood()}>Guardar estado de ánimo</Button>
            </Dialog.Content>
          </Dialog>
        </Portal>
      </View>
  )
}

const styles = StyleSheet.create({
  title: {
    fontSize: 24,
    textAlign: 'center',
    fontWeight: 'bold'
  },
  small: {
    textAlign: 'center',
  },
  sliderInput: {
    marginTop: 25,
    marginBottom: 25
  }
})

export default Mood