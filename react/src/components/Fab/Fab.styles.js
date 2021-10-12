import { StyleSheet } from 'react-native'

export const styles = ({ primary}) => StyleSheet.create({
    fab: {
        color: 'white',
        backgroundColor: primary,
        position: 'absolute',
        margin: 15,
        right: 0,
        top: 0
    }
})