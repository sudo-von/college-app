import React from 'react'
import { StyleSheet, View } from 'react-native'
/* React native paper. */
import { TextInput, HelperText } from 'react-native-paper'

const Input = ({ error, ...rest }) =>
    <View>
        <TextInput
            mode='outlined'
            style={styles.input}
            theme={{ colors: { primary: '#2d798a' }}}
            {...rest}
        />
        { error && 
            <HelperText type="error" visible={true} style={styles.helperText}>
                {error}
            </HelperText>
        }
    </View>


 
const styles = StyleSheet.create({
    helperText: {
        padding: 0,
        margin: 0,
    },
    input: {
        marginTop: 5,
        backgroundColor: 'white'
    }
})


export default Input
