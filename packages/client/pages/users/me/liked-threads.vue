<template>
  <div v-if="threads">
    <PageTitle title="お気に入りスレ" />

    <Menu :items="menuItems" />

    <template v-if="threads.data">
      <Pagination :totalCount="threads.totalCount" :limit="threadLimit" />

      <ThreadList filter="owner" :items="threads.data" :isInfiniteScroll="false" :threadLimit="threadLimit" />

      <Pagination :totalCount="threads.totalCount" :limit="threadLimit" />
    </template>
  </div>
</template>

<script setup lang="ts">
import PageTitle from '~/components/PageTitle.vue';
import ThreadList from '~/components/thread/ThreadList.vue';
import type { IThread } from '~/types/thread';
import type { ICollection } from '~/types/collection';

definePageMeta({ middleware: ['logged-in-access-only'] });

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { $api } = nuxtApp;
const { getThreadViewHistory } = useStorage();
const threadLimit = 10;

const menuItems = [
  { title: 'マイスレ', clicked: () => router.push('/users/me/threads'), icon: 'mdi-file-document-multiple-outline' },
  { title: 'ユーザー情報', clicked: () => router.push('/users/me'), icon: 'mdi-account-cog-outline' },
  { title: 'マイレス', clicked: () => router.push('/users/me/comments'), icon: 'mdi-message-text-outline' },
  {
    title: 'お気に入りスレ',
    clicked: () => router.push('/users/me/liked-threads'),
    icon: 'mdi-star-box-multiple-outline',
  },
  { title: 'お気に入りレス', clicked: () => router.push('/users/me/liked-comments'), icon: 'mdi-message-star-outline' },
];

const threads = ref<ICollection<IThread>>();

onMounted(async () => {
  await fetchThreads();
});

async function fetchThreads() {
  const response = await $api.get<ICollection<IThread>>('/users/me/liked-threads', {
    params: {
      threadIds: getThreadViewHistory(),
      limit: threadLimit,
    },
  });
  threads.value = response.data;
}
</script>
