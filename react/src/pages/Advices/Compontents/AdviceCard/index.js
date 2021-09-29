import React from 'react'
import { StyleSheet } from 'react-native'
/* Custom components. */
import AdviceCardActions from './Components/AdviceCardActions'
import AdviceCardContent from './Components/AdviceCardContent'
/* React native paper. */
import { Card } from 'react-native-paper'

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

const styles = StyleSheet.create({
  card: {
    margin: 10
  }
})

export default AdviceCard