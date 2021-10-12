import React from 'react'
import { View, Platform, StyleSheet } from 'react-native'
/* React native community. */
import Slider from '@react-native-community/slider'
/* React native paper. */
import { useTheme } from 'react-native-paper'
/* Custom components. */
import { Small } from 'src/components'

const SliderInput = ({ initialValue, minimumValue, maximumValue, changeValue, minimumText, maximumText, ...rest }) => {
    const { colors } = useTheme()
    return (
        <View {...rest}>
            <Slider
                style={styles.slider}
                minimumValue={minimumValue}
                maximumValue={maximumValue}
                minimumTrackTintColor={colors.primary}
                maximumTrackTintColor={colors.primary}
                thumbTintColor={colors.primary}
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
