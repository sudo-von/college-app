import React, { useEffect, useState } from 'react'
import { View, Alert } from 'react-native'
import { Formik, Field } from 'formik'
import { Input, PasswordInput, SelectInput, Datepicker, Option, Button, Loader } from 'src/components'
import { signup } from 'src/services/user.service'
import { getUniversities } from 'src/services/university.service'
import { useNavigation } from '@react-navigation/native'
import { styles } from './SignupForm.styles'

const SignupForm = () => {

    /* Handles form. */
    const [ formLoading, setFormLoading ] = useState(false)
    /* Handles universities. */
    const navigation = useNavigation()
    const [ universities, setUniversities ] = useState([])
    const [ loadingUniversities, setLoadingUniversities ] = useState(false)

    /* Formik configuration. */
    const initialValues = { 
        name: 'VoN',
        birth_date: '1997-04-17',
        email : 'martinez-angel@uadec.edu.mx', 
        registration_number: '16190775',
        university_id: '', 
        password: 'college-app'
    }

    const handleValidation = ({ name, birth_date, email, registration_number, university_id, password }) => {
        let errors = {}
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
    }

    const onSubmit = async (form, { resetForm }) => {
        try{
            setFormLoading(true)
            const response = await signup(form)
            setFormLoading(false)
            resetForm()
            Alert.alert('¡Felicidades!', 'Ya puedes iniciar sesión.')
            navigation.navigate('/login')
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
            setFormLoading(false)
        }
    }

    /* Get universities for the select component. */
    useEffect(() => {
        const searchUniversities = async () => {
            try{
                setLoadingUniversities(true)
                const response = await getUniversities()
                /* Creates a new array with the special structure that the SelectInput needs. */
                const universities = response.results.map((university) => (
                    { 
                        label : university.name, 
                        value : university.id, 
                        custom: <Option source={university.profile_picture} label={university.name}/>
                    }
                ))
                setUniversities(universities)
            }catch(error){
                Alert.alert('¡Ha ocurrido un error!', error.message)
            }finally{
                setLoadingUniversities(false)
            }
        }
        searchUniversities()
    }, [])

    return (
        <Formik 
            initialValues={initialValues}
            validate={handleValidation}
            validateOnChange={false}
            validateOnBlur={false} 
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
                                label={loadingUniversities ? 'Cargando universidades...' : 'Selecciona tu universidad'}
                                data={universities}
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
                        loading={formLoading} 
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

export default SignupForm
