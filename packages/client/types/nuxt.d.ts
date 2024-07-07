import { AxiosInstance } from 'axios';
import Storage from '~/plugins/storage';

declare module '#app' {
  interface NuxtApp {
    $api: AxiosInstance;
    $storage: Storage;
  }
}

declare module 'vue' {
  interface ComponentCustomProperties {
    $api: AxiosInstance;
    $storage: Storage;
  }
}
