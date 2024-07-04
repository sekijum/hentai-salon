import { defineNuxtRouteMiddleware, navigateTo, useNuxtApp } from '#app';
import { getAuthenticatedUser } from '@/actions/user';
import Storage from '@/utils/storage';

export default defineNuxtRouteMiddleware(async (to, from) => {
  const nuxtApp = useNuxtApp();
  const token = Storage.getItem<string>('token');

  if (token) {
    try {
      const authenticatedUser = await getAuthenticatedUser();
      nuxtApp.provide('user', authenticatedUser);
    } catch (error) {
      console.error('Error fetching user info:', error);
      Storage.removeItem('token');
      return navigateTo('/signin');
    }
  } else {
    nuxtApp.provide('user', null);
    if (to.name !== 'signin') {
      return navigateTo('/signin');
    }
  }
});
