// plugins/vuetify.js
import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';
import { useState } from '#app';

export default defineNuxtPlugin(nuxtApp => {
  const theme = useState('theme', () => 'light');

  const vuetify = createVuetify({
    components,
    directives,
    theme: {
      defaultTheme: theme.value,
      themes: {
        dark: {
          dark: true,
          colors: {
            background: '#000000',
            surface: '#121212',
            primary: '#BB86FC',
            secondary: '#03DAC6',
            error: '#CF6679',
          },
        },
        light: {
          dark: false,
          colors: {
            background: '#FFFFFF',
            surface: '#FFFFFF',
            primary: '#1E88E5',
            secondary: '#03A9F4',
            error: '#FF5252',
          },
        },
      },
    },
  });

  nuxtApp.vueApp.use(vuetify);

  const toggleTheme = () => {
    theme.value = theme.value === 'dark' ? 'light' : 'dark';
    vuetify.theme.global.name.value = theme.value;
  };

  nuxtApp.provide('toggleTheme', toggleTheme);
  nuxtApp.provide('theme', theme);
});
