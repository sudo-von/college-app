import React from 'react'
import { Alert, StyleSheet, Linking } from 'react-native'
import { IconButton, Card } from 'react-native-paper'
import { deleteAdviceByID, updateStudentsNumber } from 'src/services/advice.service'
import { useNavigation } from '@react-navigation/native'

const AdviceCardActions = ({ studentsWillAttend, adviceID, adviceUser, userID, setAdvices }) => {

    const attended = studentsWillAttend.includes(userID)
    const navigation = useNavigation()
    /* Updates the number of students that will go to the advice. */
    const handleAssistance = async () => {
        try{
            await updateStudentsNumber(adviceID)
            if(attended){
                setAdvices(advices => advices.map(advice => advice.id == adviceID ?
                    {...advice, students_will_attend: advice.students_will_attend.filter(
                        studentID => studentID != userID)
                    } :
                    advice
                ))
            }else{
                setAdvices(advices => advices.map(advice => advice.id == adviceID ?
                    {...advice, students_will_attend: [...advice.students_will_attend, userID]} :
                    advice
                ))
            }
        }catch(error){
            Alert.alert('¡Ha ocurrido un error!', error.message)
        }
    }

    const handleEdit = () => {
        navigation.navigate('/update-advice', {
            id: adviceID
        })
    }

    /* Delete the advice and updates the advices list. */
    const handleDelete = () => {
        Alert.alert(
            '¿Quieres eliminar esta asesoría?',
            'Tu asesoría puede ayudar mucho a otros estudiantes. Recuerda que esta acción no podrá ser revertida.',
            [
                {
                    text: 'Cancelar',
                },
                { 
                    text: 'Eliminar asesoría', 
                    onPress: async () => {
                        try{
                            await deleteAdviceByID(adviceID)
                            setAdvices(advices => advices.filter(advice => advice.id != adviceID))
                        }catch(error){
                            Alert.alert('¡Ha ocurrido un error!', error.message)
                        }
                    }
                }
            ],
        )
    }

    /* Open a new mail. */
    const handleEmail = () => {
        Linking.openURL(`mailto:${adviceUser.email}`)
    }

    const buttons = [
        {
            icon: 'thumb-up',
            color: attended ? '#4FB1F8' : '#A7BBCA',
            onPress: handleAssistance,
            validateUser: false
        },
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
        <Card.Actions>
            { buttons.map(({ icon, color, onPress, validateUser }) => {
                if(validateUser){
                    if(userID == adviceUser.id){
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

const styles = StyleSheet.create({
})

export default AdviceCardActions