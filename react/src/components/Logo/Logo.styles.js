import { StyleSheet } from 'react-native'

export const styles = ({ size, colors }) => StyleSheet.create({ 
    view: {
        borderRadius: 100,
        width: size,
        height: size,
        backgroundColor: colors.primary,
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center'
    },
    image: {
        width: size-15,
        resizeMode: 'contain'
    }
})