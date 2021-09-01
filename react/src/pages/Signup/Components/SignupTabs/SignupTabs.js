import React from 'react'
import { StyleSheet, View } from 'react-native'
/* Custom components. */
import UserForm from '../UserForm'
import StudentForm from '../StudentForm'

const SignupTabs = ({ page, setPage, user, setUser, doSignup }) => {
    const tabs = [UserForm, StudentForm]
    return (
        <View style={styles.container}>
            {tabs && tabs.map(( Tab, i ) => {
                if( page == i ){
                    return(
                        <Tab 
                            page={page} 
                            setPage={setPage} 
                            user={user}
                            setUser={setUser}
                            doSignup={doSignup}
                        />    
                    )
                }
            })}
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        width: '100%'
    }
})

export default SignupTabs
