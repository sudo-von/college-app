import { StyleSheet } from 'react-native'

export const styles = (backgroundColor) => StyleSheet.create({
    view: {
        marginTop: 10,
        flexDirection: 'row',
        alignItems: 'center',
        width: '100%',
        borderRadius: 50,
        backgroundColor
    }
})