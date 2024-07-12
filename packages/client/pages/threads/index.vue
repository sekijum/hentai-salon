<template>
  <div>
    <PageTitle title="スレ一覧" />

    <v-divider></v-divider>

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
      v-if="threads?.threadsByHistory"
      title="閲覧履歴"
      :items="threads?.threadsByHistory"
      :navigate="() => router.push({ query: { queryCriteria: ['history'] } })"
    />
    <ThreadList
      v-if="threads?.threadsByPopular"
      title="人気"
      :items="threads?.threadsByPopular"
      :navigate="() => router.push({ query: { queryCriteria: ['popularity'] } })"
    />
    <ThreadList
      v-if="threads?.threadsByNewest"
      title="新着"
      :items="threads?.threadsByNewest"
      :navigate="() => router.push({ query: { queryCriteria: ['newest'] } })"
    />
    <ThreadList
      v-if="threads?.threadsByKeyword"
      title="全板検索"
      :items="threads?.threadsByKeyword"
      :navigate="() => router.push({ query: { queryCriteria: ['keyword'] } })"
    />
    <ThreadList
      v-if="threads?.threadsByRelated"
      title="関連"
      :items="threads?.threadsByRelated"
      :navigate="() => router.push({ query: { queryCriteria: ['related'] } })"
    />
    <ThreadList
      v-if="threads?.threadsByBoard"
      title="板"
      :items="threads?.threadsByBoard"
      :navigate="() => router.push({ query: { queryCriteria: ['board'] } })"
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

const keyword = ref(route.query.keyword ?? '');
interface ThreadResponse {
  threadsByPopular: IThread[];
  threadsByNewest: IThread[];
  threadsByHistory: IThread[];
  threadsByKeyword: IThread[];
  threadsByRelated: IThread[];
  threadsByBoard: IThread[];
}
const nuxtApp = useNuxtApp();
const { payload, $api } = nuxtApp;
const { getThreadViewHistory } = useThreadViewHistory();

const threads = ref<ThreadResponse>();

const menuItems = [
  {
    title: '関連',
    navigate: () => router.push({ query: { queryCriteria: 'related' } }),
    icon: 'mdi-format-list-bulleted',
  },
  {
    title: '人気',
    navigate: () => router.push({ query: { queryCriteria: 'popularity' } }),
    icon: 'mdi-fire',
  },
  {
    title: '関連履歴',
    navigate: () => router.push({ query: { queryCriteria: 'history' } }),
    icon: 'mdi-earth',
  },
  {
    title: '新着',
    navigate: () => router.push({ query: { queryCriteria: 'newest' } }),
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
  if (route.query.boardId) {
    await router.push({ ...{ query: { ...{ boardId: route.query.boardId }, ...{ queryCriteria: 'board' } } } });
  }
  if (route.query.keyword) {
    await router.push({ ...{ query: { ...{ keyword: keyword.value }, ...{ queryCriteria: 'keyword' } } } });
  }

  const response = await $api.get<ThreadResponse>('/threads', {
    params: {
      queryCriteria: route.query.queryCriteria ? [route.query.queryCriteria] : ['popularity', 'newest'],
      threadIds: getThreadViewHistory(),
      keyword: keyword.value,
      boardId: route.query.boardId,
      limit: 10,
    },
  });
  threads.value = response.data;
}

watchEffect(() => {
  if (route.query.queryCriteria) {
    fetchThreads();
  }
});
</script>
