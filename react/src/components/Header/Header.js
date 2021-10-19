import React from 'react'
import PropTypes from 'prop-types'
import { View } from 'react-native'
import Bold from 'src/components/Bold/Bold'
import { Title, Caption } from 'react-native-paper'
import { styles } from './Header.styles.js'

const Header = ({ title='', subtitle='', ...rest }) => {
    return (
        <View style={styles.container} {...rest}>
            <Title><Bold>{title}</Bold></Title>
            <Caption style={styles.subtitle}>{subtitle}</Caption>
        </View>
    )
}

Header.propTypes = {
    title: PropTypes.string,
    subtitle: PropTypes.string
}

export default Header