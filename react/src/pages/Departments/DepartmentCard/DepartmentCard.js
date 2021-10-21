import React from 'react'
import DepartmentCardContent from './DepartmentCardContent/AdviceCardContent'
import DepartmentCardActions from './DepartmentCardActions/DepartmentCardActions'
import { Card } from 'react-native-paper'
import { styles } from './DepartmentCard.styles'

const DepartmentCard = ({ id, available, cost, description, neighborhood, street, user, userID, setDepartments }) =>
  <Card style={styles.card}>
    <DepartmentCardContent
      available={available}
      cost={cost}
      description={description}
      neighborhood={neighborhood}
      street={street}
      departmentUser={user}
    />
    <DepartmentCardActions 
      id={id}
      departmentUser={user}
      setDepartments={setDepartments}
      userID={userID}
    />
  </Card>

export default DepartmentCard