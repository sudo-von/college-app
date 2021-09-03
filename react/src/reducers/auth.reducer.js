const authReducer = (state, action) => {
    switch (action.type) {
        case 'login': {
            let { user } = action
            return {...state, isLoggedIn: true, user}
        }
        case 'logout': {
            return {
                isLoggedIn: false,
                user: null
            }
        }
        default:
            return state
    }
}

export default authReducer