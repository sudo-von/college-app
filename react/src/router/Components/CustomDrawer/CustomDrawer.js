import React from 'react'
import { View } from 'react-native'
/* React navigation. */
import { DrawerContentScrollView, DrawerItemList, DrawerItem } from '@react-navigation/drawer';
/* Custom components. */
import Bold from 'src/components/Bold'
import Logo from 'src/components/Logo'
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
                            <View style={{width: 240, alignItems: 'center'}}>
                                <Logo source={logo}/>
                            </View>
                            <Title>Facultad de electrones</Title>
                        </View>
                    )
                }
            />
            <DrawerItemList {...props}/>
        </DrawerContentScrollView>
    ) 
}

export default CustomDrawer