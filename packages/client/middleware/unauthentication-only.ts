import { defineNuxtRouteMiddleware, navigateTo } from '#app';

export default defineNuxtRouteMiddleware((to, from) => {
  const nuxtApp = useNuxtApp();
  const { payload } = nuxtApp;

  console.log(payload);
  if (payload.isLoggedIn) {
    return navigateTo('/');
  }
});
