import React from 'react'
import { View, StyleSheet, Image } from 'react-native'
/* React navigation. */
import { DrawerContentScrollView, DrawerItemList } from '@react-navigation/drawer'
/* Custom components. */
import { Header } from 'src/components'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'

const DrawerContent = (props) => {
    
    const { authState } = useAuth()
    const { user } = authState
    const { university_name, university_profile_picture } = user
    
    return (
        <DrawerContentScrollView {...props}>
            <View style={styles.drawerContent}>
                <View style={styles.imageContainer}>
                    <Image style={styles.image} source={{ uri: university_profile_picture }}/>
                </View>
                <Header
                    title={university_name}
                    subtitle='Universidad Autónoma de Gojs'
                />
            </View>
            <DrawerItemList {...props}/>
        </DrawerContentScrollView>
    ) 
}

const styles = StyleSheet.create({
    drawerContent: {
        margin: 10,
        padding: 10
    },
    imageContainer: {
        alignItems: 'center'
    },
    image: {
        width: 150,
        height: 150,
        resizeMode: 'contain'
    }
})

export default DrawerContent