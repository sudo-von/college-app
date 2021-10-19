import React, { useState } from 'react'
/* React native paper. */
import { TextInput } from 'react-native-paper'
/* Custom components. */
import Input from 'src/components/Input/Input'

const PasswordInput = ({ ...rest }) => {
    const [ secureTextEntry, seSecureTextEntry ] = useState(true)
    return (
        <Input
            secureTextEntry={secureTextEntry ? true : false}
            right={
                <TextInput.Icon 
                    onPress={() => seSecureTextEntry(!secureTextEntry)} 
                    color='gray' 
                    name={secureTextEntry ? 'eye-off' : 'eye'}
                />
            }
            {...rest}
        />
    )
}

export default PasswordInput
