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
        fontSize: 22,
        fontWeight: 'bold',
        textAlign: 'left',
        padding: 0,
        margin: 0
    },
    subtitle: {
        lineHeight: 10,
        fontSize: 10,
        textAlign: 'left',
    }
})

export default Header