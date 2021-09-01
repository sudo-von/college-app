import React from 'react'
import { View } from 'react-native'
/* Formik. */
import { Formik } from 'formik'
/* Custom components. */
import Button from 'src/components/Button'
import Input from 'src/components/Input'
import Small from 'src/components/Small'

const UserForm = ({ page, setPage, user, setUser}) => {

    const onSubmit = (data) => {
        const { name, birthDate } = data
        setUser({
            ...user, 
            'name': name, 
            'birth_date' : birthDate
        })
        setPage(page+1)
    }

    return (
        <Formik 
            initialValues={{ name: user.name, birthDate: user.birth_date}}
            validate={values => {
                const { name, birthDate } = values
                const errors = {}
                if(!name){
                    errors.name = 'Nombre requerido'
                }
                if(!birthDate){
                    errors.birthDate = 'Fecha de nacimiento requerida'
                }
                return errors
            }}
            onSubmit={onSubmit}
        >
            {({ handleChange, handleBlur, handleSubmit, errors, values }) => (
                <View>
                    <Small>PASO {page+1} DE 2</Small>
                    <Small>Cuéntanos sobre ti</Small>
                    <Input
                        label='Ingresa tu nombre'
                        onChangeText={handleChange('name')}
                        onBlur={handleBlur('name')}
                        value={values.name}
                        error={errors.name}
                    />
                    <Input
                        label='Selecciona tu fecha de nacimiento'
                        onChangeText={handleChange('birthDate')}
                        onBlur={handleBlur('birthDate')}
                        value={values.birthDate}
                        error={errors.birthDate}
                    />
                    <Button onPress={handleSubmit}>Siguiente</Button>
                </View>
            )}
        </Formik>
    )
}

export default UserForm
