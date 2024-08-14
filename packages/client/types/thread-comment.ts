import type { ICollection } from './collection';
import type { IThreadCommentAttachment } from './thread-comment-attachment';
import type { IThread } from './thread';
import type { IUser } from './user';

export interface IThreadComment {
  id: string;
  guestName?: string;
  content: string;
  createdAt: string;
  updatedAt: string;
  attachments: IThreadCommentAttachment[];
  user?: IUser;
  thread: IThread;
  parentCommentId: string;
  parentComment: IThreadComment;
  replies: ICollection<IThreadComment>;
  replyCount: number;
  isLiked: boolean;
}
