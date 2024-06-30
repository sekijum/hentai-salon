<template>
  <div>
    <h1 class="thread-title">{{ item.title }}</h1>

    <v-divider></v-divider>

    <v-chip-group>
      <v-chip>Chip 1</v-chip>

      <v-chip>Chip 2</v-chip>

      <v-chip>Chip 3</v-chip>
    </v-chip-group>

    <v-divider></v-divider>

    <CommentForm />
    <CommentList />
    <CommentForm />
  </div>
</template>

<script setup>
import Header from '~/components/Header.vue';
import ThreadTable from '~/components/thread/ThreadTable.vue';
import CommentList from '~/components/comment/CommentList.vue';
import CommentForm from '~/components/comment/CommentForm.vue';

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

<style scoped>
.thread-title {
  font-size: 1rem;
  text-align: center;
}
</style>
