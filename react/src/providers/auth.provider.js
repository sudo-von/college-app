import React, {createContext, useReducer, useContext} from 'react'
/* Reducers. */
import authReducer from '../reducers/auth.reducer'

const AuthContext = createContext()

const useAuth = () => useContext(AuthContext)

const AuthProvider = ({ children }) => {

    const [state, dispatch] = useReducer(authReducer, {
        isLoggedIn: false,
        user: null
    })

    return (
        <AuthContext.Provider value={{state, dispatch}}>
            {children}
        </AuthContext.Provider>
    )
}

export { AuthProvider, useAuth }