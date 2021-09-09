import React from 'react'
import { StyleSheet } from 'react-native'
/* React native paper. */
import { FAB, useTheme } from 'react-native-paper'

const CustomFab = ({ icon, ...rest}) => {
    const { colors } = useTheme()
    return(
        <FAB 
            style={styles(colors).fab}
            color={colors.background}
            icon={icon}
            {...rest}
        />
    )
}

const styles = (colors) => StyleSheet.create({
    fab: {
        color: 'white',
        backgroundColor: colors.primary,
        position: 'absolute',
        margin: 15,
        right: 0,
        top: 0
    }
})

export default CustomFab