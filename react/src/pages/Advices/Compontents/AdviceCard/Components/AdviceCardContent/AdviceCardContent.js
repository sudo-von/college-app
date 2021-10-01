import React from 'react'
import { StyleSheet } from 'react-native'
/* Custom components. */
import Bold from 'src/components/Bold'
import Small from 'src/components/Small'
/* React native paper. */
import { Card, Title, Paragraph, useTheme } from 'react-native-paper'
/* Moment. */
import moment from 'moment'

const AdviceCardContent = ({ adviceUser, adviceDate, classroom, studentsNumber, subject }) => {
    
    const { colors } = useTheme()
    const date = moment(adviceDate).format('DD-MM-YYYY')
    const time = moment(adviceDate).format('hh:mm a')

    return (
        <Card.Content>  
            <Title>
                <Bold>{subject}</Bold>
            </Title>
            <Small style={styles(colors).small}>
                <Bold>{adviceUser.name}</Bold> / {adviceUser.email}
            </Small>
            <Paragraph>
                El día {date} a las {time} impartiré una asesoría de {subject} en el salón {classroom.name}.
            </Paragraph>
            { studentsNumber > 0 && 
                <Paragraph>Se espera la asistencia de {studentsNumber}
                    { studentsNumber == 1 ? ' estudiante' : ' estudiantes'}.
                </Paragraph>
            }
        </Card.Content>
    )
}

const styles = (colors) => StyleSheet.create({
    small: {
      color: colors.primary,
    },
})
  
export default AdviceCardContent
