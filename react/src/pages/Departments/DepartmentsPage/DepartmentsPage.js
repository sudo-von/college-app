import React, { useEffect } from 'react'
import { View, ScrollView } from 'react-native'
import { Loader, Container } from 'src/components'
import { IconButton, useTheme } from 'react-native-paper'
import { Tabs, TabScreen } from 'react-native-paper-tabs'
import { useDepartments } from 'src/hooks/useDepartments'
import { styles } from './DepartmentsPage.styles'
import DepartmentCard from '../DepartmentCard/DepartmentCard'

const DepartmentsPage = () => {

    const { colors } = useTheme()
    const { loading, departments } = useDepartments()

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
                                        {...data}
                                    />    
                                )}
                            </ScrollView>
                        </TabScreen>
                        <TabScreen label='Mis departamentos'>
                            <ScrollView>
                                { departments.map(data => 
                                    <DepartmentCard
                                        {...data}
                                    />    
                                )}
                            </ScrollView>
                        </TabScreen>
                    </Tabs>
                }
            </View>
        </Container>
    )
}

export default DepartmentsPage