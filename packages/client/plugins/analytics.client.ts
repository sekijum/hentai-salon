import VueGtag from 'vue-gtag';

export default defineNuxtPlugin(nuxtApp => {
  const router = useRouter();
  const config = useRuntimeConfig();

  nuxtApp.vueApp.use(
    VueGtag,
    {
      appName: '変態サロン',
      pageTrackerScreenviewEnabled: true,
      config: { id: config.public.gaMeasurementId },
    },
    router,
  );
});
