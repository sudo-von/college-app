import React from 'react'
import PropTypes from 'prop-types'
import { View } from 'react-native'
import { styles } from './Container.styles'

const Container = ({ children, justifyContent='center', style={}, ...rest}) => 
    <View style={styles({ justifyContent, style}).container} {...rest}>
        {children}
    </View>

Container.propTypes = {
    children: PropTypes.any,
    justifyContent: PropTypes.string,
    style: PropTypes.object,
} 

export default Container