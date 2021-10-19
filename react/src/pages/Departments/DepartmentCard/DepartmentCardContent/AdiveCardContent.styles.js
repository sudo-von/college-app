import { StyleSheet } from 'react-native'

export const styles = (available) => StyleSheet.create({
    badge: {
        backgroundColor: available ? 'green' : 'red'
    },
    content: {
      padding: 8,
    },
    cost: {
        fontSize: 20
    },
    paragraph: {
        color: 'gray'
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