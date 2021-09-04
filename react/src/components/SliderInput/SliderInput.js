import React, { Fragment, useState } from 'react'
import { View, Platform, StyleSheet } from 'react-native'
/* React native community. */
import Slider from '@react-native-community/slider'
/* React native paper. */
import { useTheme } from '@react-navigation/native'
import Small from 'src/components/Small'
import Container from 'src/components/Container'

const SliderInput = ({ initialValue, minimumValue, maximumValue, changeValue, minimumText, maximumText}) => {
    const { colors } = useTheme()
    return (
        <View>
            <Slider
                style={styles.slider}
                minimumValue={minimumValue}
                maximumValue={maximumValue}
                minimumTrackTintColor={colors.primary}
                maximumTrackTintColor={colors.primary}
                value={initialValue}
                onValueChange={(currentValue) => changeValue(currentValue)}
            />
            <View style={styles.view}>
                <Small>{minimumText}</Small>
                <Small>{maximumText}</Small>
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    slider: {
        marginLeft: Platform.select({ ios: 0, android: -15 }),
        marginRight: Platform.select({ ios: 0, android: -15 })  
    },
    view: {
        flexDirection: 'row', 
        justifyContent: 'space-between', 
    }
})

export default SliderInput
