import React from 'react'
import { View, StyleSheet, TouchableHighlight } from 'react-native'
/* React native paper. */
import { Caption, IconButton } from 'react-native-paper'
/* React navigation. */
import { useNavigation } from '@react-navigation/native';

const NavigationBadge = ({ url, title, icon, color, backgroundColor }) => {
    const navigation = useNavigation()
    return (
            <TouchableHighlight onPress={() => navigation.navigate(url)} underlayColor="white">
                <View style={styles(backgroundColor).view}>
                    <IconButton icon={icon} color={color}/>
                    <Caption style={styles.caption}>{title}</Caption>
                </View>
            </TouchableHighlight>
    )
}

/*
const styles = StyleSheet.create({
    error : {
        backgroundColor: '#f7d8d5',
        color: '#611a15'
    },
    warning: {
        backgroundColor: '#fff4e5',
        color: '#663c00'
    },
    info: {
        backgroundColor: '#e8f4fd',
        color: '#0d5a43'
    },
    success: {
        backgroundColor: '#edf7ed',
        color: '#1e4620'
    }
})
*/

const styles = (backgroundColor) => StyleSheet.create({
    view: {
        marginTop: 20,
        borderRadius: 2,
        flexDirection: 'row',
        alignItems: 'center',
        width: '100%',
        backgroundColor
    },
    caption: {
        flexShrink: 1
    }
})

export default NavigationBadge