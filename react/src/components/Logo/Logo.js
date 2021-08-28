import React from 'react'
import { StyleSheet, View, Image } from 'react-native'

const Logo = ({ source }) => {
    return (
        <View style={styles.view}>
            <Image
                style={styles.image}
                source={source}
            />
        </View>
    )
}


const styles = StyleSheet.create({ 
    view: {
        width: 180,
        height: 180,
        borderRadius: 100,
        backgroundColor: '#4C9DAF',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center'
    },
    image: {
        width: 156,
        height: 110,
        resizeMode: 'contain'
    }
})


export default Logo
