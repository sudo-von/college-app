import { createRef } from 'react'

const navigationRef = createRef(null)

const navigate = (name) => {
    if (navigationRef.current) {
        navigationRef.current.navigate(name)
    }
}

export { 
    navigationRef,
    navigate
}