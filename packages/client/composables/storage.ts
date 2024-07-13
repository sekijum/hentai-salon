export const useStorage = () => {
  const { $storage } = useNuxtApp();

  const THREAD_VIEW_HISTORY_KEY = 'thread_view_history';
  const COMMENT_SORT_ORDER = 'comment_sort_order';
  const COMMENT_LIMIT = 'comment_limit';
  const THEME = 'theme';

  const getTheme = (): 'dark' | 'light' => {
    return $storage.getItem<'dark' | 'light'>(THEME) || 'light';
  };

  const setTheme = (theme: 'dark' | 'light'): void => {
    $storage.setItem(THEME, theme);
  };

  const getCommentLimit = (): boolean => {
    return $storage.getItem<boolean>(COMMENT_LIMIT) || false;
  };

  const setCommentLimit = (isBool: boolean): void => {
    $storage.setItem(COMMENT_LIMIT, isBool);
  };

  const getCommentSortOrder = (): 'asc' | 'desc' => {
    return $storage.getItem<'asc' | 'desc'>(COMMENT_SORT_ORDER) || 'desc';
  };

  const setCommentSortOrder = (order: 'asc' | 'desc'): void => {
    $storage.setItem(COMMENT_SORT_ORDER, order);
  };

  const getThreadViewHistory = (): number[] => {
    return $storage.getItem(THREAD_VIEW_HISTORY_KEY) || [];
  };

  const setThreadViewHistory = (threadId: number): void => {
    const history = getThreadViewHistory() || [];
    const index = history.indexOf(threadId);
    if (index > -1) {
      history.splice(index, 1);
    }
    history.unshift(threadId);
    $storage.setItem(THREAD_VIEW_HISTORY_KEY, history);
  };

  const clearThreadViewHistory = (): void => {
    if (confirm('本当にスレッドの閲覧履歴を削除しますか？')) {
      $storage.removeItem(THREAD_VIEW_HISTORY_KEY);
    }
  };

  return {
    getThreadViewHistory,
    setThreadViewHistory,
    clearThreadViewHistory,
    getCommentSortOrder,
    setCommentSortOrder,
    getCommentLimit,
    setCommentLimit,
    getTheme,
    setTheme,
  };
};
