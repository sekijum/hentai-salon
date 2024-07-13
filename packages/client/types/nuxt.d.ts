import { AxiosInstance } from 'axios';
import { IStorage } from '~/plugins/storage';
import dayjs from 'dayjs';

declare module '#app' {
  interface NuxtApp {
    $toggleTheme: () => void;
    $api: AxiosInstance;
    $storage: IStorage;
    $dayjs: dayjs;
    $formatDate: (
      date: Date | string,
      {
        timezone,
        format,
      }?: {
        timezone?: string | undefined;
        format?: string | undefined;
      },
    ) => string;
  }
}

declare module 'vue' {
  interface ComponentCustomProperties {
    $toggleTheme: () => void;
    $api: AxiosInstance;
    $storage: IStorage;
    $dayjs: dayjs;
    $formatDate: (
      date: Date | string,
      {
        timezone,
        format,
      }?: {
        timezone?: string | undefined;
        format?: string | undefined;
      },
    ) => string;
  }
}
