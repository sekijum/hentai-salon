import type { ICollection } from './collection';
import type { IThreadCommentAttachment } from './thread-comment-attachment';
import type { IUser } from './user';

export interface IThreadComment {
  id: number;
  guestName?: string;
  content: string;
  createdAt: string;
  updatedAt: string;
  attachments: IThreadCommentAttachment[];
  user?: IUser;
  thread: IThreadComment;
  parentCommentId: number;
  parentComment: IThreadComment;
  replies: ICollection<IThreadComment>;
  replyCount: number;
  isLiked: boolean;
}
