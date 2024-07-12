export default defineNuxtPlugin(nuxtApp => {
  nuxtApp.provide('storage', Storage);
});

export const useThreadViewHistory = () => {
  const { $storage } = useNuxtApp();

  const THREAD_VIEW_HISTORY_KEY = 'threadViewHistory';

  const getThreadViewHistory = (): number[] => {
    return $storage.getItem<number[]>(THREAD_VIEW_HISTORY_KEY) || [];
  };

  const setThreadViewHistory = (threadId: number): void => {
    const history = $storage.getItem<number[]>(THREAD_VIEW_HISTORY_KEY) || [];
    if (!history.includes(threadId)) {
      history.push(threadId);
      $storage.setItem(THREAD_VIEW_HISTORY_KEY, history);
    }
  };

  const clearThreadViewHistory = (): void => {
    $storage.removeItem(THREAD_VIEW_HISTORY_KEY);
  };

  return { getThreadViewHistory, setThreadViewHistory, clearThreadViewHistory };
};
