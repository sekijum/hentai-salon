<template>
  <div>
    <PageTitle title="マイスレ" />

    <Menu :items="menuItems" />

    <ThreadList v-if="threads.length" queryCriteria="owner" :items="threads" :isInfiniteScroll="true" />
  </div>
</template>

<script setup lang="ts">
import PageTitle from '~/components/PageTitle.vue';
import ThreadList from '~/components/thread/ThreadList.vue';
import type { IThread } from '~/types/thread';

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { $api } = nuxtApp;
const { getThreadViewHistory } = useStorage();

const menuItems = [
  { title: 'ユーザー情報', clicked: () => router.push('/users/me'), icon: 'mdi-update' },
  { title: 'マイスレ', clicked: () => router.push('/users/me/threads'), icon: 'mdi-new-box' },
  { title: 'マイレス', clicked: () => router.push('/users/me/comments'), icon: 'mdi-format-list-bulleted' },
];

const threads = ref<IThread[]>([]);

onMounted(async () => {
  await fetchThreads();
});

async function fetchThreads() {
  const response = await $api.get<IThread[]>('/threads/', {
    params: {
      queryCriteria: 'owner',
      threadIds: getThreadViewHistory(),
      limit: 10,
    },
  });
  console.log(response);
  threads.value = response.data;
}
</script>
