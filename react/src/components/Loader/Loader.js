import React from 'react'
/* Custom components. */
import { Container, Center, Small }  from 'src/components'
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
