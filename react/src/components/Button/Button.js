import React from 'react'
import PropTypes from 'prop-types'
import { Button as PaperButton, useTheme } from 'react-native-paper'
import { styles } from './Button.styles'

const Button = ({ loading=false, loadingMessage='', children, ...rest }) => {
    const { colors } = useTheme()
    return(
        <PaperButton 
            style={styles(colors).button} 
            mode='contained' 
            loading={loading} 
            disabled={loading} 
            labelStyle={styles(colors).label}
            {...rest}
        >
            { loading ? loadingMessage : children }
        </PaperButton>
    )
}

Button.propTypes = {
    loading: PropTypes.bool,
    loadingMessage: PropTypes.string,
    children: PropTypes.any
} 

export default Button