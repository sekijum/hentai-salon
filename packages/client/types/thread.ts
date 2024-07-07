type ThreadCommentResource = {
  id: number;
  userId?: number;
  guestName?: string;
  content: string;
  ipAddress: string;
  status: number;
  parentCommentID: number;
  createdAt: string;
  updatedAt: string;
  attachments: ThreadCommentAttachmentResource[];
  totalReplies: number;
};

type ThreadCommentAttachmentResource = {
  url: string;
  displayOrder: number;
  type: string;
};

type TThreadBoard = {
  id: number;
  title: string;
};

export type TThread = {
  id: number;
  board: TThreadBoard | null;
  title: string;
  description: string;
  thumbnailUrl: string;
  tags: string[];
  createdAt: string;
  commentCount: number;
  popularity: string;
  comments: ThreadCommentResource;
};

export type TThreadList = {
  threadsByPopular: TThread[];
  threadsByNewest: TThread[];
  threadsByHistories: TThread[];
};
