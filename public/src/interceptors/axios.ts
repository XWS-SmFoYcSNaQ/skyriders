import axios from 'axios';

axios.defaults.baseURL = process.env.REACT_APP_API;

let refresh = false;

axios.interceptors.response.use(resp => resp, async error => {
    if (error.response.status === 401 && !refresh) {
        refresh = true;
        const response = await axios.get('auth/refresh', { withCredentials: true });

        if (response.status === 200) {
            axios.defaults.headers.common['Authorization'] = `Bearer ${response.data['access_token']}`

            return axios(error.config)
        }
    }
    refresh = false;
    return error;
});