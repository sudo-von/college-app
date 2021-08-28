import React from 'react'
import { StyleSheet } from 'react-native'
/* React native paper. */
import { Button } from 'react-native-paper'

const CustomButton = ({...rest}) =>
    <Button style={styles.button} mode='contained' {...rest}>
        Iniciar sesión
    </Button>

const styles = StyleSheet.create({
    button: {
        marginTop: 35,
        backgroundColor: '#4C9DAF',
        color: 'white'
    }
})

export default CustomButton