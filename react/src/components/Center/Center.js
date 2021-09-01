import React from 'react'
import { View, StyleSheet } from 'react-native'

const Center = ({ children, style }) => {
    return (
        <View style={styles(style).view}>
            {children}
        </View>
    )
}

const styles = (style) => StyleSheet.create({
    view: {
        alignItems: 'center',
        ...style
    }
})

export default Center