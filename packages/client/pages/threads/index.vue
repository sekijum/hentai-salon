<template>
  <div>
    <PageTitle title="スレ一覧" />

    <v-divider />

    <Menu :items="menuItems" />

    <v-text-field
      v-model="keyword"
      label="スレを検索"
      variant="outlined"
      hide-details
      counter
      single-line
      clearable
      dense
      density="compact"
    />
    <v-btn type="submit" color="primary" block @click="search">全板検索</v-btn>

    <ThreadList
      v-if="route.query.filter === 'history' && threads.threadsByHistory.length"
      filter="history"
      title="閲覧履歴"
      :items="threads?.threadsByHistory"
      :isInfiniteScroll="false"
      :threadLimit="threadLimit"
    />
    <ThreadList
      v-if="(route.query.filter === 'popularity' || !route.query.filter) && threads.threadsByPopular.length"
      filter="popularity"
      title="人気"
      :items="threads?.threadsByPopular"
      :isInfiniteScroll="true"
      :threadLimit="threadLimit"
    />
    <ThreadList
      v-if="(route.query.filter === 'newest' || !route.query.filter) && threads.threadsByNewest.length"
      filter="newest"
      title="新着"
      :items="threads?.threadsByNewest"
      :isInfiniteScroll="true"
      :threadLimit="threadLimit"
    />
    <ThreadList
      v-if="route.query.filter === 'keyword' && threads.threadsByKeyword.length"
      filter="keyword"
      title="全板検索"
      :items="threads?.threadsByKeyword"
      :isInfiniteScroll="true"
      :threadLimit="threadLimit"
    />
    <ThreadList
      v-if="route.query.filter === 'related' && threads.threadsByRelated.length"
      filter="related"
      title="関連"
      :items="threads?.threadsByRelated"
      :isInfiniteScroll="true"
      :threadLimit="threadLimit"
    />
    <ThreadList
      v-if="route.query.filter === 'board' && threads.threadsByBoard.length"
      filter="board"
      :title="`#${board?.title}`"
      :items="threads?.threadsByBoard"
      :isInfiniteScroll="true"
      :threadLimit="threadLimit"
    />
  </div>
</template>

<script setup lang="ts">
import ThreadList from '~/components/thread/ThreadList.vue';
import Menu from '~/components/Menu.vue';
import PageTitle from '~/components/PageTitle.vue';
import type { IThread } from '~/types/thread';
import type { IBoard } from '~/types/board';

const route = useRoute();
const router = useRouter();
const nuxtApp = useNuxtApp();
const { getThreadViewHistory } = useStorage();

const { $api } = nuxtApp;

const keyword = ref(route.query.keyword ?? '');
const threads = ref<{
  threadsByPopular: IThread[];
  threadsByNewest: IThread[];
  threadsByHistory: IThread[];
  threadsByKeyword: IThread[];
  threadsByRelated: IThread[];
  threadsByBoard: IThread[];
}>({
  threadsByPopular: [],
  threadsByNewest: [],
  threadsByHistory: [],
  threadsByKeyword: [],
  threadsByRelated: [],
  threadsByBoard: [],
});
const threadLimit = 10;
const board = ref<IBoard>();

const menuItems = [
  {
    title: 'トップ',
    clicked: () => router.push({ query: {} }),
    icon: 'mdi-home',
  },
  {
    title: '関連',
    clicked: () => router.push({ query: { filter: 'related' } }),
    icon: 'mdi-link-variant',
  },
  {
    title: '人気',
    clicked: () => router.push({ query: { filter: 'popularity' } }),
    icon: 'mdi-fire',
  },
  {
    title: '関連履歴',
    clicked: () => router.push({ query: { filter: 'history' } }),
    icon: 'mdi-history',
  },
  {
    title: '新着',
    clicked: () => router.push({ query: { filter: 'newest' } }),
    icon: 'mdi-new-box',
  },
];

async function search() {
  await router.push({ ...{ query: { ...{ keyword: keyword.value }, ...{ filter: 'keyword' } } } });
  await fetchThreads();
}

onMounted(async () => {
  await fetchThreads();
});

async function fetchThreads() {
  if (route.query.filter) {
    const response = await $api.get<IThread[]>('/threads/', {
      params: {
        filter: route.query.filter,
        threadIds: getThreadViewHistory(),
        keyword: keyword.value,
        boardId: route.query.boardId,
        limit: threadLimit,
      },
    });
    if (route.query.filter === 'history') {
      threads.value.threadsByHistory = response.data;
    } else if (route.query.filter === 'popularity') {
      threads.value.threadsByPopular = response.data;
    } else if (route.query.filter === 'newest') {
      threads.value.threadsByNewest = response.data;
    } else if (route.query.filter === 'keyword') {
      threads.value.threadsByKeyword = response.data;
    } else if (route.query.filter === 'related') {
      threads.value.threadsByRelated = response.data;
    } else if (route.query.filter === 'board') {
      board.value = response.data[0].board;
      threads.value.threadsByBoard = response.data;
    }
  } else {
    await Promise.all(
      ['newest'].map(async filter => {
        const response = await $api.get<IThread[]>('/threads/', {
          params: {
            filter,
            threadIds: getThreadViewHistory(),
            keyword: keyword.value,
            boardId: route.query.boardId,
            limit: threadLimit,
          },
        });
        if (filter === 'popularity') {
          threads.value.threadsByPopular = response.data;
        } else if (filter === 'newest') {
          threads.value.threadsByNewest = response.data;
        }
        console.log(threads.value);
      }),
    );
  }
}

watch(
  () => route.query.filter,
  () => fetchThreads(),
);

useHead({
  title: '変態サロン | スレ一覧',
  meta: [
    {
      property: 'og:title',
      content: '変態サロン | スレ一覧',
    },
    {
      property: 'og:url',
      content: location.href,
    },
  ],
});
</script>
