import React from 'react'
import { StyleSheet } from 'react-native'
/* React native paper. */
import { Caption } from 'react-native-paper'

const Small = ({ children }) => {
    return (
        <Caption style={styles.message}>{children}</Caption>
    )
}

const styles = StyleSheet.create({
    message: {
        color: '#666666',
        fontSize: 12,
        textAlign: 'center'
    }
})

export default Small