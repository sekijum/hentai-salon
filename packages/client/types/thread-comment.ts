import type { IListResource } from './list-resource';
import type { IThreadCommentAttachment } from './thread-comment-attachment';
import type { IThreadComment as IComment } from './thread';

interface IThread {
  id: number;
  title: string;
  description: string;
}

interface IUser {
  id: number;
  name: string;
  profileLink?: string;
}

export interface IThreadComment {
  id: number;
  guestName?: string;
  content: string;
  createdAt: string;
  updatedAt: string;
  attachments: IThreadCommentAttachment[];
  user?: IUser;
  thread: IThread;
  parentCommentId: number;
  parentComment: IComment;
  replies: IListResource<IComment>;
  replyCount: number;
}
