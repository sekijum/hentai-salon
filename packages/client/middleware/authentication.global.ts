import { defineNuxtRouteMiddleware } from '#app';

export default defineNuxtRouteMiddleware(async (to, from) => {
  const nuxtApp = useNuxtApp();
  const { $api, $storage } = nuxtApp;

  const token = $storage.getItem('access_token');

  if (token) {
    try {
      const response = await $api.get('/whoami');
      const authenticatedUser = response.data;
      nuxtApp.payload.user = authenticatedUser;
      nuxtApp.payload.isLoggedIn = true;
      nuxtApp.payload.isAdmin = authenticatedUser.role === 'Admin';
      nuxtApp.payload.isMember = authenticatedUser.role === 'Member';
    } catch (error) {
      console.error('Error fetching user info:', error);
      // $storage.removeItem('access_token');
    }
  } else {
    nuxtApp.payload.user = null;
    nuxtApp.payload.isLoggedIn = false;
    nuxtApp.payload.isAdmin = false;
    nuxtApp.payload.isMember = false;
  }
});