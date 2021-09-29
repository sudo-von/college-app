import React from 'react'
import { StyleSheet, View } from 'react-native'
/* Custom components. */
import Container from 'src/components/Container'
import Small from 'src/components/Small'
import Bold from 'src/components/Bold'
import NavigationBadge from 'src/components/NavigationBadge'
import Mood from './Components/Mood'
/* React native paper. */
import { Title } from 'react-native-paper'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'

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
        url: '/suggestions',
        title: 'Sugerencias',
        icon: 'email-outline',
        color: '#F8DD4F',
        backgroundColor: '#FEFFB7'
    }
]

const Home = () => {

    /* Destructuring user properties. */
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
                <Title><Bold>¡Hola {user_name}!</Bold></Title>
                <Small>Mantente al día con las últimas{'\n'}noticias visitando la página de tu escuela.</Small>
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

const styles = StyleSheet.create({
    view: {
        marginVertical: 40,
    },
    title: {
        fontSize: 24
    }
})

export default Home
