import { AxiosInstance } from 'axios';
import { IStorage } from '~/plugins/storage';
import dayjs, { type formatDate } from 'dayjs';

interface IType {
  payload: {
    user?: {
      id: number,
      name: string,
      role: string,
      email: string,
      profileLink: string,
      createdAt: string,
      updatedAt: string,
    },
    isLoggedIn: boolean,
    isAdmin: boolean,
    isMember: boolean,
  };
  $api: AxiosInstance;
  $storage: IStorage;
  $dayjs: dayjs;
  $formatDate: formatDate;
}

declare module '#app' {
  interface NuxtApp extends IType
}

declare module 'vue' {
  interface ComponentCustomProperties extends IType
}
