import React from 'react'
import PropTypes from 'prop-types'
import { View } from 'react-native'
import { styles } from './Center.styles'

const Center = ({ children, ...rest }) => {
    return (
        <View style={styles.view} {...rest}>
            {children}
        </View>
    )
}

Center.propTypes = {
    children: PropTypes.any
} 


export default Center