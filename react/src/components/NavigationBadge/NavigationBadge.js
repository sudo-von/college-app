import React from 'react'
import { View, TouchableHighlight } from 'react-native'
import { Caption, IconButton } from 'react-native-paper'
import { useNavigation } from '@react-navigation/native'
import { styles } from './NavigationBadge.styles'

const NavigationBadge = ({ url, title, icon, color, backgroundColor }) => {
    
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

export default NavigationBadge