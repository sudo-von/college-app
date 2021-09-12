import React, { useEffect, useState } from 'react'
import { View, Alert, StyleSheet } from 'react-native'
/* Formik. */
import { Formik, Field } from 'formik'
/* Custom components. */
import Input from 'src/components/Input'
import PasswordInput from 'src/components/PasswordInput'
import SelectInput from 'src/components/SelectInput'
import Datepicker from 'src/components/Datepicker'
import Option from 'src/components/Option'
import Button from 'src/components/Button'
/* Services. */
import { signup } from 'src/services/user.service'
import { getUniversities } from 'src/services/university.service'
/* React navigation. */
import { useNavigation } from '@react-navigation/native'

const SignupForm = () => {

    /* Handles university list. */
    const [ universityList, setUniversityList ] = useState([])
    useEffect(() => {
        const searchUniversities = async () => {
            try{
                const response = await getUniversities()
                /* Creates a new array with the special structure that the SelectInput needs. */
                const universities = response.map((university) => (
                    { 
                        label : university.name, 
                        value : university.id, 
                        custom: <Option source={university.profile_picture} label={university.name}/>
                    }
                ))
                setUniversityList(universities)
            }catch(error){
                Alert.alert('¡Ha ocurrido un error!', error.message)
            }
        }
        searchUniversities()
    },[])

    /* Handles form submit. */
    const navigation = useNavigation()
    const [ loading, setLoading ] = useState(false)
    const onSubmit = async (form, { resetForm }) => {
        try{
            setLoading(true)
            const response = await signup(form)
            setLoading(false)
            resetForm()
            Alert.alert(response, 'Ya puedes iniciar sesión.')
            navigation.navigate('/login')
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
            setLoading(false)
        }
    }

    return (
        <Formik 
            initialValues={{ 
                name: 'VoN',
                birth_date: '1997-04-17',
                email : 'martinez-angel@uadec.edu.mx', 
                registration_number: '16190775',
                university_id: '', 
                password: 'college-app'
            }}
            validate={values => {
                const { name, birth_date, email, registration_number, university_id, password } = values
                const errors = {}
                if(!name){
                    errors.name = 'Nombre requerido'
                }
                if(!birth_date){
                    errors.birth_date = 'Fecha de nacimiento requerida'
                }
                if(!email){
                    errors.email = 'Correo requerido'
                }
                if(!registration_number){
                    errors.registration_number = 'Matrícula requerida'
                }
                if(!university_id){
                    errors.university_id = 'Universidad requerida'
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
                    <Input
                        label='Ingresa tu nombre'
                        onChangeText={handleChange('name')}
                        onBlur={handleBlur('name')}
                        value={values.name}
                        error={errors.name}
                    />
                    <Field name='birth_date'>
                        {({ field, form, meta }) => (
                            <Datepicker 
                                label='Selecciona tu fecha de nacimiento'
                                field={field} 
                                form={form} 
                                meta={meta}
                            />
                        )}
                    </Field>
                    <Input
                        label='Ingresa tu correo institucional'
                        onChangeText={handleChange('email')}
                        onBlur={handleBlur('email')}
                        value={values.email}
                        error={errors.email}
                    />
                    <Input
                        label='Ingresa tu matrícula'
                        onChangeText={handleChange('registration_number')}
                        onBlur={handleBlur('registration_number')}
                        value={values.registration_number}
                        error={errors.registration_number}
                        keyboardType='phone-pad'
                    />
                    <Field name='university_id'>
                        {({ field, form, meta }) => (
                            <SelectInput
                                label='Selecciona tu facultad'
                                data={universityList}
                                field={field} 
                                form={form} 
                                meta={meta}
                                error={errors.university_id}
                            />
                        )}
                    </Field>
                    <PasswordInput
                        label='Ingresa tu contraseña'
                        onChangeText={handleChange('password')}
                        onBlur={handleBlur('password')}
                        value={values.password}
                        error={errors.password}
                    />
                    <Button 
                        loading={loading} 
                        loadingMessage='Registrando usuario...' 
                        onPress={handleSubmit}
                        style={styles.button}
                    >
                        Registrarse
                    </Button>
                </View>
            )}
        </Formik>
    )
}

const styles = StyleSheet.create({
    button: {
        marginVertical: 20
    }
})

export default SignupForm
