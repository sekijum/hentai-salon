import { defineNuxtPlugin, useRuntimeConfig } from '#app';
import axios, { type AxiosInstance } from 'axios';

export default defineNuxtPlugin(nuxtApp => {
  const config = useRuntimeConfig();
  const { $storage, $api } = nuxtApp;
  console.log(nuxtApp);

  const api: AxiosInstance = axios.create({
    baseURL: config.public.apiBaseUrl,
    timeout: 10000,
    headers: {
      'Content-Type': 'application/json',
    },
  });

  api.interceptors.request.use(
    config => {
      const token = $storage.getItem<string>('access_token');
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
        $storage.removeItem('access_token');
        window.location.href = '/signin';
      }
      return Promise.reject(error);
    },
  );

  nuxtApp.provide('api', api);
});
