import React from 'react'
import { View } from 'react-native'
import { Container, Small, Bold } from 'src/components'
import { Title } from 'react-native-paper'
import { styles } from './ConfigurationPage.styles'
import NavigationBadge from 'src/components/NavigationBadge'

const routes = [
    {
        url: '/account-configuration',
        title: 'Configuración de la cuenta',
        icon: 'account-cog-outline',
        color: '#498FE2',
        backgroundColor: '#e6f3ff'
    },
    {
        url: '/contact-configuration',
        title: 'Configuración de contacto',
        icon: 'whatsapp',
        color: '#49E286',
        backgroundColor: '#e6ffe8'
    }
]

const ConfigurationPage = () => 
    <Container justifyContent='flex-start'>
        <View style={styles.view}>
            <Title><Bold>¡Bienvenido al{'\n'}panel de configuración!</Bold></Title>
            <Small>Modifica tu información en cualquier momento.</Small>
        </View>
        { routes && routes.map(({ url, title, icon, color, backgroundColor }, index) =>
            <NavigationBadge
                key={`${url}-${title}-${index}`}
                url={url}
                title={title}
                icon={icon}
                color={color}
                backgroundColor={backgroundColor}
            />
        )}
    </Container>

export default ConfigurationPage
