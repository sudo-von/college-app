import { StyleSheet } from 'react-native'

export const styles = ({ primary }) => StyleSheet.create({
    button: {
        backgroundColor: primary,
        borderRadius: 50
    },
    label: {
        color: 'white'
    }
})