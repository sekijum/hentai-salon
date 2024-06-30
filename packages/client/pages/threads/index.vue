<template>
  <div>
    <PageTitle title="スレ一覧" />

    <v-divider></v-divider>

    <MenuSection
      :items="[
        {
          title: '関連順',
          to: '',
          icon: 'mdi-format-list-bulleted',
        },
        {
          title: '人気',
          to: '',
          icon: 'mdi-fire',
        },
        {
          title: '閲覧履歴',
          to: '',
          icon: 'mdi-update',
        },
        {
          title: '閲覧順',
          to: '',
          icon: 'mdi-earth',
        },
        {
          title: '新着順',
          to: '',
          icon: 'mdi-new-box',
        },
        {
          title: 'コメント数',
          to: '',
          icon: 'mdi-cog',
        },
      ]"
    />

    <v-autocomplete label="スレを検索" :items="['California', 'Colorado', 'Florida', 'Georgia', 'Texas', 'Wyoming']" v-model="keyword" />

    <ThreadTable title="" :items="newsItems" moreLink="/news" :maxItems="5" />
  </div>
</template>

<script setup>
import Header from '~/components/Header.vue';
import ThreadTable from '~/components//thread/ThreadTable.vue';
import MenuSection from '~/components//MenuSection.vue';
import PageTitle from '~/components/PageTitle.vue';

const route = useRoute();

const keyword = ref(route.query.keyword ?? '');

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
  { title: 'ニコニコ動画', subtitle: '流出情報「ダウンロードや拡散はお控えください」', link: '/news/3', comments: 5, board: 'ニュース' },
]);
</script>
