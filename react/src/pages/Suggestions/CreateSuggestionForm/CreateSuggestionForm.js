import React from 'react'
import { View } from 'react-native'
import { Formik } from 'formik'
import { Button, Input } from 'src/components'
import { useSuggestion } from 'src/hooks/useSuggestion'
import { styles } from './CreateSuggestionForm.style'

const CreateSuggestionForm = () => {

    const { loading, handleSuggestion } = useSuggestion()

    const initialValues = { 
        suggestion : '' 
    }

    const handleValidation = ({ suggestion }) => {
        let errors = {}
        if(!suggestion){
            errors.suggestion = 'Sugerencia requerida'
        }
        return errors
    }
        
    return (
        <Formik 
            initialValues={initialValues} 
            validate={handleValidation}
            validateOnChange={false}
            validateOnBlur={false}
            onSubmit={handleSuggestion}
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
