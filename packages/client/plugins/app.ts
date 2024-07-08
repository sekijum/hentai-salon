import { nextTick } from 'vue';

export default defineNuxtPlugin(nuxtApp => {
  const router = useRouter();
  router.options.scrollBehavior = (to, from, savedPosition) => {
    if (savedPosition) {
      return savedPosition;
    }
    if (to.hash) {
      return new Promise(resolve => {
        setTimeout(() => {
          resolve({ el: to.hash, behavior: 'smooth' });
        }, 300);
      });
    } else {
      if (to.path === from.path) return;
      return new Promise((resolve, reject) => {
        nuxtApp.hook('page:finish', () => {
          nextTick().then(() => {
            resolve({ top: 0, left: 0 });
          });
        });
      });
    }
  };
});
