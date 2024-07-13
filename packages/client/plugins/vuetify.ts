import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';

export default defineNuxtPlugin(nuxtApp => {
  const { getTheme } = useStorage();

  const vuetify = createVuetify({
    components,
    directives,
    theme: {
      defaultTheme: getTheme(),
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
});
