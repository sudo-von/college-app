import React from 'react'
import { StyleSheet, View } from 'react-native'

const Container = ({ children, justifyContent, style, ...rest}) => 
    <View 
        style={styles({ style, justifyContent}).container} 
        {...rest}
    >
        {children}
    </View>

const styles = ({ style, justifyContent}) => StyleSheet.create({
    container: {
        flex: 1,
        padding: 25,
        justifyContent: justifyContent ? justifyContent : 'center',
        ...style
    }
})

export default Container