import React, { useState, useEffect } from 'react'
import { Alert } from 'react-native'
import { Dialog, Portal } from 'react-native-paper'
import { Small, Button, Fab, SliderInput } from 'src/components'
import { getMood, sendMood } from 'src/services/mood.service'
import { styles } from './Mood.styles'

const Mood = ({ initialMoodValue, minimumValue, maximumValue, minimumText, maximumText, userID }) => {

  /* If user has not sent its mood today then the modal mood will be show. */
  const [mood, setMood] = useState(initialMoodValue)
  const [show, setShow] = useState(false)
  
  useEffect(() => {
      const searchMood = async () => {
        try{
          await getMood(userID)
        }catch(error){
          setShow(true)
        }
      }
      searchMood()
  }, [])

  const handleMood = async () => {
    try{
      await sendMood(userID, {mood})
    }catch(error){
      Alert.alert('¡Ha ocurrido un error!', error.message)
    }finally{
      setShow(false)
    }
  }

  const handleShow = (show) => setShow(show) 

  return (
    <Portal>
      <Dialog visible={showMood}>
        <Dialog.Content>
          <Fab
            small
            icon='close-thick'
            onPress={handleShow(false)}
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

export default Mood