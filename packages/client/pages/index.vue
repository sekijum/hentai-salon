<template>
  <div>
    <v-autocomplete label="話題を検索" :items="['California', 'Colorado', 'Florida', 'Georgia', 'Texas', 'Wyoming']"></v-autocomplete>
    <MenuSection />
    <ThreadTable title="閲覧履歴" :items="historyItems" moreLink="/history" :maxItems="3" />
    <ThreadTable title="人気" :items="newsItems" moreLink="/news" :maxItems="5" />
    <ThreadTable title="新着" :items="popularItems" moreLink="/popular" :maxItems="3" />
  </div>
</template>

<script setup>
import Header from '~/components/Header.vue';
import ThreadTable from '~/components/ThreadTable.vue';
import { ref } from 'vue';

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

const popularItems = ref([
  { title: 'こいせん7 全レス転載禁止', link: '/popular/1', comments: 30, board: '人気' },
  { title: 'とらせん8', link: '/popular/2', comments: 25, board: '人気' },
  { title: '巨専9', link: '/popular/3', comments: 20, board: '人気' },
  { title: 'こいせん祝勝会 全レス転載禁止', link: '/popular/4', comments: 15, board: '人気' },
  { title: 'こいせん4 全レス転載禁止', link: '/popular/5', comments: 10, board: '人気' },
]);

const historyItems = ref([
  { title: 'スレッドA', link: '/history/1', comments: 3, board: '閲覧履歴' },
  { title: 'スレッドB', link: '/history/2', comments: 4, board: '閲覧履歴' },
  { title: 'スレッドC', link: '/history/3', comments: 2, board: '閲覧履歴' },
]);
</script>
