import React from 'react'
import { StyleSheet } from 'react-native'
/* React native paper. */
import { Button } from 'react-native-paper'
import { useTheme } from '@react-navigation/native'

const CustomButton = ({ loading, loadingMessage, children, ...rest}) => {
    const { colors } = useTheme()
    return(
        <Button 
            style={styles(colors).button} 
            mode='contained' 
            loading={loading} 
            disabled={loading} 
            labelStyle={styles(colors).label}
            {...rest}
        >
            { loading ? loadingMessage : children }
        </Button>
    )
}

const styles = (colors) => StyleSheet.create({
    label: {
        color: 'white'
    },
    button: {
        backgroundColor: colors.primary,
    }
})

export default CustomButton