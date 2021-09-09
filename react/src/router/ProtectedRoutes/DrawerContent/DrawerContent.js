import React from 'react'
import { View, StyleSheet, Image } from 'react-native'
/* React navigation. */
import { DrawerContentScrollView, DrawerItemList, DrawerItem } from '@react-navigation/drawer'
/* Custom components. */
import Header from 'src/components/Header'
/* Contexts. */
import { useAuth } from 'src/providers/auth.provider'

const DrawerContent = (props) => {
    const { authState } = useAuth()
    const { user : { university_name, university_profile_picture } } = authState
    return (
        <DrawerContentScrollView {...props}>
            <DrawerItem
                label={({ focused, color }) => 
                    (
                        <View>
                            <View style={styles.view}>
                                <Image style={styles.image} source={{ uri: university_profile_picture }}/>
                            </View>
                            <Header
                                title={university_name}
                                subtitle='Universidad Autónoma de Gojs'
                            />
                        </View>
                    )
                }
            />
            <DrawerItemList {...props}/>
        </DrawerContentScrollView>
    ) 
}

const styles = StyleSheet.create({
    view: {
        width: 240,
        alignItems: 'center'
    },
    title: {
        fontSize: 16,
        marginTop: 10
    },
    image: {
        width: 180,
        height: 180,
        resizeMode: 'contain'
    }
})

export default DrawerContent