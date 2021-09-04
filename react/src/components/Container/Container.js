import React from 'react'
import { StyleSheet, View } from 'react-native'

const Container = ({ children, style, ...rest}) => 
    <View style={styles(style).container} {...rest}>
        {children}
    </View>

const styles = (style) => StyleSheet.create({
    container: {
        padding: 25,
        flex: 1,
        justifyContent: 'center',
        ...style
    }
})

export default Container