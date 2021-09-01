import React from 'react'
import { View } from 'react-native'
/* Formik. */
import { Formik } from 'formik'
/* Custom components. */
import Button from 'src/components/Button'
import Input from 'src/components/Input'
import PasswordInput from 'src/components/PasswordInput'
import Small from 'src/components/Small'

const StudentForm = ({ page, user, setUser, doSignup}) => {

    const onSubmit = async (data) => {
        const { email, registrationNumber, universityID, password } = data
        setUser({
            ...user, 
            'email' : email,
            'registration_number' : registrationNumber,
            'university_id' : universityID,
            'password' : password
        })
        await doSignup()
    }

    return (
        <Formik 
            initialValues={{ email : user.email, registrationNumber: user.registration_number, universityID: user.university_id, password: user.password}}
            validate={values => {
                const { email, registrationNumber, universityID, password } = values
                const errors = {}
                if(!email){
                    errors.email = 'Correo requerido'
                }
                if(!registrationNumber){
                    errors.registrationNumber = 'Matrícula requerida'
                }
                if(!universityID){
                    errors.universityID = 'Universidad requerida'
                }
                if(!password){
                    errors.password = 'Contraseña requerida'
                }
                return errors
            }}
            onSubmit={onSubmit}
        >
            {({ handleChange, handleBlur, handleSubmit, errors, values }) => (
                <View>
                    <Small>PASO {page+1} DE 2</Small>
                    <Small>Complementa tus datos uniersitarios</Small>
                    <Input
                        label='Ingresa tu correo institucional'
                        onChangeText={handleChange('email')}
                        onBlur={handleBlur('email')}
                        value={values.email}
                        error={errors.email}
                    />
                    <Input
                        label='Ingresa tu matrícula'
                        onChangeText={handleChange('registrationNumber')}
                        onBlur={handleBlur('registrationNumber')}
                        value={values.registrationNumber}
                        error={errors.registrationNumber}
                    />
                    <Input
                        label='Selecciona tu facultad'
                        onChangeText={handleChange('universityID')}
                        onBlur={handleBlur('universityID')}
                        value={values.universityID}
                        error={errors.universityID}
                    />
                    <PasswordInput
                        label='Ingresa tu contraseña'
                        onChangeText={handleChange('password')}
                        onBlur={handleBlur('password')}
                        value={values.password}
                        error={errors.password}
                    />
                    <Button onPress={handleSubmit}>Registrarse</Button>
                </View>
            )}
        </Formik>
    )
}

export default StudentForm
