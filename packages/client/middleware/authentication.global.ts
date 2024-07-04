import { defineNuxtRouteMiddleware, navigateTo, useNuxtApp } from '#app';
import Storage from '@/utils/storage';

export default defineNuxtRouteMiddleware(async (to, from) => {
  const nuxtApp = useNuxtApp();
  const api = nuxtApp.$api;

  const token = Storage.getItem('access_token');

  if (token) {
    try {
      const response = await api.get('/whoami');
      const authenticatedUser = response.data;
      nuxtApp.payload.user = authenticatedUser;
    } catch (error) {
      console.error('Error fetching user info:', error);
      Storage.removeItem('access_token');
    }
  } else {
    nuxtApp.payload.user = null;
  }
});
