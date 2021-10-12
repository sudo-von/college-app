import { StyleSheet } from 'react-native'

export const styles = (rest) => StyleSheet.create({
    caption: {
        color: '#666666',
        fontSize: 12,
        ...rest
    }
})