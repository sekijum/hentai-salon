export type TThreadBoard = {
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
};

export type TThreadList = {
  threadsByPopular: TThread[];
  threadsByNewest: TThread[];
  threadsByHistories: TThread[];
};
