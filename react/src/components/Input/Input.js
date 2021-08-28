import React from 'react'
import { StyleSheet } from 'react-native'
/* React native paper. */
import { TextInput } from 'react-native-paper'

const Input = ({ ...rest }) =>
    <TextInput
        mode='outlined'
        style={styles.input}
        theme={{ colors: { primary: '#2d798a' }}}
        {...rest}
    />
 
const styles = StyleSheet.create({
    input: {
        marginTop: 5,
        backgroundColor: 'white'
    }
})


export default Input
