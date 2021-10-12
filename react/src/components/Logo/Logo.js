import React from 'react'
import PropTypes from 'prop-types'
import { View, Image } from 'react-native'
import { useTheme } from 'react-native-paper'
import { styles } from './Logo.styles'

const Logo = ({ source, size=180 }) => {
    const { colors } = useTheme()
    return (
        <View style={styles({ size, colors}).view}>
            <Image
                style={styles({ size, colors}).image}
                source={source}
            />
        </View>
    )
}

Logo.propTypes = {
    source: PropTypes.any,
    size: PropTypes.number
}

export default Logo
