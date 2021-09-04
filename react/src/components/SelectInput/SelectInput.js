import React, { useState } from 'react'
import { StyleSheet, View } from 'react-native'
/* React native paper. */
import DropDown from 'react-native-paper-dropdown'
import { HelperText } from 'react-native-paper'

const SelectInput = ({ field, form, meta, error, label, data }) => {
  
  /* Destructuring props. */
  const { name, value } = field
  const { setFieldValue } = form
  /* Checks if there is a formated date from the props, if not then will return the current date. */
  const [showDropDown, setShowDropDown] = useState(false)
  const handleSelectedOption = (selected) => {
    setFieldValue(name, selected)
  }

  return (
    <View style={styles.containerStyle}>
      <DropDown
        label={label}
        mode='outlined'
        visible={showDropDown}
        showDropDown={() => setShowDropDown(true)}
        onDismiss={() => setShowDropDown(false)}
        value={value}
        setValue={handleSelectedOption}
        list={data}
      />
      { error && 
        <HelperText 
            padding='none' 
            type='error' 
            visible={true}
        >
            {error}
        </HelperText>
      }
    </View>
  )
}
  
const styles = StyleSheet.create({
  containerStyle: {
    marginTop: 5,
  },
})
  
export default SelectInput