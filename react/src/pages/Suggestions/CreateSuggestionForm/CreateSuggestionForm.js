import React, { useState } from 'react'
import { View, Alert } from 'react-native'
import { Formik } from 'formik'
import { Button, Input } from 'src/components'
import { sendSuggestion } from 'src/services/suggestion.service'
import { styles } from './CreateSuggestionForm.style'

const CreateSuggestionForm = () => {

    const [ loading, setLoading ] = useState(false)

    /* Formik configuration. */
    const initialValues = { suggestion : ''}

    const handleValidation = ({ suggestion }) => {
        let errors = {}
        if(!suggestion){
            errors.suggestion = 'Sugerencia requerida'
        }
        return errors
    }

    const onHandleSubmit = async (form, { resetForm }) => {
        try{
            setLoading(true)
            await sendSuggestion(form)
            resetForm()
            Alert.alert('¡Felicidades!', 'Tu sugerencia ha sido enviada con éxito.')
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }finally{
            setLoading(false)
        }
    }
        
    return (
        <Formik 
            initialValues={initialValues} 
            validate={handleValidation}
            validateOnChange={false}
            validateOnBlur={false}
            onSubmit={onHandleSubmit}
        >
            {({ handleChange, handleBlur, handleSubmit, errors, values }) => (
                <View style={styles.container}>
                    <Input
                        label='Escribe tu sugerencia'
                        onChangeText={handleChange('suggestion')}
                        onBlur={handleBlur('suggestion')}
                        value={values.suggestion}
                        error={errors.suggestion}
                        multiline={true}
                        numberOfLines={5}
                    />
                    <Button 
                        loading={loading} 
                        loadingMessage='Enviando sugerencia' 
                        style={styles.button}
                        onPress={handleSubmit}
                    >
                        Enviar sugerencia
                    </Button>
                </View>
            )}
        </Formik>
    )
}

export default CreateSuggestionForm
