<template>
  <div>
    <PageTitle title="マイスレ" />

    <Menu :items="menuItems" />

    <ThreadItem filter="owner" :isInfiniteScroll="true" />
  </div>
</template>

<script setup lang="ts">
import PageTitle from '~/components/PageTitle.vue';
import ThreadItem from '~/components/thread/ThreadItem.vue';
import type { IThread } from '~/types/thread';
import type { ICollection } from '~/types/collection';

definePageMeta({ middleware: ['logged-in-access-only'] });

const router = useRouter();
const nuxtApp = useNuxtApp();
const { $api } = nuxtApp;
const { getThreadViewHistory } = useStorage();
const threadLimit = 10;

const menuItems = [
  { title: 'マイスレ', clicked: () => router.push('/mypage/threads'), icon: 'mdi-file-document-multiple-outline' },
  { title: 'ユーザー情報', clicked: () => router.push('/mypage'), icon: 'mdi-account-cog-outline' },
  { title: 'マイレス', clicked: () => router.push('/mypage/comments'), icon: 'mdi-message-text-outline' },
  {
    title: 'お気に入りスレ',
    clicked: () => router.push('/mypage/liked-threads'),
    icon: 'mdi-star-box-multiple-outline',
  },
  { title: 'お気に入りレス', clicked: () => router.push('/mypage/liked-comments'), icon: 'mdi-message-star-outline' },
];

const threads = ref<ICollection<IThread>>();

onMounted(async () => {
  await fetchThreads();
});

async function fetchThreads() {
  const response = await $api.get<ICollection<IThread>>('/users/me/threads', {
    params: {
      threadIds: getThreadViewHistory(),
      limit: threadLimit,
    },
  });
  threads.value = response.data;
}
</script>
