import type { IThread } from '~/types/thread';
import type { IThreadComment } from '~/types/thread-comment';
import type { IListResource } from '~/types/list-resource';

export interface IUser {
  id: number;
  name: string;
  role: string;
  email: string;
  profileLink: string;
  createdAt: string;
  updatedAt: string;
  threads: IListResource<IThread>;
  comments: IListResource<IThreadComment>;
}
