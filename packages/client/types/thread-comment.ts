import type { ICollection } from './collection';
import type { IThreadCommentAttachment } from './thread-comment-attachment';
import type { IThreadComment as IComment } from './thread';
import type { IUser } from './user';

interface IThread {
  id: number;
  title: string;
  description: string;
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
  replies: ICollection<IComment>;
  replyCount: number;
}
