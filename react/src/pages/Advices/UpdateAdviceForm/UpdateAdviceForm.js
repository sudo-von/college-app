import React, { useEffect, useState } from 'react'
import { View, Alert } from 'react-native'
import { Formik, Field } from 'formik'
import { Input, SelectInput, Datepicker, Timepicker, Button } from 'src/components'
import { updateAdviceByID } from 'src/services/advice.service'
import { getUniversityByID } from 'src/services/university.service'
import { useAuth } from 'src/providers/auth.provider'
import { styles } from './UpdateAdviceForm.styles'
import withLoading from 'src/hocs/withLoading'
import moment from 'moment'

const UpdateAdviceForm = ({ advice }) => {
   
    const { authState } = useAuth()
    const { user } = authState
    const { university_id } = user

    const [ loading, setLoading ] = useState(false)
    const [ classrooms, setClassrooms ] = useState([])

    useEffect(() => {
        const searchUniversityByID = async (universityID) => {
            try{
                const university = await getUniversityByID(universityID)
                setClassrooms(university.classrooms.map(classroom => ({ label: classroom.name, value: classroom.id })))
            }catch(error){
                Alert.alert('¡Ha ocurrido un error!', error.message)
            }
        }
        searchUniversityByID(university_id)
    },  [])

    /* Formik configuration. */
    const initialValues = { 
        subject: advice.subject,
        advice_date: moment(advice.advice_date).format('YYYY-MM-DD'),
        advice_time: moment(advice.advice_date).format('HH:mm'),
        classroom_id: advice.classroom.id
    }

    const handleValidation = ({ subject, advice_date, advice_time, classroom_id }) => {
        let errors = {}
        if(!subject){
            errors.subject = 'Materia requerida'
        }
        if(!advice_date){
            errors.advice_date = 'Fecha de la asesoría requerida'
        }
        if(!advice_time){
            errors.advice_time = 'Hora de la asesoría requerida'
        }
        if(!classroom_id){
            errors.classroom_id = 'Salón requerido'
        }
        return errors
    }

    const onHandleSubmit = async (form) => {
        try{
            setLoading(true)
            await updateAdviceByID(advice.id, { ...form, advice_date: `${form.advice_date} ${form.advice_time}`})
            Alert.alert('¡Felicidades!', '¡Has actualizado la asesoría con éxito!')
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }finally{
            setLoading(false)
        }
    }

    return (
        <Formik 
            initialValues={initialValues}
            validateOnChange={false}
            validateOnBlur={false}
            validate={handleValidation}
            onSubmit={onHandleSubmit}
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
                    <Field name='advice_time'>
                        {({ field, form, meta }) => (
                            <Timepicker 
                                label='Selecciona la hora de la asesoría'
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
                        loadingMessage='Actualizando asesoría...' 
                        onPress={handleSubmit}
                        style={styles.button}
                    >
                        Actualizar asesoría
                    </Button>
                </View>
            )}
        </Formik>
    )
}

export default withLoading(UpdateAdviceForm, 'Cargando asesoría...')