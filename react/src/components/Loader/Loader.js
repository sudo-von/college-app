import React from 'react'
import PropTypes from 'prop-types'
import { Container, Center, Small }  from 'src/components'
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
