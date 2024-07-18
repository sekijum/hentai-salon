import axios, { type AxiosInstance } from 'axios';

function API(): AxiosInstance {
  const config = useRuntimeConfig();
  const { $storage } = useNuxtApp();

  const api: AxiosInstance = axios.create({
    baseURL: config.public.apiBaseUrl,
    timeout: 10000,
    headers: {
      'Content-Type': 'application/json',
    },
  });

  api.interceptors.request.use(
    config => {
      const token = $storage.getItem('access_token');
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
      // if (error.response && error.response.status === 401) {
      //   $storage.removeItem('access_token');
      //   window.location.href = '/signin';
      // }
      return Promise.reject(error);
    },
  );

  return api;
}

export default defineNuxtPlugin(nuxtApp => {
  nuxtApp.provide('api', API());
});
