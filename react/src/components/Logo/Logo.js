import React from 'react'
import { StyleSheet, View, Image } from 'react-native'
/* React native paper. */
import { useTheme } from 'react-native-paper'

const Logo = ({ source, size }) => {
    const { colors } = useTheme()
    return (
        <View style={styles({ size, colors}).view}>
            <Image
                style={styles({ size, colors}).image}
                source={source}
            />
        </View>
    )
}

const styles = ({ size, colors }) => StyleSheet.create({ 
    view: {
        borderRadius: 100,
        width: size ? size : 180,
        height: size ? size : 180,
        backgroundColor: colors.primary,
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center'
    },
    image: {
        width: size ? size-25 : 155,
        resizeMode: 'contain'
    }
})


export default Logo
