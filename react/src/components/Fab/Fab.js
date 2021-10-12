import React from 'react'
import PropTypes from 'prop-types'
import { FAB as PaperFab, useTheme } from 'react-native-paper'
import { styles } from './Fab.styles'

const Fab = ({ icon='home', ...rest}) => {
    const { colors } = useTheme()
    return(
        <PaperFab
            style={styles(colors).fab}
            color={colors.background}
            icon={icon}
            {...rest}
        />
    )
}

Fab.propTypes = {
    icon: PropTypes.string
}

export default Fab