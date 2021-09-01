import React from 'react'
import { StyleSheet } from 'react-native'
/* React native paper. */
import { Caption } from 'react-native-paper'

const Small = ({ children, ...rest }) => {
    return (
        <Caption style={styles.message} {...rest}>{children}</Caption>
    )
}

const styles = StyleSheet.create({
    message: {
        color: '#666666',
        fontSize: 12
    }
})

export default Small
