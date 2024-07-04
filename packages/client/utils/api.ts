import axios, { type AxiosInstance } from 'axios';
import Storage from './storage';
import { useRuntimeConfig } from '#app';

const config = useRuntimeConfig();

const api: AxiosInstance = axios.create({
  baseURL: 'http://localhost:8080' || config.public.apiBaseUrl,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

api.interceptors.request.use(
  config => {
    const token = Storage.getItem<string>('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  error => {
    return Promise.reject(error);
  },
);

api.interceptors.response.use(
  response => {
    return response;
  },
  error => {
    if (error.response && error.response.status === 401) {
      Storage.removeItem('token');
      window.location.href = '/signin';
    }
    return Promise.reject(error);
  },
);

export default api;
