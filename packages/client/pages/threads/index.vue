<template>
  <div>
    <PageTitle title="スレ一覧" />

    <v-divider></v-divider>

    <v-autocomplete
      label="スレを検索"
      :items="['California', 'Colorado', 'Florida', 'Georgia', 'Texas', 'Wyoming']"
      v-model="keyword"
    />

    <Menu :items="menuItems" />

    <ThreadTable title="" :items="newsItems" moreLink="/news" :maxItems="5" />
  </div>
</template>

<script setup lang="ts">
import Header from '~/components/Header.vue';
import ThreadTable from '~/components/thread/ThreadTable.vue';
import Menu from '~/components/Menu.vue';
import PageTitle from '~/components/PageTitle.vue';

const route = useRoute();
import { useRouter } from 'vue-router';

const router = useRouter();
const keyword = ref(route.query.keyword ?? '');

const menuItems = [
  {
    title: '関連順',
    navigate: () => router('/'),
    icon: 'mdi-format-list-bulleted',
  },
  {
    title: '人気',
    navigate: () => router('/'),
    icon: 'mdi-fire',
  },
  {
    title: '閲覧履歴',
    navigate: () => router('/'),
    icon: 'mdi-update',
  },
  {
    title: '閲覧順',
    navigate: () => router('/'),
    icon: 'mdi-earth',
  },
  {
    title: '新着順',
    navigate: () => router('/'),
    icon: 'mdi-new-box',
  },
  {
    title: 'コメント数',
    navigate: () => router('/'),
    icon: 'mdi-cog',
  },
];

const newsItems = ref([
  {
    title:
      'ラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴え',
    subtitle: '食事中にイヤホンやめて',
    link: '/threads/1',
    comments: 12,
    board: 'ニュース',
  },
  { title: '都知事選', subtitle: '蓮舫氏、神谷外務次官の是非問う', link: '/news/2', comments: 8, board: 'ニュース' },
  {
    title: 'ニコニコ動画',
    subtitle: '流出情報「ダウンロードや拡散はお控えください」',
    link: '/news/3',
    comments: 5,
    board: 'ニュース',
  },
]);
</script>
