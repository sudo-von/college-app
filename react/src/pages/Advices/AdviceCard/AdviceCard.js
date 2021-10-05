import React from 'react'
import { Card } from 'react-native-paper'
import { styles } from './AdviceCard.styles'
import AdviceCardActions from './AdviceCardActions/AdviceCardActions'
import AdviceCardContent from './AdviceCardContent/AdviceCardContent'

const AdviceCard = ({ data, userID, setAdvices }) => {
  const { advice_date, classroom, id, students_will_attend, subject, user } = data
  return(
    <Card style={styles.card}>
      <AdviceCardContent
        adviceDate={advice_date}
        adviceUser={user}
        classroom={classroom}
        studentsNumber={students_will_attend.length}
        subject={subject}
      />
      <AdviceCardActions 
        adviceID={id} 
        adviceUser={user}
        userID={userID}
        studentsWillAttend={students_will_attend}
        setAdvices={setAdvices}
      />
    </Card>
  )
}

export default AdviceCard