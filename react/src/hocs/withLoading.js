import React from 'react'
/* Custom components. */
import Loader from 'src/components/Loader'

const withLoading = (Component, loadingMessage) => {

    const WithLoadingComponent = ({ isLoading, ...props }) => {

        if (isLoading){
            return (
                <Loader loadingMessage={loadingMessage}/>
            )
        }
        return <Component {...props}/>
    }

    return WithLoadingComponent
}

export default withLoading