<template>
  <div>
    <Menu :items="menuItems" />

    <!-- <ThreadTable title="閲覧履歴" :items="historyItems" moreLink="/history" link="/hoge" :maxItems="3" /> -->
    <ThreadTable title="人気" :items="threadsByPopular" moreLink="/news" />
    <ThreadTable title="新着" :items="threadsByNewest" moreLink="/popular" />
  </div>
</template>

<script setup lang="ts">
import Menu from '~/components/Menu.vue';
import ThreadTable from '~/components/thread/ThreadTable.vue';
import type { IThread } from '~/types/thread';

interface ThreadResponse {
  threadsByPopular: IThread[];
  threadsByNewest: IThread[];
}

const router = useRouter();
const nuxtApp = useNuxtApp();
const { payload, $api } = nuxtApp;
console.log('ユーザー情報:', payload.user);

const threadsByPopular = ref<IThread[]>([]);
const threadsByNewest = ref<IThread[]>([]);

const menuItems = [
  { title: 'お知らせ', navigate: () => router.push('/'), icon: 'mdi-update' },
  { title: 'スレ一覧', navigate: () => router.push('/threads'), icon: 'mdi-new-box' },
  { title: '板一覧', navigate: () => router.push('/boards'), icon: 'mdi-format-list-bulleted' },
  { title: '設定', navigate: () => router.push('/setting'), icon: 'mdi-cog' },
  { title: 'サインイン', navigate: () => router.push('/signin'), icon: 'mdi-login' },
  { title: 'サインアップ', navigate: () => router.push('/signup'), icon: 'mdi-account-plus' },
  { title: 'サインアウト', navigate: () => router.push('/signup'), icon: 'mdi-logout' },
  { title: 'スレ作成', navigate: () => router.push('/threads/new'), icon: 'mdi-forum' },
  { title: '板作成', navigate: () => router.push('/boards/new'), icon: 'mdi-forum' },
  { title: '管理画面', navigate: () => router.push('/admin'), icon: 'mdi-forum' },
];

onMounted(async () => {
  await fetchThreads();
});

async function fetchThreads() {
  const response = await $api.get<ThreadResponse>('/threads', {
    params: {
      orders: ['popularity', 'newest'],
      limit: 10,
    },
  });
  threadsByPopular.value = response.data.threadsByPopular;
  threadsByNewest.value = response.data.threadsByNewest;
}
</script>
