import React, { useState } from 'react'
import { StyleSheet, View } from 'react-native'
/* Custom components. */
import Logo from 'src/components/Logo'
import Alert from 'src/components/Alert'
/* React native paper. */
import SignupTabs from './Components/SignupTabs'
/* Services. */
import { signup } from 'src/services/user.service'

const Signup = () => {

    const logo = require('../../../assets/gopher.png')
    /* Handle current page. */
    const [page, setPage ] = useState(0)
    /* Handle both forms data. */
    const [user, setUser] = useState({
        'name' : 'asd',
        'birth_date' : '1997-04-17',
        'email' : 'sudovonctf@gmail.com',
        'registration_number' : '16190774',
        'university_id' : '612803c400e131d7b8163642',
        'password' : '123456',
    })
    /* States. */
    const [ loading, setLoading ] = useState(false)
    const [ error, setError ] = useState(null) 
    const [ success, setSuccess ] = useState(null) 

    const doSignup = async () => {
        try{
            setLoading(true)
            setError(null)
            setSuccess(null)
            console.log(user)
            const response = await signup(user)
            setSuccess(response)
        }catch(error){
            console.log(error)
            setError(error.message)
        }finally{
            setLoading(false)
        }
    }

    return (
        <View style={styles.container}>
            <Logo 
                source={logo} 
                style={styles.logo}
            />
            { error && <Alert title={error} type='error'/>}
            { success && <Alert title={success} type='success'/>}
            <SignupTabs 
                page={page} 
                setPage={setPage} 
                user={user}
                setUser={setUser}
                doSignup={doSignup}
            />
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        padding: 25,
        flex: 1,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center'
    },
    logo: {
        width: 125,
        height: 125
    }
})

export default Signup