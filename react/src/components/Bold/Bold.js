import React from 'react'
import { StyleSheet, Text } from 'react-native'

const Bold = ({ children, theme }) => {
    return (
        <Text style={styles.text}>
            {children}
        </Text>
    )
}

const styles = StyleSheet.create({
    text: {
        fontWeight: 'bold',
    }
})

export default Bold
