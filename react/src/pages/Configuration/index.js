import React from 'react'
import { StyleSheet, View } from 'react-native'
/* Custom components. */
import Container from 'src/components/Container'
import Small from 'src/components/Small'
import Bold from 'src/components/Bold'
import NavigationBadge from 'src/components/NavigationBadge'
/* React native paper. */
import { Title } from 'react-native-paper'

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

const Configuration = () => 
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

const styles = StyleSheet.create({
    view: {
        marginVertical: 40,
    }
})

export default Configuration