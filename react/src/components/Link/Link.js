import React from 'react'
import { StyleSheet } from 'react-native'
/* React router native. */
import { Link } from 'react-router-native'
/* React native paper. */
import { Caption } from 'react-native-paper'

const CustomLink = ({ url, children }) =>
    <Link
        to={url}
        style={styles.link}
        underlayColor="#f0f4f7"
    >
        <Caption style={styles.signup}>{children}</Caption>
    </Link>

const styles = StyleSheet.create({
    link: {
        color: 'red'
    },
    signup: {
        fontSize: 14,
        color: '#4C9DAF',
        textAlign: 'center',
        lineHeight: 13
    }
})

export default CustomLink
