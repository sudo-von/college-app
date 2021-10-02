import React, { useEffect, useState } from 'react'
import { View, Alert, StyleSheet } from 'react-native'
/* Formik. */
import { Formik, Field } from 'formik'
/* Custom components. */
import Input from 'src/components/Input'
import SelectInput from 'src/components/SelectInput'
import Datepicker from 'src/components/Datepicker'
import Button from 'src/components/Button'
/* Services. */
import { createAdvice } from 'src/services/advice.service'
import { getUniversityByID } from 'src/services/university.service'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'

const CreateAdviceForm = () => {

    /* Gets user id. */
    const { authState } = useAuth()
    const { user } = authState
    const { university_id } = user
    /* Handles classrooms list. */
    const [ classrooms, setClassrooms ] = useState([])
    useEffect(() => {
        const searchUniversityByID = async () => {
            try{
                const university = await getUniversityByID(university_id)
                console.log(university)
                setClassrooms(university.classrooms.map(classroom => ({ label: classroom.name, value: classroom.id })))
            }catch(error){
                Alert.alert('¡Ha ocurrido un error!', error.message)
            }
        }
        searchUniversityByID()
    },[])

    /* Handles form submit. */
    const [ loading, setLoading ] = useState(false)
    const onSubmit = async (form, { resetForm }) => {
        try{
            setLoading(true)
            await createAdvice(form)
            resetForm()
            Alert.alert(response, '¡Has registrado la asesorái con éxito!')
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }finally{
            setLoading(false)
        }
    }

    return (
        <Formik 
            initialValues={{ 
                subject: '',
                advice_date: '',
                classroom_id : ''
            }}
            validate={values => {
                const { subject, advice_date, classroom_id } = values
                const errors = {}
                if(!subject){
                    errors.subject = 'Materia requerida'
                }
                if(!advice_date){
                    errors.advice_date = 'Fecha de la asesoría requerida'
                }
                if(!classroom_id){
                    errors.classroom_id = 'Salón requerido'
                }
                return errors
            }}
            onSubmit={onSubmit}
        >
            {({ handleChange, handleBlur, handleSubmit, errors, values }) => (
                <View>
                    <Input
                        label='Ingresa la materia'
                        onChangeText={handleChange('subject')}
                        onBlur={handleBlur('subject')}
                        value={values.subject}
                        error={errors.subject}
                    />
                    <Field name='advice_date'>
                        {({ field, form, meta }) => (
                            <Datepicker 
                                label='Selecciona la fecha de la asesoría'
                                field={field} 
                                form={form} 
                                meta={meta}
                            />
                        )}
                    </Field>
                    <Field name='classroom_id'>
                        {({ field, form, meta }) => (
                            <SelectInput
                                label='Selecciona el salón'
                                data={classrooms}
                                field={field} 
                                form={form} 
                                meta={meta}
                                error={errors.classroom_id}
                            />
                        )}
                    </Field>
                    <Button 
                        loading={loading} 
                        loadingMessage='Registrando asesoría...' 
                        onPress={handleSubmit}
                        style={styles.button}
                    >
                        Registrar asesoría
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

export default CreateAdviceForm
