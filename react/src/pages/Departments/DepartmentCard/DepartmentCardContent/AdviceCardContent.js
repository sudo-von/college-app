import React from 'react'
import { View } from 'react-native'
import { Bold, Small } from 'src/components'
import { Card, Title, Paragraph, Badge } from 'react-native-paper'
import { styles } from './AdiveCardContent.styles'

const DepartmentCardContent = ({ available, cost, description, departmentUser, neighborhood, street }) => {
    
    const formatedCost = '$'+cost.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',')

    return (
        <Card.Content style={styles({}).content}>
            <View style={styles({}).view}>
                <Small>{departmentUser.name} / {departmentUser.email}</Small>
                <Badge size={20} style={styles({ available }).badge}>{ available ? 'Disponible' : 'No disponible' }</Badge>
            </View>
                <Paragraph style={styles({}).paragraph}><Bold>{street} / {neighborhood}</Bold></Paragraph>
            <Title style={styles({}).cost}><Bold>{formatedCost}</Bold></Title> 
            <Small>{description}</Small>
        </Card.Content>
    )
}

export default DepartmentCardContent
