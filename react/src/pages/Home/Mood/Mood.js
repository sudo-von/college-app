import React from 'react'
import { Dialog, Portal } from 'react-native-paper'
import { Small, Button, Fab, SliderInput } from 'src/components'
import { useMood } from 'src/hooks/useMood'
import { styles } from './Mood.styles'

const Mood = ({ initialMoodValue, minimumValue, maximumValue, minimumText, maximumText, userID }) => {

  const { show, setMood, handleShow, handleUpdateMood } = useMood(initialMoodValue, userID) 

  return (
    <Portal>
      <Dialog visible={show}>
        <Dialog.Content>
          <Fab
            small
            icon='close-thick'
            onPress={() => handleShow(false)}
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
          <Button onPress={() => handleUpdateMood()}>Guardar estado de ánimo</Button>
        </Dialog.Content>
      </Dialog>
    </Portal>
  )
}

export default Mood