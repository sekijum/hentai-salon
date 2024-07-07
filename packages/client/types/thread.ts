export type TThreadComment = {
  id: number;
  userId?: number;
  guestName?: string;
  content: string;
  ipAddress: string;
  status: number;
  parentCommentID: number;
  createdAt: string;
  updatedAt: string;
  attachments: TThreadCommentAttachment[];
  totalReplies: number;
};

type TThreadCommentAttachment = {
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
  comments: TThreadComment[];
};

export type TThreadList = {
  threadsByPopular: TThread[];
  threadsByNewest: TThread[];
  threadsByHistories: TThread[];
};
