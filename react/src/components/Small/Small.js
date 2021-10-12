import React from 'react'
import PropTypes from 'prop-types'
import { Caption } from 'react-native-paper'
import { styles } from './Small.styles'

const Small = ({ children, style={}, ...rest }) =>
    <Caption 
        style={styles(style).caption} 
        {...rest}
    >
        {children}
    </Caption>

Small.PropTypes = {
    children: PropTypes.any,
    styles: PropTypes.object
}

export default Small