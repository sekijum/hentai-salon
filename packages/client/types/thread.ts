import type { ICollection } from './collection';
import type { IThreadCommentAttachment, IThreadCommentAttachmentForThread } from './thread-comment-attachment';

interface IThreadBoard {
  id: number;
  title: string;
}

export interface IThreadComment {
  id: number;
  user?: { id: number; name: string; profileLink?: string };
  guestName?: string;
  content: string;
  parentCommentId?: number;
  attachments: IThreadCommentAttachment[];
  replyCount: number;
  isLiked: boolean;
}

export interface IThread {
  id: number;
  board?: IThreadBoard;
  title: string;
  description: string;
  thumbnailUrl: string;
  tagNameList: string[];
  commentCount: number;
  comments: ICollection<IThreadComment>;
  attachments: ICollection<IThreadCommentAttachmentForThread>;
  isLiked: boolean;
}
