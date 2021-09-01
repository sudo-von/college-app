import React from 'react'
import { View, StyleSheet } from 'react-native'
/* React native paper. */
import { Caption, IconButton } from 'react-native-paper'

const Alert = ({ title, type }) => {

    let alert = {}
    switch (type) {
        case 'error':
            alert = setAlert({ 
                icon: 'alert-circle', 
                styles: styles.error
            })
            break
        case 'warning':
            alert = setAlert({ 
                icon: 'alert', 
                styles: styles.warning
            })
            break
        case 'info':
            alert = setAlert({ 
                icon: 'information-outline', 
                styles: styles.info
            })
            break
        case 'success':
            alert = setAlert({ 
                icon: 'check', 
                styles: styles.success
            })
            break
        default:
            break
    }

    return (
        <View style={alert.styles}>
            <IconButton icon={alert.icon} color={alert.styles.color}/>
            <Caption style={{flexShrink: 1}}>{title}</Caption>
        </View>
    )
}

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

const setAlert = ({icon, styles}) => (
    {
        icon,
        styles : {
            marginTop: 10,
            borderRadius: 2,
            flexDirection: 'row',
            justifyContent: 'center',
            alignItems: 'center',
            ...styles
        }
    }
)

export default Alert