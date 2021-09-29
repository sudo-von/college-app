import React from 'react'
import { ScrollView } from 'react-native'
/* Custom components. */
import AdviceCard from '../AdviceCard'

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
