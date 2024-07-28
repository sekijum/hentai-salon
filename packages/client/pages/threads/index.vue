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
      v-if="route.query.queryCriteria === 'history' && threads.threadsByHistory.length"
      queryCriteria="history"
      title="閲覧履歴"
      :items="threads?.threadsByHistory"
      :isInfiniteScroll="true"
      :threadLimit="threadLimit"
    />
    <ThreadList
      v-if="
        (route.query.queryCriteria === 'popularity' || !route.query.queryCriteria) && threads.threadsByPopular.length
      "
      queryCriteria="popularity"
      title="人気"
      :items="threads?.threadsByPopular"
      :isInfiniteScroll="true"
      :threadLimit="threadLimit"
    />
    <ThreadList
      v-if="(route.query.queryCriteria === 'newest' || !route.query.queryCriteria) && threads.threadsByNewest.length"
      queryCriteria="newest"
      title="新着"
      :items="threads?.threadsByNewest"
      :isInfiniteScroll="true"
      :threadLimit="threadLimit"
    />
    <ThreadList
      v-if="route.query.queryCriteria === 'keyword' && threads.threadsByKeyword.length"
      queryCriteria="keyword"
      title="全板検索"
      :items="threads?.threadsByKeyword"
      :isInfiniteScroll="true"
      :threadLimit="threadLimit"
    />
    <ThreadList
      v-if="route.query.queryCriteria === 'related' && threads.threadsByRelated.length"
      queryCriteria="related"
      title="関連"
      :items="threads?.threadsByRelated"
      :isInfiniteScroll="true"
      :threadLimit="threadLimit"
    />
    <ThreadList
      v-if="route.query.queryCriteria === 'board' && threads.threadsByBoard.length"
      queryCriteria="board"
      title="板"
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

const menuItems = [
  {
    title: 'トップ',
    clicked: () => router.push({ query: {} }),
    icon: 'mdi-home',
  },
  {
    title: '関連',
    clicked: () => router.push({ query: { queryCriteria: 'related' } }),
    icon: 'mdi-link-variant',
  },
  {
    title: '人気',
    clicked: () => router.push({ query: { queryCriteria: 'popularity' } }),
    icon: 'mdi-fire',
  },
  {
    title: '関連履歴',
    clicked: () => router.push({ query: { queryCriteria: 'history' } }),
    icon: 'mdi-history',
  },
  {
    title: '新着',
    clicked: () => router.push({ query: { queryCriteria: 'newest' } }),
    icon: 'mdi-new-box',
  },
];

async function search() {
  await router.push({ ...{ query: { ...{ keyword: keyword.value }, ...{ queryCriteria: 'keyword' } } } });
  await fetchThreads();
}

onMounted(async () => {
  await fetchThreads();
});

async function fetchThreads() {
  if (route.query.queryCriteria) {
    const response = await $api.get<IThread[]>('/threads/', {
      params: {
        queryCriteria: route.query.queryCriteria,
        threadIds: getThreadViewHistory(),
        keyword: keyword.value,
        boardId: route.query.boardId,
        limit: threadLimit,
      },
    });
    if (route.query.queryCriteria === 'history') {
      threads.value.threadsByHistory = response.data;
    } else if (route.query.queryCriteria === 'popularity') {
      threads.value.threadsByPopular = response.data;
    } else if (route.query.queryCriteria === 'newest') {
      threads.value.threadsByNewest = response.data;
    } else if (route.query.queryCriteria === 'keyword') {
      threads.value.threadsByKeyword = response.data;
    } else if (route.query.queryCriteria === 'related') {
      threads.value.threadsByRelated = response.data;
    } else if (route.query.queryCriteria === 'board') {
      threads.value.threadsByBoard = response.data;
    }
  } else {
    await Promise.all(
      ['newest', 'popularity'].map(async queryCriteria => {
        const response = await $api.get<IThread[]>('/threads/', {
          params: {
            queryCriteria,
            threadIds: getThreadViewHistory(),
            keyword: keyword.value,
            boardId: route.query.boardId,
            limit: threadLimit,
          },
        });
        if (queryCriteria === 'popularity') {
          threads.value.threadsByPopular = response.data;
        } else if (queryCriteria === 'newest') {
          threads.value.threadsByNewest = response.data;
        }
      }),
    );
  }
}

watch(
  () => route.query.queryCriteria,
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
