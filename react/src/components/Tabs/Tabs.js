import React from 'react'
import PropTypes from 'prop-types'
import { View } from 'react-native'
import { styles } from './Tabs.style'

const Tabs = ({ tab=0, setTab, children }) =>
    <View style={styles.view}>
        {   
            children && children.map(( child, i ) => {
            if( tab === i ){
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

Tabs.PropTypes = {
    children: PropTypes.any,
    tab: PropTypes.number,
    setTab: PropTypes.func
}

export default Tabs
