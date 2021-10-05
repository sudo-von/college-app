import { StyleSheet } from 'react-native'

export const styles = (colors) => StyleSheet.create({
    container: {
        padding: 0
    },
    view: {
        backgroundColor: '#F2F2F2',
        flex: 1,
    },
    bottomView: {
        alignItems: 'center',
        backgroundColor: colors.primary
    }
})