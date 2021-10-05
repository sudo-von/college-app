import React from 'react'
import { ScrollView } from 'react-native'
import AdviceCard from '../AdviceCard/AdviceCard'

const AdvicesList = ({ advices, setAdvices, userID }) =>
    <ScrollView>
        { advices.map((data) => 
            <AdviceCard 
                key={data.id} 
                userID={userID} 
                data={data} 
                setAdvices={setAdvices}
            />
        )}
    </ScrollView>

export default AdvicesList
