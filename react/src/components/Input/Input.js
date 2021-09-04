import React from 'react'
import { StyleSheet, View } from 'react-native'
/* React native paper. */
import { TextInput, HelperText } from 'react-native-paper'
/* React navigation. */
import { useTheme } from '@react-navigation/native'

const Input = ({ error, ...rest }) => {
    const { colors } = useTheme()
    return(
        <View>
            <TextInput
                mode='outlined'
                style={styles.input}
                theme={{ colors: { primary: colors.primary }}}
                {...rest}
            />
            { error && 
                <HelperText 
                    padding='none' 
                    type='error' 
                    visible={true}
                >
                    {error}
                </HelperText>
            }
        </View>
    )
}


 
const styles = StyleSheet.create({
    input: {
        marginTop: 5,
    }
})


export default Input
