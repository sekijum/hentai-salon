import type { ICollection } from './collection';
import type { IThreadCommentAttachment, IThreadCommentAttachmentForThread } from './thread-comment-attachment';
import type { IBoard } from './board';
import type { IThreadComment } from './thread-comment';

export interface IThread {
  id: number;
  board?: IBoard;
  title: string;
  description: string;
  thumbnailUrl: string;
  tagNameList: string[];
  commentCount: number;
  comments: ICollection<IThreadComment>;
  attachments: ICollection<IThreadCommentAttachmentForThread>;
  isLiked: boolean;
}
