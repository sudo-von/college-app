import React, { useState, useEffect } from 'react'
import { StyleSheet, Alert } from 'react-native'
/* React native paper. */
import {Dialog, Portal } from 'react-native-paper'
/* Custom components. */
import Small from 'src/components/Small'
import Button from 'src/components/Button'
import Fab from 'src/components/Fab'
import SliderInput from 'src/components/SliderInput'
/* Services. */
import { getMood, sendMood } from 'src/services/mood.service'

const Mood = ({ initialMoodValue, minimumValue, maximumValue, minimumText, maximumText, userID }) => {

  /* If user has not sent its mood today then the modal mood will be show. */
  const [showMood, setShowMood] = useState(false)
  useEffect(() => {
      const getDailyMood = async () => {
        try{
          await getMood(userID)
        }catch(error){
          setShowMood(true)
        }
      }
      getDailyMood()
  }, [])
  /* Sends user mood. */
  const [mood, setMood] = useState(initialMoodValue)
  const handleMood = async () => {
    try{
      await sendMood(userID, {mood})
      setShowMood(false)
    }catch(error){
      Alert.alert('',error.message)
    }
  }

  return (
    <Portal>
      <Dialog visible={showMood}>
        <Dialog.Content>
          <Fab
            small
            icon='close-thick'
            onPress={() => setShowMood(false)}
          />
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
  )
}

const styles = StyleSheet.create({
  fab: {
    position: 'absolute',
    margin: 15,
    right: 0,
    top: 0
  },
  title: {
    marginTop: 50,
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