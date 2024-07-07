<template>
  <div>
    <Menu :items="menuItems" />

    <!-- <ThreadTable title="閲覧履歴" :items="historyItems" moreLink="/history" link="/hoge" :maxItems="3" /> -->
    <ThreadTable title="人気" :items="threadsByPopular" moreLink="/news" :maxItems="5" />
    <ThreadTable title="新着" :items="threadsByNewest" moreLink="/popular" :maxItems="3" />
  </div>
</template>

<script setup lang="ts">
import Menu from '~/components/Menu.vue';
import ThreadTable from '~/components/thread/ThreadTable.vue';
import type { TThread, TThreadList } from '~/types/thread';

const router = useRouter();
const nuxtApp = useNuxtApp();
const { payload, $api } = nuxtApp;
console.log('ユーザー情報:', payload.user);

const threadsByPopular = ref<TThread[]>([]);
const threadsByNewest = ref<TThread[]>([]);

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
  const response = await $api.get<TThreadList>('/threads', {
    params: {
      orders: ['popularity', 'newest'],
    },
  });
  threadsByPopular.value = response.data.threadsByPopular;
  threadsByNewest.value = response.data.threadsByNewest;
}
</script>
