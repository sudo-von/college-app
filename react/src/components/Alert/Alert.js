import React, { useState } from 'react'
import { View, StyleSheet } from 'react-native'
/* React native paper. */
import { IconButton, Button, Paragraph, Dialog, Portal } from 'react-native-paper'

const Alert = ({ title, message, type }) => {

    const [visible, setVisible] = useState(true)
    const hideDialog = () => setVisible(false)

    let dialogStyles = {}
    switch (type) {
        case 'error':
            dialogStyles = styles.error
            break
        case 'warning':
            dialogStyles = styles.warning
            break
        case 'info':
            dialogStyles = styles.info
            break
        case 'success':
            dialogStyles = styles.success
            break
        default:
            break
    }

    return (
        <View>
            <Portal>
            <Dialog visible={visible} onDismiss={hideDialog} style={dialogStyles}>
                {title &&
                    <Dialog.Title 
                        theme={{ colors: { text: dialogStyles.color }}}>{title}</Dialog.Title>
                }
                {message &&
                    <Dialog.Content>
                        <Paragraph style={dialogStyles.color}>{message}</Paragraph>
                    </Dialog.Content>
                }
                <Dialog.Actions>
                <Button color={dialogStyles.color} onPress={hideDialog}>Cerrar</Button>
                </Dialog.Actions>
            </Dialog>
            </Portal>
        </View>
    )
}

const styles = StyleSheet.create({
    error : {
        backgroundColor: '#f7d8d5',
        fontSize: 40,
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


export default Alert