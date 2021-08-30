import React from 'react'
import { StyleSheet } from 'react-native'
/* React native paper. */
import { Button } from 'react-native-paper'

const CustomButton = ({loading, children, ...rest}) =>
    <Button style={styles.button} mode='contained' loading={loading} disabled={loading} {...rest}>
        {loading ? 'Cargando... ' : children}
    </Button>

const styles = StyleSheet.create({
    button: {
        marginTop: 40,
        marginBottom: 40,
        backgroundColor: '#4C9DAF',
        color: 'white'
    }
})

export default CustomButton