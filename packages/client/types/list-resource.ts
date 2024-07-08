export interface IListResource<T> {
  totalCount: number;
  limit: number;
  offset: number;
  data: T[];
}
