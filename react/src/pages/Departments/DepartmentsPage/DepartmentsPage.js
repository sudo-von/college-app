import React from 'react'
import { View, ScrollView } from 'react-native'
import { Loader, Container } from 'src/components'
import { IconButton, useTheme } from 'react-native-paper'
import { Tabs, TabScreen } from 'react-native-paper-tabs'
import { useUser } from 'src/hooks/useUser'
import { useDepartments } from 'src/hooks/useDepartments'
import { styles } from './DepartmentsPage.styles'
import { useNavigation } from '@react-navigation/native'
import DepartmentCard from '../DepartmentCard/DepartmentCard'

const DepartmentsPage = () => {

    const navigation = useNavigation()
    const { colors } = useTheme()
    const { user: { user_id } } = useUser()
    const { loading, departments, setDepartments } = useDepartments()

    return (
        <Container style={styles(colors).container}>
            <View style={styles(colors).view}>
                { loading ? 
                    <Loader loadingMessage='Cargando departamentos'/> :
                    <Tabs style={styles(colors).tabs}>
                        <TabScreen label='Todos'>
                            <ScrollView>
                                { departments.map(data => 
                                    <DepartmentCard
                                        key={data.id}
                                        userID={user_id}
                                        setDepartments={setDepartments}
                                        {...data}
                                    />    
                                )}
                            </ScrollView>
                        </TabScreen>
                        <TabScreen label='Mis departamentos'>
                            <ScrollView>
                                { departments.map(data => 
                                    <DepartmentCard
                                        key={data.id}
                                        userID={user_id}
                                        setDepartments={setDepartments}
                                        {...data}
                                    />    
                                )}
                            </ScrollView>
                        </TabScreen>
                    </Tabs>
                }
            </View>
            <View style={styles(colors).bottomView}>
                <IconButton
                    icon='plus'
                    color='gray'
                    onPress={() => navigation.navigate('/create-department')}
                />
            </View>
        </Container>
    )
}

export default DepartmentsPage