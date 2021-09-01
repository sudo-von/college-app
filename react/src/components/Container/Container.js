import React from 'react'
import { StyleSheet, View } from 'react-native'

const Container = ({ children, ...rest}) => 
    <View style={styles.container} {...rest}>
        {children}
    </View>

const styles = StyleSheet.create({
    container: {
        padding: 25,
        flex: 1,
        justifyContent: 'center'
    }
})

export default Container