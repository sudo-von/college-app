import React from 'react'
import { StyleSheet } from 'react-native'
/* React navigation. */
import { Link } from '@react-navigation/native'
/* React native paper. */
import { Caption } from 'react-native-paper'

const CustomLink = ({ url, children }) =>
    <Link
        to={{ screen: url }}
        underlayColor="#f0f4f7"
    >
        <Caption style={styles.signup}>{children}</Caption>
    </Link>

const styles = StyleSheet.create({
    signup: {
        fontSize: 14,
        color: '#4C9DAF',
        lineHeight: 13
    }
})

export default CustomLink
