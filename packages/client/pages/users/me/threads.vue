<template>
  <div v-if="threads">
    <PageTitle title="マイスレ" />

    <Menu :items="menuItems" />

    <Pagination :totalCount="threads.totalCount" :limit="limit" />

    <ThreadList queryCriteria="owner" :items="threads.data" :isInfiniteScroll="false" />

    <Pagination :totalCount="threads.totalCount" :limit="limit" />
  </div>
</template>

<script setup lang="ts">
import PageTitle from '~/components/PageTitle.vue';
import ThreadList from '~/components/thread/ThreadList.vue';
import type { IThread } from '~/types/thread';
import type { IListResource } from '~/types/list-resource';

definePageMeta({ middleware: ['logged-in-access-only'] });

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { $api } = nuxtApp;
const { getThreadViewHistory } = useStorage();
const limit = 10;

const menuItems = [
  { title: 'ユーザー情報', clicked: () => router.push('/users/me'), icon: 'mdi-account' },
  { title: 'マイスレ', clicked: () => router.push('/users/me/threads'), icon: 'mdi-note' },
  { title: 'マイレス', clicked: () => router.push('/users/me/comments'), icon: 'mdi-comment' },
];

const threads = ref<IListResource<IThread>>();

onMounted(async () => {
  await fetchThreads();
});

async function fetchThreads() {
  const response = await $api.get<IListResource<IThread>>('/users/me/threads', {
    params: {
      threadIds: getThreadViewHistory(),
      limit,
    },
  });
  threads.value = response.data;
}
</script>
