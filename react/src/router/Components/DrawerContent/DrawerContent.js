import React from 'react'
import { View, Image } from 'react-native'
import { DrawerContentScrollView, DrawerItemList } from '@react-navigation/drawer'
import { Header } from 'src/components'
import { useUser } from 'src/hooks/useUser'
import { styles } from './DrawerContent.styles'

const DrawerContent = (props) => {
    
    const { user } = useUser()
    const { university_name, university_profile_picture } = user
    
    return (
        <DrawerContentScrollView {...props}>
            <View style={styles.drawerContent}>
                <View style={styles.imageContainer}>
                    <Image style={styles.image} source={{ uri: university_profile_picture }}/>
                </View>
                <Header
                    title={university_name}
                    subtitle='Universidad Autónoma de Kyoto'
                />
            </View>
            <DrawerItemList {...props}/>
        </DrawerContentScrollView>
    ) 
}

export default DrawerContent