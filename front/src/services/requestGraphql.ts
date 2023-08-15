import axios from 'axios';
import { API_BASE_URL } from "../constants/urlConstants";
const makeGraphQLCall = async (query: string, variables: any = {}) => {
    const config = {
        method: 'POST',
        url: API_BASE_URL,
        headers: {
            'Content-Type': 'application/json',
        },
    };

    const response = await axios({
        ...config,
        data: JSON.stringify({ query, variables }),
    });

    if (response.status === 200) {
        return response;
    } else {
        throw new Error(response.statusText);
    }
};

export default makeGraphQLCall;
