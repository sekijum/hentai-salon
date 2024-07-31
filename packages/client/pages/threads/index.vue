<template>
  <div>
    <PageTitle title="スレ一覧" />

    <v-divider />

    <Menu :items="menuItems" />

    <v-text-field
      v-model="keyword"
      label="スレを検索"
      variant="outlined"
      hide-details
      counter
      single-line
      clearable
      dense
      density="compact"
    />
    <v-btn type="submit" color="primary" block @click="search">全板検索</v-btn>

    <ThreadItem
      v-if="!route.query.filter || route.query.filter === 'popularity'"
      filter="popularity"
      title="人気"
      :isInfiniteScroll="true"
    />
    <ThreadItem v-if="route.query.filter === 'history'" filter="history" title="閲覧履歴" :isInfiniteScroll="false" />
    <ThreadItem v-if="route.query.filter === 'newest'" filter="newest" title="新着" :isInfiniteScroll="true" />
    <ThreadItem v-if="route.query.filter === 'keyword'" filter="keyword" title="全板検索" :isInfiniteScroll="true" />
    <ThreadItem v-if="route.query.filter === 'related'" filter="related" title="関連" :isInfiniteScroll="true" />
    <ThreadItem v-if="route.query.filter === 'board'" filter="board" title="板" :isInfiniteScroll="true" />
  </div>
</template>

<script setup lang="ts">
import ThreadItem from '~/components/thread/ThreadItem.vue';
import Menu from '~/components/Menu.vue';
import PageTitle from '~/components/PageTitle.vue';

const route = useRoute();
const router = useRouter();
const nuxtApp = useNuxtApp();

const keyword = ref(route.query.keyword ?? '');

const menuItems = [
  {
    title: 'トップ',
    clicked: () => router.push({ query: {} }),
    icon: 'mdi-home',
  },
  {
    title: '関連',
    clicked: () => router.push({ query: { filter: 'related' } }),
    icon: 'mdi-link-variant',
  },
  {
    title: '人気',
    clicked: () => router.push({ query: { filter: 'popularity' } }),
    icon: 'mdi-fire',
  },
  {
    title: '関連履歴',
    clicked: () => router.push({ query: { filter: 'history' } }),
    icon: 'mdi-history',
  },
  {
    title: '新着',
    clicked: () => router.push({ query: { filter: 'newest' } }),
    icon: 'mdi-new-box',
  },
];

async function search() {
  await router.push({ ...{ query: { ...{ keyword: keyword.value }, ...{ filter: 'keyword' } } } });
}

useHead({
  title: '変態サロン | スレ一覧',
  meta: [
    {
      property: 'og:title',
      content: '変態サロン | スレ一覧',
    },
    {
      property: 'og:url',
      content: location.href,
    },
  ],
});
</script>
