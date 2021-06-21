import axios from 'axios';

const commonInst = axios.create({
  baseURL: process.env.VUE_APP_API_URL,
});

function registerUser(userData) {
  return commonInst.post('signup', userData);
}

function loginUser(userData) {
  return commonInst.post('login', userData);
}

export { registerUser, loginUser };
