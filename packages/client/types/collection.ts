export interface ICollection<T> {
  totalCount: number;
  limit: number;
  offset: number;
  data: T[];
}
