import React from 'react'
import PropTypes from 'prop-types'
import { View, TouchableHighlight } from 'react-native'
import { Caption, IconButton } from 'react-native-paper'
import { useNavigation } from '@react-navigation/native'
import { styles } from './NavigationBadge.styles'

const NavigationBadge = ({ url='/', title='Inicio', icon='home', color='white', backgroundColor='gray' }) => {
    
    const navigation = useNavigation()
    const handleNavigation = () => navigation.navigate(url)

    return (
        <TouchableHighlight onPress={handleNavigation} underlayColor='white'>
            <View style={styles(backgroundColor).view}>
                <IconButton icon={icon} color={color}/>
                <Caption>{title}</Caption>
            </View>
        </TouchableHighlight>
    )
}

NavigationBadge.propTypes = {
    url: PropTypes.string,
    title: PropTypes.string,
    icon: PropTypes.string,
    color: PropTypes.string,
    backgroundColor: PropTypes.string
}

export default NavigationBadge