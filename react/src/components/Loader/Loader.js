import React from 'react'
import PropTypes from 'prop-types'
import Container from 'src/components/Container/Container'
import Center from 'src/components/Center/Center'
import Small from 'src/components/Small/Small'
import { ActivityIndicator } from 'react-native-paper'

const Loader = ({ loadingMessage='Cargando...', size=80 }) => {
    return (
        <Container>
            <ActivityIndicator size={size}/>
            <Center>
                <Small>{loadingMessage}</Small>
            </Center>
        </Container>
    )
}

Loader.propTypes = {
    loadingMessage: PropTypes.string,
    size: PropTypes.number
}

export default Loader
