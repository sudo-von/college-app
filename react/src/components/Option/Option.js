import React from 'react'
import { StyleSheet } from 'react-native'
import { View, Text, Image } from 'react-native'

const Option = ({ source, label }) => {
    return (
        <View style={styles.view}>
            <Image style={styles.image} source={{ uri: source }}/>
            <Text>{label}</Text>
        </View>
    )
}

const styles = StyleSheet.create({
    view: {
        flex: 1,
        width: '100%',
        flexDirection: 'row',
        justifyContent: 'center',
        alignItems: 'center'
    },
    image: {
        marginRight: 10,
        width: 30,
        height: 30,
        resizeMode: 'contain'
    }
})

export default Option
