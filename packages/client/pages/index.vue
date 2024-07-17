<template>
  <div>
    <v-alert
      v-if="payload.isLoggedIn"
      :text="`*${payload?.user?.name}*さん、あなたのアカウントにログインしています。`"
      prominent
      class="small-text"
    />

    <Menu :items="menuItems" />

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
    <ThreadList title="新着" :items="threadsByNewest" :isInfiniteScroll="true" />
  </div>
</template>

<script setup lang="ts">
import Menu from '~/components/Menu.vue';
import ThreadList from '~/components/thread/ThreadList.vue';
import type { IThread } from '~/types/thread';

interface ThreadResponse {
  threadsByPopular: IThread[];
  threadsByNewest: IThread[];
  threadsByHistory: IThread[];
}

const router = useRouter();
const nuxtApp = useNuxtApp();
const { getThreadViewHistory } = useStorage();
const isMenuModal = useState('isMenuModal', () => false);

const { payload, $api } = nuxtApp;

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
  { title: 'メニュー', clicked: () => (isMenuModal.value = true), icon: 'mdi-forum' },
];

onMounted(async () => {
  await fetchThreads();
});

async function fetchThreads() {
  const queryCriteria = ['newest', 'popularity'];
  if (getThreadViewHistory().length) {
    queryCriteria.push('history');
  }
  const response = await $api.get<ThreadResponse>('/threads/', {
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

<style scoped>
.small-text {
  font-size: 0.75rem;
}
</style>
