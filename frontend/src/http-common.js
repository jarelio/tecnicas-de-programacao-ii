import axios from 'axios';

//Define a URL base da origem para consumo do servico
export default axios.create({
  baseURL: 'http://192.168.254.129:8000/',
  headers: {
    'Content-type': 'application/json',
  },
});
