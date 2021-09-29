const authReducer = (state, action) => {
    switch (action.type) {
        case 'login': {
            return { 
                isLoggedIn: true, 
                user : action.user
            }
        }
        case 'logout': {
            return {
                isLoggedIn: false,
                user: null
            }
        }
        case 'update': {
            let { user } = action
            return {...state, user : {...state.user, "user_name": user.name}}
        }
        default:
            return state
    }
}

export default authReducer