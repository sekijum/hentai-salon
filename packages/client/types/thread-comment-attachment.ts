export interface IThreadCommentAttachment {
  url: string;
  displayOrder: number;
  type: 'Video' | 'Image';
}

export interface IThreadCommentAttachmentForThread extends IThreadCommentAttachment {
  commentId: number;
  idx: number;
}
