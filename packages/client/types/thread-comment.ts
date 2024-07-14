import type { IListResource } from './list-resource';
import type { IThreadCommentAttachment } from './thread-comment-attachment';
import type { IThreadComment as IComment } from './thread';

export interface IThreadComment {
  id: number;
  guestName?: string;
  content: string;
  createdAt: string;
  updatedAt: string;
  attachments: IThreadCommentAttachment[];
  user?: { id: number; name: string; profileLink?: string };
  thread: {
    id: number;
    title: string;
    description: string;
  };
  parentCommentId: number;
  parentComment: IComment;
  replies: IListResource<IComment>;
  totalReplies: number;
}
