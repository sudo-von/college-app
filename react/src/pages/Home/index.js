import React, { useEffect, useState } from 'react'
import { StyleSheet, View } from 'react-native'
/* Custom components. */
import Container from 'src/components/Container'
import Small from 'src/components/Small'
import NavigationBadge from 'src/components/NavigationBadge'
import Mood from 'src/components/Mood'
/* React native paper. */
import { Title } from 'react-native-paper'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'
/* Services. */
import { getMood } from 'src/services/mood.service'

const Home = () => {
    
    const [ modal, setModal ] = useState(false)

    const { authState } = useAuth()
    const { user } = authState
    const { user_name, user_id } = user

    useEffect(() => {
        
        const getDailyMood = async () => {
            try{
                await getMood(user_id)
            }catch(error){
            }
        }
        getDailyMood()
    }, [])

    return (
        <Container style={styles.container}>
            <Mood/>
            <Title>¡Hola {user_name}!</Title>
            <Small>Mantente al día con las últimas noticias visitando la página de tu escuela.</Small>
            <NavigationBadge
                url='/'
                title='Botón de pánico'
                icon='alert-circle-outline'
                color='#611a15'
                backgroundColor='#f7d8d5'
            />
            <NavigationBadge
                url='/'
                title='Asesorías'
                icon='book-open-variant'
                color='#4F89F8'
                backgroundColor='#e8f4fd'
            />
            <NavigationBadge
                url='/'
                title='Sugerencias'
                icon='email-outline'
                color='#F8DD4F'
                backgroundColor='#FEFFB7'
            />
        </Container>
    )
}

const styles = StyleSheet.create({
    container: {
        justifyContent: 'flex-start'
    }
})

export default Home
