export interface IThreadCommentAttachment {
  url: string;
  displayOrder: number;
  type: 'video' | 'image';
  duration?: number;
}

export interface IThreadCommentAttachmentForThread extends IThreadCommentAttachment {
  commentId: number;
  idx: number;
  commentAuthorName?: string;
  commentContent?: string;
  createdAt?: string;
}
