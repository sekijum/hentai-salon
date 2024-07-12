<template>
  <div>
    <PageTitle title="スレ一覧" />

    <v-divider></v-divider>

    <Menu :items="menuItems" />

    <v-autocomplete
      label="スレを検索"
      :items="['California', 'Colorado', 'Florida', 'Georgia', 'Texas', 'Wyoming']"
      v-model="keyword"
    />

    <ThreadTable title="人気" :items="threadsByPopular" moreLink="/news" />
    <ThreadTable title="新着" :items="threadsByNewest" moreLink="/popular" />
  </div>
</template>

<script setup lang="ts">
import ThreadTable from '~/components/thread/ThreadTable.vue';
import Menu from '~/components/Menu.vue';
import PageTitle from '~/components/PageTitle.vue';
import type { IThread } from '~/types/thread';
const route = useRoute();
const router = useRouter();

const keyword = ref(route.query.keyword ?? '');
interface ThreadResponse {
  threadsByPopular: IThread[];
  threadsByNewest: IThread[];
}
const nuxtApp = useNuxtApp();
const { payload, $api } = nuxtApp;

const threadsByPopular = ref<IThread[]>([]);
const threadsByNewest = ref<IThread[]>([]);

const menuItems = [
  {
    title: '関連順',
    navigate: () => router.push('/'),
    icon: 'mdi-format-list-bulleted',
  },
  {
    title: '人気',
    navigate: () => router.push('/'),
    icon: 'mdi-fire',
  },
  {
    title: '閲覧履歴',
    navigate: () => router.push('/'),
    icon: 'mdi-update',
  },
  {
    title: '閲覧順',
    navigate: () => router.push('/'),
    icon: 'mdi-earth',
  },
  {
    title: '新着順',
    navigate: () => router.push('/'),
    icon: 'mdi-new-box',
  },
  {
    title: 'コメント数',
    navigate: () => router.push('/'),
    icon: 'mdi-cog',
  },
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
