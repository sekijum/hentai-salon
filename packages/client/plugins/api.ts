import { defineNuxtPlugin, useRuntimeConfig } from '#app';
import axios from 'axios';
import Storage from '@/utils/storage';

export default defineNuxtPlugin(nuxtApp => {
  const config = useRuntimeConfig();

  const api = axios.create({
    baseURL: config.public.apiBaseUrl,
    timeout: 10000,
    headers: {
      'Content-Type': 'application/json',
    },
  });

  api.interceptors.request.use(
    config => {
      const token = Storage.getItem<string>('access_token');
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
        Storage.removeItem('access_token');
        window.location.href = '/signin';
      }
      return Promise.reject(error);
    },
  );

  nuxtApp.provide('api', api);
});
