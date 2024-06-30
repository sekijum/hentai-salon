<template>
  <div>
    <h2>{{ item.title }}</h2>

    <v-divider></v-divider>

    <CommentList />
  </div>
</template>

<script setup>
import Header from '~/components/Header.vue';
import ThreadTable from '~/components/thread/ThreadTable.vue';
import CommentList from '~/components/comment/CommentList.vue';

const route = useRoute();

const keyword = ref(route.query.keyword ?? '');

const item = ref({
  title:
    'ラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴え',
  subtitle: '食事中にイヤホンやめて',
  link: '/news/1',
  comments: 12,
  board: 'ニュース',
});
</script>
