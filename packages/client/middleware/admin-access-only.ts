import { defineNuxtRouteMiddleware, navigateTo } from '#app';

export default defineNuxtRouteMiddleware((to, from) => {
  const nuxtApp = useNuxtApp;
  const { $user } = nuxtApp;

  if ($user) {
    const isAdmin = $user.role === 'admin';

    if (!isAdmin) {
      return navigateTo('/403');
    }
  } else {
    if (to.name !== 'signin') {
      return navigateTo('/signin');
    }
  }
});
