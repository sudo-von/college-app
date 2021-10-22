import { StyleSheet } from 'react-native'

export const styles = ({ primary, background }) => StyleSheet.create({
    container: {
        padding: 0
    },
    view: {
        backgroundColor: '#F2F2F2',
        flex: 1,
    },
    bottomView: {
        alignItems: 'center',
        backgroundColor: primary
    },
    tabs: {
        backgroundColor: primary,
        color: background
    },
        bottomView: {
        alignItems: 'center',
        backgroundColor: primary
    }
})