import { StyleSheet } from 'react-native'

export const styles = ({ justifyContent, style}) => StyleSheet.create({
    container: {
        flex: 1,
        padding: 25,
        justifyContent,
        ...style
    }
})