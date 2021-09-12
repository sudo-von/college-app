import React from 'react'
import { StyleSheet, Text } from 'react-native'

const Bold = ({ children, ...rest }) =>
        <Text 
            style={styles.text}
            {...rest}
        >
            {children}
        </Text>

const styles = StyleSheet.create({
    text: {
        fontWeight: 'bold',
    }
})

export default Bold