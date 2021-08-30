import React, {createContext, useReducer, useContext} from 'react'
/* Reducers. */
import authReducer from '../reducers/auth.reducer'

const AuthContext = createContext()

const useAuth = () => useContext(AuthContext)

const AuthProvider = ({ children }) => {

    const [authState, authDispatch] = useReducer(authReducer, {
        isLoggedIn: false,
        user: null
    })

    return (
        <AuthContext.Provider value={{authState, authDispatch}}>
            {children}
        </AuthContext.Provider>
    )
}

export { AuthProvider, useAuth }