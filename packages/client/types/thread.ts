import type { IListResource } from './list-resource';
import type { IThreadCommentAttachment, IThreadCommentAttachmentForThread } from './thread-comment-attachment';

interface IThreadBoard {
  id: number;
  title: string;
}

export interface IThreadComment {
  id: number;
  idx: number;
  user?: { id: number; name: string; profileLink?: string };
  guestName?: string;
  content: string;
  parentCommentIdx?: number;
  parentCommentId: number;
  createdAt: string;
  updatedAt: string;
  attachments: IThreadCommentAttachment[];
  totalReplies: number;
}

export interface IThread {
  id: number;
  board?: IThreadBoard;
  title: string;
  description: string;
  thumbnailUrl: string;
  tags: string[];
  createdAt: string;
  commentCount: number;
  comments: IListResource<IThreadComment>;
  attachments: IThreadCommentAttachmentForThread[];
}
