import React from 'react'
import { StyleSheet } from 'react-native'
/* React navigation. */
import { Link } from '@react-navigation/native'
/* React native paper. */
import { useTheme, Caption } from 'react-native-paper'

const CustomLink = ({ url, children }) => {
    const { colors } = useTheme()
    return(
        <Link to={{ screen: url }}>
            <Caption style={styles({ colors }).signup}>
                {children}
            </Caption>
        </Link>
    )
}

const styles = ({ colors }) => StyleSheet.create({
    signup: {
        fontSize: 16,
        color: colors.primary
    }
})

export default CustomLink
