<template>
  <div>
    <Menu :items="menuItems" />

    <!-- <ThreadList title="閲覧履歴" :items="historyItems" moreLink="/history" link="/hoge" :maxItems="3" /> -->
    <ThreadList
      title="スレッド閲覧履歴"
      :items="threadsByHistory"
      :clicked="() => router.push({ path: '/threads', query: { queryCriteria: 'history' } })"
      :isInfiniteScroll="false"
    />
    <ThreadList
      title="人気"
      :items="threadsByPopular"
      :clicked="() => router.push({ path: '/threads' })"
      :isInfiniteScroll="false"
    />
    <ThreadList
      title="新着"
      :items="threadsByNewest"
      :clicked="() => router.push({ path: '/threads', query: { queryCriteria: 'newest' } })"
      :isInfiniteScroll="false"
    />
  </div>
</template>

<script setup lang="ts">
import Menu from '~/components/Menu.vue';
import ThreadList from '~/components/thread/ThreadList.vue';
import type { IThread } from '~/types/thread';
const { getThreadViewHistory } = useStorage();

interface ThreadResponse {
  threadsByPopular: IThread[];
  threadsByNewest: IThread[];
  threadsByHistory: IThread[];
}

const router = useRouter();
const nuxtApp = useNuxtApp();
const { payload, $api } = nuxtApp;
console.log('ユーザー情報:', payload.user);

const threadsByPopular = ref<IThread[]>([]);
const threadsByNewest = ref<IThread[]>([]);
const threadsByHistory = ref<IThread[]>([]);

const menuItems = [
  { title: 'お知らせ', clicked: () => router.push('/'), icon: 'mdi-update' },
  { title: 'スレ一覧', clicked: () => router.push('/threads'), icon: 'mdi-new-box' },
  { title: '板一覧', clicked: () => router.push('/boards'), icon: 'mdi-format-list-bulleted' },
  { title: '設定', clicked: () => router.push('/setting'), icon: 'mdi-cog' },
  { title: 'サインイン', clicked: () => router.push('/signin'), icon: 'mdi-login' },
  { title: 'サインアップ', clicked: () => router.push('/signup'), icon: 'mdi-account-plus' },
  { title: 'サインアウト', clicked: () => router.push('/signup'), icon: 'mdi-logout' },
  { title: 'スレ作成', clicked: () => router.push('/threads/new'), icon: 'mdi-forum' },
  { title: '板作成', clicked: () => router.push('/boards/new'), icon: 'mdi-forum' },
  { title: '管理画面', clicked: () => router.push('/admin'), icon: 'mdi-forum' },
];

onMounted(async () => {
  await fetchThreads();
});

async function fetchThreads() {
  const queryCriteria = ['newest', 'popularity'];
  if (getThreadViewHistory().length) {
    queryCriteria.push('history');
  }
  console.log(queryCriteria);
  const response = await $api.get<ThreadResponse>('/threads', {
    params: {
      queryCriteria,
      threadIds: getThreadViewHistory(),
      limit: 10,
    },
  });
  threadsByPopular.value = response.data.threadsByPopular;
  threadsByNewest.value = response.data.threadsByNewest;
  threadsByHistory.value = response.data.threadsByHistory;
}
</script>
