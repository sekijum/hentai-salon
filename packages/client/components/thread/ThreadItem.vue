<template>
  <div v-if="items.length">
    <h2 v-if="title" class="font-weight-regular">{{ title }}</h2>

    <div class="thread-section">
      <template v-if="isInfiniteScroll">
        <v-infinite-scroll :onLoad="load" height="600px">
          <ThreadItemRow :items="items" />
          <template v-slot:empty />
        </v-infinite-scroll>
      </template>
      <template v-else>
        <ThreadItemRow :items="items" />

        <div v-if="clicked" class="more-link" @click="clicked">
          {{ title }}をもっと見る<v-icon down>mdi-chevron-down</v-icon>
        </div>
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import ThreadItemRow from '~/components/thread/ThreadItemRow.vue';
import type { IThread } from '~/types/thread';

const nuxtApp = useNuxtApp();

const props = defineProps<{
  title?: string;
  clicked?: () => void;
  isInfiniteScroll?: boolean;
  filter: string;
  limit?: number;
}>();

const router = useRouter();
const route = useRoute();

const { $api } = nuxtApp;
const { getThreadViewHistory } = useStorage();
const threadLimit = 10;

const offset = ref(0);
const items = ref<IThread[]>([]);

async function load({ done }: { done: (status: 'loading' | 'error' | 'empty' | 'ok') => void }) {
  offset.value += threadLimit;
  const { canNextLoad } = await fetchThreads(offset.value);
  canNextLoad ? done('ok') : done('empty');
}

async function fetchThreads(offset: number = 0) {
  let threads: IThread[];
  if (props.filter === 'liked') {
    const response = await $api.get<IThread[]>('/users/me/liked-threads', {
      params: {
        threadIds: getThreadViewHistory(),
        limit: threadLimit,
        offset: offset,
      },
    });
    threads = response.data;
  } else if (props.filter === 'owner') {
    const response = await $api.get<IThread[]>('/users/me/threads', {
      params: {
        threadIds: getThreadViewHistory(),
        limit: threadLimit,
        offset: offset,
      },
    });
    threads = response.data;
  } else {
    const response = await $api.get<IThread[]>('/threads', {
      params: {
        filter: props.filter,
        threadIds: getThreadViewHistory(),
        keyword: route.query.keyword,
        tagNameList: Array.isArray(route.query.tagNameList) ? route.query.tagNameList : [route.query.tagNameList],
        boardId: route.query.boardId,
        limit: threadLimit,
        offset: offset,
      },
    });
    if (props.limit) {
      threads = response.data.slice(0, props.limit);
    } else {
      threads = response.data;
    }
  }

  if (threads) {
    items.value = [...items.value, ...threads];
  }
  if (!threads || threads.length < threadLimit) {
    return { canNextLoad: false };
  }
  return { canNextLoad: true };
}

watch(
  () => route.query.tagNameList?.length || route.query.keyword || route.query.filter,
  () => {
    offset.value = 0;
    items.value = [];
    fetchThreads();
  },
);

onMounted(() => {
  fetchThreads();
});
</script>

<style scoped>
.thread-section {
  cursor: pointer;
  width: 100%;
}

.more-link {
  text-align: center;
  cursor: pointer;
  padding: 10px;
  text-decoration: underline;
}

:global(.v-infinite-scroll__side) {
  display: none !important;
}
</style>
