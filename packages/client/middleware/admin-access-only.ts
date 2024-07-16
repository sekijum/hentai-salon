import { defineNuxtRouteMiddleware, navigateTo } from '#app';

export default defineNuxtRouteMiddleware((to, from) => {
  const nuxtApp = useNuxtApp();
  const { payload } = nuxtApp;

  if (payload.isLoggedIn) {
    if (!payload.isAdmin) {
      return navigateTo('/403');
    }
  } else {
    if (to.name !== 'signin') {
      return navigateTo('/signin');
    }
  }
});
