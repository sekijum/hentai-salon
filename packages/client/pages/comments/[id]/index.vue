<template>
  <div>
    <PageTitle title="コメント詳細" />

    <v-divider />

    <CommentItem :comment="comment" />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import CommentItem from '~/components/comment/CommentItem.vue';

const route = useRoute();
const keyword = ref(route.query.keyword ?? '');

const item = ref({
  title: 'ラーメン店主異例の訴え',
  subtitle: '食事中にイヤホンやめて',
  link: '/news/1',
  comments: 12,
  board: 'ニュース',
});

const comment = ref({
  id: 4,
  username: '名無しさん@恐縮です',
  date: '2024/06/30',
  time: '14:28',
  content: '歌舞伎俳優も未成年の舞妓と疑惑があっても伝統芸能の嗜みでOKのナゾのルールがあったみたいなものかw',
  board: '垢版',
  type: '大砲',
  idCode: 'mswGlHt',
  commentCount: '11',
  media: [
    { type: 'image', url: 'https://via.placeholder.com/300' },
    { type: 'image', url: 'https://via.placeholder.com/300' },
    { type: 'image', url: 'https://via.placeholder.com/300' },
    {
      type: 'video/mp4',
      url: 'https://www.w3schools.com/html/mov_bbb.mp4',
      thumbnail: 'https://via.placeholder.com/300',
    },
  ],
});
</script>

<style scoped>
.fab {
  position: fixed;
  right: 16px;
}

.fab-top {
  bottom: 72px; /* 下のFABとの間にスペースを確保 */
}

.fab-bottom {
  bottom: 16px;
}
</style>
