import React, { useState } from 'react'
/* React native paper. */
import { TextInput } from 'react-native-paper'
/* Custom components. */
import Input from 'src/components/Input'

const PasswordInput = ({ ...rest }) => {
    const [ showPassword, setShowPassword ] = useState(false)
    return (
        <Input
            secureTextEntry={showPassword ? true : false}
            right={
                <TextInput.Icon 
                    onPress={() => setShowPassword(!showPassword)} 
                    color='gray' 
                    name={showPassword ? 'eye-off' : 'eye'}
                />
            }
            {...rest}
        />
    )
}

export default PasswordInput
