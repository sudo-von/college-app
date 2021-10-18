import { StyleSheet } from 'react-native'

export const styles = ({ primary, available }) => StyleSheet.create({
    badge: {
        backgroundColor: available ? 'green' : 'red'
    },
    content: {
      padding: 10,
    },
    cost: {
        fontSize: 18
    },
    title:{
        fontSize: 18,
    },
    paragraph: {
        color: '#36454F'
    },
    small: {
        color: 'gray',
    },
    view: {
        flexDirection: 'row', 
        justifyContent: 'space-between', 
        alignItems: 'center'
    }
})