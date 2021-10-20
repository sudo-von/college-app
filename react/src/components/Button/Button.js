import React from 'react'
import PropTypes from 'prop-types'
import { FAB, useTheme } from 'react-native-paper'
import { styles } from './Button.styles'

const Button = ({ loading=false, loadingMessage='', children, margin, style, ...rest }) => {
    const { colors } = useTheme()
    return(
        <FAB 
            style={{...styles(colors).button,...style}} 
            label={loading ? loadingMessage : children}
            loading={loading} 
            color='white'
            disabled={loading} 
            {...rest}
        />
    )
}

Button.propTypes = {
    loading: PropTypes.bool,
    loadingMessage: PropTypes.string,
    children: PropTypes.any
} 

export default Button