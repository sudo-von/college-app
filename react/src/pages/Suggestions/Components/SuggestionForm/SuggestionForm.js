import React, { useState } from 'react'
import { StyleSheet, View, Alert } from 'react-native'
/* Formik. */
import { Formik } from 'formik'
/* Custom components. */
import Button from 'src/components/Button'
import Input from 'src/components/Input'
/* Services. */
import { sendSuggestion } from 'src/services/suggestion.service'

const SuggestionForm = () => {

    const [ loading, setLoading ] = useState(false)
    const onSubmit = async ({ suggestion }, { resetForm }) => {
        try{
            setLoading(true)
            await sendSuggestion({ suggestion })
            resetForm()
            Alert.alert('¡Increíble!', 'Tu sugerencia ha sido enviada con éxito.')
        }catch(error){
            console.log(error)
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }finally{
            setLoading(false)
        }
    }
        
    return (
        <Formik 
            initialValues={{ suggestion: '' }} 
            validate={({suggestion}) => {
                const errors = {}
                if(!suggestion){
                    errors.suggestion = 'Sugerencia requerida'
                }
                return errors
            }}
            onSubmit={onSubmit}
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

const styles = StyleSheet.create({
    container: {
    },
    button: {
        marginTop: 10,
    }
})

export default SuggestionForm
