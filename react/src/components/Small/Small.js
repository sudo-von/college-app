import React from 'react'
import { StyleSheet } from 'react-native'
/* React native paper. */
import { Caption } from 'react-native-paper'

const Small = ({ children, style, ...rest }) => {
    return (
        <Caption style={styles(style).message} {...rest}>{children}</Caption>
    )
}

const styles = (style) => StyleSheet.create({
    message: {
        color: '#666666',
        fontSize: 12,
        ...style
    }
})

export default Small
