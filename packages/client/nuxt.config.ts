import vuetify, { transformAssetUrls } from 'vite-plugin-vuetify';
export default defineNuxtConfig({
  ssr: false,

  build: {
    transpile: ['vuetify'],
  },

  runtimeConfig: {
    public: {
      apiBaseUrl: process.env.NUXT_PUBLIC_API_BASE_URL,
      clientId: process.env.IMGUR_CLIENT_ID,
    },
  },

  css: ['vuetify/styles', '@mdi/font/css/materialdesignicons.css'],

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
      titleTemplate: '変態サロン',
      meta: [
        {
          name: 'description',
          content: '変態に特化したサロン。',
        },
      ],
    },
  },

  compatibilityDate: '2024-07-07',

  plugins: ['~/plugins/storage'],
});
