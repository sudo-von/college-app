import React, { useState } from 'react'
import { StyleSheet, View } from 'react-native'
/* React native paper. */
import DropDown from 'react-native-paper-dropdown'
  
const SelectInput = ({ label, data }) => {
  
  const [showDropDown, setShowDropDown] = useState(false)
  const [selectedOption, setSelectedOption] = useState('')
  
  return (
    <View style={styles.containerStyle}>
      <DropDown
        label={label}
        mode='outlined'
        visible={showDropDown}
        showDropDown={() => setShowDropDown(true)}
        onDismiss={() => setShowDropDown(false)}
        value={selectedOption}
        setValue={setSelectedOption}
        list={data}
      />
    </View>
  )
}
  
const styles = StyleSheet.create({
  containerStyle: {
    marginTop: 5,
  },
})
  
export default SelectInput