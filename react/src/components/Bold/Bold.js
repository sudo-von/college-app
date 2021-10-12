import React from 'react'
import { Text } from 'react-native'
import { styles } from './Bold.styles'

const Bold = ({ children, ...rest }) =>
    <Text style={styles.text} {...rest}>
        {children}
    </Text>

export default Bold