import React from 'react'
import PropTypes from 'prop-types'
import { Text } from 'react-native'
import { styles } from './Bold.styles'

const Bold = ({ children, ...rest }) =>
    <Text style={styles.text} {...rest}>
        {children}
    </Text>

Bold.propTypes = {
    children: PropTypes.any
} 

export default Bold