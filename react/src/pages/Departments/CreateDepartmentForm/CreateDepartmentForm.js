import React from 'react'
import { View } from 'react-native'
import { Formik } from 'formik'
import { Input, Button } from 'src/components'
import { styles } from './CreateDepartmentForm.styles'
import { useDepartment } from 'src/hooks/useDepartment'

const CreateDepartmentForm = () => {

    const { loading, handleCreateDepartment } = useDepartment()

    const initialValues = { 
        cost: '',
        description: '',
        neighborhood: '',
        street: ''
    }

    const handleValidation = ({ cost, description, neighborhood, street }) => {
        let errors = {}
        if(!cost){
            errors.cost = 'Costo requerido'
        }
        if(isNaN(cost)){
            errors.cost = 'El costo sólo debe contener números'
        }
        if(!description){
            errors.description = 'Descripción requerida'
        }
        if(!neighborhood){
            errors.neighborhood = 'Colonia requerida'
        }
        if(!street){
            errors.street = 'Calle requerida'
        }
        return errors
    }

    return (
        <Formik 
            initialValues={initialValues}
            validateOnChange={false}
            validateOnBlur={false}
            validate={handleValidation}
            onSubmit={handleCreateDepartment}
        >
            {({ handleChange, handleBlur, handleSubmit, errors, values }) => (
                <View>
                    <Input
                        label='Ingresa la calle'
                        onChangeText={handleChange('street')}
                        onBlur={handleBlur('street')}
                        value={values.street}
                        error={errors.street}
                    />
                    <Input
                        label='Ingresa la colonia'
                        onChangeText={handleChange('neighborhood')}
                        onBlur={handleBlur('neighborhood')}
                        value={values.neighborhood}
                        error={errors.neighborhood}
                    />
                    <Input
                        label='Ingresa la descripción'
                        onChangeText={handleChange('description')}
                        onBlur={handleBlur('description')}
                        value={values.description}
                        error={errors.description}
                    />
                    <Input
                        label='Ingresa el costo (MXN)'
                        keyboardType='numeric'
                        onChangeText={handleChange('cost')}
                        onBlur={handleBlur('cost')}
                        value={values.cost}
                        error={errors.cost}
                    />
                    <Button 
                        loading={loading} 
                        loadingMessage='Registrando departamento...' 
                        onPress={handleSubmit}
                        style={styles.button}
                    >
                        Registrar departamento
                    </Button>
                </View>
            )}
        </Formik>
    )
}

export default CreateDepartmentForm