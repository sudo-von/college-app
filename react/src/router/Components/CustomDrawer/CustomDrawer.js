import React from 'react'
import { View, StyleSheet } from 'react-native'
/* React navigation. */
import { DrawerContentScrollView, DrawerItemList, DrawerItem } from '@react-navigation/drawer';
/* Custom components. */
import Logo from 'src/components/Logo'
import Small from 'src/components/Small'
import Header from 'src/components/Header'
/* React native paper. */
import { Title } from 'react-native-paper'

const CustomDrawer = (props) => {
    const logo = require('../../../../assets/gopher.png')
    return (
        <DrawerContentScrollView {...props}>
            <DrawerItem
                label={({ focused, color }) => 
                    (
                        <View>
                            <View style={styles.view}>
                                <Logo source={logo}/>
                            </View>
                            <Header
                                title='Facultad de electrones'
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
    }
})

export default CustomDrawer