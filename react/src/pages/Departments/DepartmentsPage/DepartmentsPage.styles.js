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
        backgroundColor: background,
        color: primary
    },
        bottomView: {
        alignItems: 'center',
        backgroundColor: background
    }
})