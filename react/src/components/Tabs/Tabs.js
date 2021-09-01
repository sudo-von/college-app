import React from 'react'
import { StyleSheet, View } from 'react-native'

const Tabs = ({ tab, setTab, children }) =>
    <View style={styles.container}>
        {children && children.map(( child, i ) => {
            if( tab == i ){
                return(
                    <child.type 
                        {...child.props} 
                        key={`tab-${i}`} 
                        tab={tab} 
                        setTab={setTab}
                    />    
                )
            }
        })}
    </View>

const styles = StyleSheet.create({
    container: {
        width: '100%'
    }
})

export default Tabs
