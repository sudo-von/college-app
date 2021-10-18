import React from 'react'
import { View, Image } from 'react-native'
import { Container, Small, Bold, NavigationBadge } from 'src/components'
import { Title } from 'react-native-paper'
import { useAuth } from 'src/providers/auth.provider'
import { styles } from './HomePage.styles'
import Mood from '../Mood/Mood'

const routes = [
    {
        url: '/panic-button',
        title: 'Botón de pánico',
        icon: 'alert-circle-outline',
        color: '#611a15',
        backgroundColor: '#f7d8d5'
    },
    {
        url: '/advices',
        title: 'Asesorías',
        icon: 'book-open-variant',
        color: '#4F89F8',
        backgroundColor: '#e8f4fd'
    },
    {
        url: '/departments',
        title: 'Departamentos',
        icon: 'home-city-outline',
        color: '#49E286',
        backgroundColor: '#e6ffe8'
    },
    {
        url: '/suggestions',
        title: 'Sugerencias',
        icon: 'email-outline',
        color: '#F8DD4F',
        backgroundColor: '#FEFFB7'
    }
]

const HomePage = () => {

    const { authState } = useAuth()
    const { user } = authState
    const { user_name, user_id } = user

    return (
        <Container justifyContent='flex-start'>
            <Mood 
                initialMoodValue={5}
                minimumValue={0}
                maximumValue={10}
                minimumText='Triste'
                maximumText='Feliz'
                userID={user_id}
            />
            <View style={styles.view}>
                <Image style={styles.image} source={require('assets/home.png')}/>
                <Title style={styles.title}><Bold>¡Hola {user_name}!</Bold></Title>
                <Small style={styles.small}>Mantente al día con las últimas{'\n'}noticias visitando la página de tu escuela.</Small>
            </View>
            { routes && routes.map(({ url, title, icon, color, backgroundColor, ...rest }, index) =>
                <NavigationBadge
                    key={`${url}-${title}-${index}`}
                    url={url}
                    title={title}
                    icon={icon}
                    color={color}
                    backgroundColor={backgroundColor}
                    {...rest}
                />
            )}
        </Container>
    )
}

export default HomePage
