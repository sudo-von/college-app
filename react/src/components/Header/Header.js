import React from 'react'
import { StyleSheet, View } from 'react-native'
/* Custom components. */
import { Bold } from 'src/components'
/* React native paper. */
import { Title, Caption } from 'react-native-paper'

const Header = ({ title, subtitle }) => {
    return (
        <View style={styles.container}>
            <Title><Bold>{title}</Bold></Title>
            <Caption style={styles.subtitle}>{subtitle}</Caption>
        </View>
    )
}

const styles = StyleSheet.create({
    container:{
        marginTop: 20
    },
    subtitle: {
        lineHeight: 12,
        fontSize: 12,
    }
})

export default Header