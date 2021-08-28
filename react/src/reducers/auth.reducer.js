const authReducer = (state, action) => {
    switch (action.type) {
        case 'login': {
            let { user } = action
            return {...state, isLoggedIn: true, user}
        }
        default:
            return state
    }
}

export default authReducer