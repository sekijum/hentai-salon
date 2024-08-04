export interface IThreadCommentAttachment {
  url: string;
  displayOrder: number;
  type: 'Video' | 'Image';
  duration?: number;
}

export interface IThreadCommentAttachmentForThread extends IThreadCommentAttachment {
  commentId: number;
  idx: number;
}
