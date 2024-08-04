import { defineNuxtPlugin } from '#app';
import VueGtag from 'vue-gtag';

export default defineNuxtPlugin(nuxtApp => {
  const router = useRouter();

  nuxtApp.vueApp.use(
    VueGtag,
    {
      appName: '変態サロン',
      pageTrackerScreenviewEnabled: true,
      config: { id: 'XXXX' },
    },
    router,
  );
});
