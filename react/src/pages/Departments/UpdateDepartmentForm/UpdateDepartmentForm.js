import React from 'react'
import { View } from 'react-native'
import { Formik, Field } from 'formik'
import { Input, Button, SelectInput } from 'src/components'
import { styles } from './UpdateDepartmentForm.styles'
import { useDepartment } from 'src/hooks/useDepartment'

const UpdateDepartmentForm = ({ id, available, cost, description, neighborhood, street }) => {

    const { loading, handleUpdateDepartment } = useDepartment()

    const initialValues = { 
        cost: cost && cost.toString(),
        description,
        neighborhood,
        street,
        available
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

    const handleSubmit = (form, { resetForm }) => handleUpdateDepartment(id, form, resetForm)

    return (
        <Formik 
            initialValues={initialValues}
            validateOnChange={false}
            validateOnBlur={false}
            validate={handleValidation}
            onSubmit={handleSubmit}
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
                    <Field name='available'>
                        {({ field, form, meta }) => (
                            <SelectInput
                                label={'Selecciona el estatus del departamento'}
                                data={[
                                    {
                                        label: 'Disponible',
                                        value: true,
                                    },
                                    {
                                        label: 'Rentado',
                                        value: false,
                                    }
                                ]}
                                value={available}
                                field={field} 
                                form={form} 
                                meta={meta}
                                error={errors.available}
                            />
                        )}
                    </Field>
                    <Button 
                        loading={loading} 
                        loadingMessage='Actualizando departamento...' 
                        onPress={handleSubmit}
                        style={styles.button}
                    >
                        Actualizar departamento
                    </Button>
                </View>
            )}
        </Formik>
    )
}

export default UpdateDepartmentForm