import React from 'react'
import PropTypes from 'prop-types'
import { Link as NativeLink } from '@react-navigation/native'
import { useTheme, Caption } from 'react-native-paper'
import { styles } from './Link.styles'

const Link = ({ url='/', children, ...rest }) => {
    const { colors } = useTheme()
    return(
        <NativeLink to={{ screen: url }} {...rest}>
            <Caption style={styles(colors).caption}>
                {children}
            </Caption>
        </NativeLink>
    )
}

Link.propTypes = {
    title: PropTypes.string,
    subtitle: PropTypes.string
}

export default Link
