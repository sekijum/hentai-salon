<template>
  <div>
    <v-alert
      v-if="payload.isLoggedIn"
      :text="`${payload?.user?.name}さん、あなたのアカウントにログインしています。`"
      prominent
      class="small-text text-center"
    />

    <Menu :items="guestMenuItems" />

    <ThreadList
      v-if="threads.threadsByHistory?.length"
      queryCriteria="history"
      title="スレッド閲覧履歴"
      :items="threads.threadsByHistory"
      :clicked="() => router.push({ path: '/threads', query: { queryCriteria: 'history' } })"
      :isInfiniteScroll="false"
    />
    <ThreadList
      v-if="threads.threadsByPopular?.length"
      queryCriteria="popularity"
      title="人気"
      :items="threads.threadsByPopular"
      :clicked="() => router.push({ path: '/threads' })"
      :isInfiniteScroll="false"
    />
    <ThreadList
      v-if="threads.threadsByNewest?.length"
      queryCriteria="newest"
      title="新着"
      :items="threads.threadsByNewest"
      :isInfiniteScroll="true"
    />
  </div>
</template>

<script setup lang="ts">
import Menu from '~/components/Menu.vue';
import ThreadList from '~/components/thread/ThreadList.vue';
import type { IThread } from '~/types/thread';

const router = useRouter();
const nuxtApp = useNuxtApp();
const { getThreadViewHistory } = useStorage();
const isMenuModal = useState('isMenuModal', () => false);

const { payload, $api } = nuxtApp;

const threads = ref<{
  threadsByPopular: IThread[];
  threadsByNewest: IThread[];
  threadsByHistory: IThread[];
}>({
  threadsByPopular: [],
  threadsByNewest: [],
  threadsByHistory: [],
});

const guestMenuItems = computed(() => {
  let items: { title: string; clicked: () => any; icon: string }[] = [
    { title: 'スレ一覧', clicked: () => router.push('/threads'), icon: 'mdi-new-box' },
    { title: '板一覧', clicked: () => router.push('/boards'), icon: 'mdi-format-list-bulleted' },
    { title: '設定', clicked: () => router.push('/setting'), icon: 'mdi-cog' },
    { title: 'メニュー', clicked: () => (isMenuModal.value = true), icon: 'mdi-forum' },
  ];

  if (!payload.isLoggedIn) {
    items = items.concat([
      { title: 'サインイン', clicked: () => router.push('/signin'), icon: 'mdi-login' },
      { title: 'サインアップ', clicked: () => router.push('/signup'), icon: 'mdi-account-plus' },
    ]);
  }

  if (payload.isMember || payload.isAdmin) {
    items = items.concat([
      { title: 'スレ作成', clicked: () => router.push('/threads/new'), icon: 'mdi-forum' },
      { title: 'ユーザー情報', clicked: () => router.push('/users/me'), icon: 'mdi-update' },
      { title: 'マイスレ', clicked: () => router.push('/users/me/threads'), icon: 'mdi-new-box' },
      { title: 'マイレス', clicked: () => router.push('/users/me/comments'), icon: 'mdi-format-list-bulleted' },
    ]);
  }

  if (payload.isAdmin) {
    items = items.concat([
      { title: '板作成', clicked: () => router.push('/boards/new'), icon: 'mdi-forum' },
      { title: 'adminer', clicked: () => open('http://localhost:8081', '_blank'), icon: 'mdi-forum' },
    ]);
  }

  return items;
});

onMounted(async () => {
  await fetchThreads();
});

async function fetchThreads() {
  await Promise.all(
    ['history', 'newest', 'popularity'].map(async queryCriteria => {
      const response = await $api.get<IThread[]>('/threads/', {
        params: {
          queryCriteria,
          threadIds: getThreadViewHistory(),
          limit: 10,
        },
      });
      if (queryCriteria === 'popularity') {
        threads.value.threadsByNewest = response.data;
      } else if (queryCriteria === 'newest') {
        threads.value.threadsByPopular = response.data;
      } else if (queryCriteria === 'history') {
        threads.value.threadsByHistory = response.data;
      }
    }),
  );
}
</script>

<style scoped>
.small-text {
  font-size: 0.75rem;
}
</style>
