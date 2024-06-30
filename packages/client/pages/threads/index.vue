<template>
  <div>
    <h2>スレ一覧</h2>

    <v-divider></v-divider>

    <v-autocomplete label="スレを検索" :items="['California', 'Colorado', 'Florida', 'Georgia', 'Texas', 'Wyoming']" v-model="keyword" />
    <ThreadTable title="" :items="newsItems" moreLink="/news" :maxItems="5" />
  </div>
</template>

<script setup>
import Header from '~/components/Header.vue';
import ThreadTable from '~/components/ThreadTable.vue';

const route = useRoute();

const keyword = ref(route.query.keyword ?? '');

const newsItems = ref([
  {
    title:
      'ラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴え',
    subtitle: '食事中にイヤホンやめて',
    link: '/news/1',
    comments: 12,
    board: 'ニュース',
  },
  { title: '都知事選', subtitle: '蓮舫氏、神谷外務次官の是非問う', link: '/news/2', comments: 8, board: 'ニュース' },
  { title: 'ニコニコ動画', subtitle: '流出情報「ダウンロードや拡散はお控えください」', link: '/news/3', comments: 5, board: 'ニュース' },
]);
</script>
