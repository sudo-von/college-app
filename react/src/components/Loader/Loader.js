import React from 'react'
/* Custom components. */
import Container from 'src/components/Container'
import Center from 'src/components/Center'
import Small from 'src/components/Small'
/* React native paper. */
import { ActivityIndicator } from 'react-native-paper'

const Loader = ({ loadingMessage, size }) => {
    return (
        <Container>
            <ActivityIndicator size={size ? size : 80}/>
            <Center>
                <Small>{loadingMessage}</Small>
            </Center>
        </Container>
    )
}

export default Loader
