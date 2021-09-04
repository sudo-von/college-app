import React from 'react'
import { StyleSheet, View, Image } from 'react-native'

const Logo = ({ source, style }) => {
    return (
        <View style={styles(style).view}>
            <Image
                style={styles(style).image}
                source={source}
            />
        </View>
    )
}


const styles = (style) => StyleSheet.create({ 
    view: {
        width: style ? style.width : 180,
        height: style ? style.height : 180,
        borderRadius: 100,
        backgroundColor: '#4C9DAF',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center'
    },
    image: {
        width: style ? style.width-15 : 150,
        resizeMode: 'contain'
    }
})


export default Logo
