import React, { useState } from 'react'
import { View, Button, Platform } from 'react-native'
/* Custom components. */
import { Input } from 'src/components'
/* React native paper. */
import { TextInput } from 'react-native-paper'
/* React native community. */
import DateTimePicker from '@react-native-community/datetimepicker'
/* Moment. */
import moment from 'moment'

const Datepicker = ({ field, form, meta, ...rest }) => {
  /* Destructuring props. */
  const { name, value } = field
  const { setFieldValue } = form
  /* Checks if there is a formated date from the props, if not then will return the current date. */
  const getDate = () => {
    if(value){
      return moment(value, 'YYYY-MM-DD').toDate()
    }
    return new Date()
  }
  /* State to handle the visibility for the modal. */
  const [visible, setVisible] = useState(false)
  /* Handle value from the passed props given the field name. */
  const onChange = (event, selectedDate) => {
    if(event.type === 'dismissed'){
      setVisible(false)
      return
    }
    setVisible(false)
    setFieldValue(name, moment(selectedDate).format('YYYY-MM-DD'))
  }

  return (
    <View>
      <View>
        <Input
          disabled
          right={
            <TextInput.Icon
              onPress={() => setVisible(true)}
              color='gray'
              icon='calendar'
            />
          }
          {...meta}
          {...field}
          {...rest}
        />
      </View>
      {visible && (
        <DateTimePicker
          testID="dateTimePicker"
          value={getDate()}
          mode={'date'}
          display="spinner"
          onChange={onChange}
        />
      )}
    </View>
  )
}

export default Datepicker