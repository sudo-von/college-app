import React from 'react'
import { Alert, Linking } from 'react-native'
import { IconButton, Card } from 'react-native-paper'
import { deleteDepartmentByID } from 'src/services/department.service'
import { useNavigation } from '@react-navigation/native'

const DepartmentCardActions = ({ id, departmentUser, userID, setDepartments }) => {

    const navigation = useNavigation()

    const handleEdit = () => {
        navigation.navigate('/update-department', { id })
    }

    const handleDelete = () => {
        Alert.alert(
            '¿Quieres eliminar este departamento?',
            'Recuerda que esta acción no podrá ser revertida.',
            [
                {
                    text: 'Cancelar',
                },
                { 
                    text: 'Eliminar', 
                    onPress: async () => {
                        try{
                            await deleteDepartmentByID(id)
                            setDepartments(departments => departments.filter(department => department.id != id))
                        }catch(error){
                            Alert.alert('¡Ha ocurrido un error!', error.message)
                        }
                    }
                }
            ],
        )
    }

    const handleEmail = () => {
        Linking.openURL(`mailto:${departmentUser.email}`)
    }

    const buttons = [
        {
            icon: 'email',
            color: '#F8DD4F',
            onPress: handleEmail,
            validateUser: false
        },
        {
            icon: 'pencil',
            color: '#F8B44F',
            onPress: handleEdit,
            validateUser: true
        },
        {
            icon: 'delete',
            color: '#F84F4F',
            onPress: handleDelete,
            validateUser: true
        }
    ]


    return(
        <Card.Actions style={{ margin: 0, padding: 0 }}>
            { buttons.map(({ icon, color, onPress, validateUser }) => 
            {
                if(validateUser){
                    if(userID == departmentUser.id){
                        return(
                            <IconButton
                                key={`${userID}-${icon}`}
                                icon={icon}
                                color={color}
                                size={20}
                                onPress={onPress}
                            />
                        )
                    }
                }else{
                    return(
                        <IconButton
                            key={`${userID}-${icon}`}
                            icon={icon}
                            color={color}
                            size={20}
                            onPress={onPress}
                        />
                    )
                }
            })}
        </Card.Actions>
    )
}


export default DepartmentCardActions