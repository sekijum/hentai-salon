import { AxiosInstance } from 'axios';
import Storage from '~/plugins/storage';

declare module '#app' {
  interface NuxtApp {
    $api: AxiosInstance;
    $storage: typeof Storage;
  }
}

declare module 'vue' {
  interface ComponentCustomProperties {
    $api: AxiosInstance;
    $storage: typeof Storage;
  }
}
