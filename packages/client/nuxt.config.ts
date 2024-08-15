import vuetify, { transformAssetUrls } from 'vite-plugin-vuetify';
export default defineNuxtConfig({
  ssr: false,

  build: {
    transpile: ['vuetify'],
  },

  runtimeConfig: {
    public: {
      apiBaseUrl: process.env.NUXT_PUBLIC_API_BASE_URL,
      appEnv: process.env.APP_ENV,
      staticUrl: process.env.STATIC_URL,
    },
  },

  css: ['vuetify/styles', '@mdi/font/css/materialdesignicons.css', 'glightbox/dist/css/glightbox.css'],

  modules: [
    (_options, nuxt) => {
      nuxt.hooks.hook('vite:extendConfig', config => {
        // @ts-expect-error
        config.plugins.push(vuetify({ autoImport: true }));
      });
    },
    //...
  ],

  vite: {
    vue: {
      template: {
        transformAssetUrls,
      },
    },
  },

  app: {
    head: {
      title: '変態サロン',
      meta: [
        {
          name: 'description',
          content: '変態に特化したサロン。',
        },
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/hentai-salon-logo/logo.png' },
        { rel: 'apple-touch-icon', sizes: '180x180', href: '/hentai-salon-logo/logo.png' },
        { rel: 'icon', type: 'image/png', sizes: '32x32', href: '/hentai-salon-logo/logo.png' },
        { rel: 'icon', type: 'image/png', sizes: '16x16', href: '/hentai-salon-logo/logo.png' },
      ],
    },
  },

  plugins: ['~/plugins/storage'],
});
