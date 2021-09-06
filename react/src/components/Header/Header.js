import React from 'react'
import { StyleSheet, View } from 'react-native'
/* React native paper. */
import { Title, Caption } from 'react-native-paper'

const Header = ({ title, subtitle }) => {
    return (
        <View style={styles.container}>
            <Title style={styles.title}>{title}</Title>
            <Caption style={styles.subtitle}>{subtitle}</Caption>
        </View>
    )
}

const styles = StyleSheet.create({
    container:{
        marginTop: 20
    },
    title: {
        fontWeight: 'bold',
        textAlign: 'left',
        padding: 0,
        margin: 0
    },
    subtitle: {
        lineHeight: 15,
        fontSize: 12,
        textAlign: 'left',
    }
})

export default Header