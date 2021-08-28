/* Base64. */
import base64 from 'react-native-base64'

export {
    login
}

/* Handles basic auth. */
const login = async ({ email, password }) => {
    
    const options = {
        method: 'POST',
        headers: { 
            'Authorization': "Basic " + base64.encode(`${email}:${password}`)
        }
    }

    try{
        const response = await fetch(`http://192.168.0.4:4000/users/login`, options)
        if(!response.ok){
            throw new Error('Credenciales inválidas')
        }
        return response.headers.map.authorization
    }catch(error){
        return error
    }
}


const handleResponse = (response) => {
    return response.text().then(text => {
        const data = text && JSON.parse(text);
        if (!response.ok) {
            if (response.status === 401) {
                // auto logout if 401 response returned from api
                logout();
                location.reload(true);
            }

            const error = (data && data.message) || response.statusText;
            return Promise.reject(error);
        }

        return data;
    });
}