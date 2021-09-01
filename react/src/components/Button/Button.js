import React from 'react'
import { StyleSheet } from 'react-native'
/* React native paper. */
import { Button } from 'react-native-paper'

const CustomButton = ({loading, loadingMessage='Cargando...', children, ...rest}) =>
    <Button 
        style={styles.button} 
        mode='contained' 
        loading={loading} 
        disabled={loading} 
        labelStyle={styles.label}
        {...rest}
    >
        {loading ? loadingMessage : children}
    </Button>

const styles = StyleSheet.create({
    label: {
        color: 'white'
    },
    button: {
        marginTop: 40,
        marginBottom: 40,
        color: 'white'
    }
})

export default CustomButton