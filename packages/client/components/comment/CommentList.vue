<template>
  <v-list class="no-padding-list">
    <CommentItem v-for="(comment, idx) in comments" :key="comment.id" :comment="comment" :idx="idx" />
  </v-list>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import CommentItem from '~/components/comment/CommentItem.vue';

const comments = ref([
  {
    id: 1,
    username: '警備員',
    level: 6,
    rank: '新芽',
    date: '2024/06/30',
    time: '14:25',
    content: '俺も恋人じゃないけど触らせてもらおう',
    board: '垢版',
    type: '大砲',
    idCode: '+Q6HhT800',
    commentCount: '11',
    media: [{ type: 'image', url: 'https://via.placeholder.com/300' }],
  },
  {
    id: 2,
    username: '警備員',
    level: 17,
    rank: '新芽',
    date: '2024/06/30',
    time: '14:26',
    content: 'もう結婚するしか逃げ道はない',
    board: '垢版',
    type: '大砲',
    idCode: 'X+z5zrq00',
    commentCount: '11',
    media: [
      { type: 'image', url: 'https://via.placeholder.com/300' },
      {
        type: 'video/mp4',
        url: 'https://www.w3schools.com/html/mov_bbb.mp4',
        thumbnail: 'https://via.placeholder.com/300',
      },
    ],
  },
  {
    id: 3,
    username: '名無しさん@恐縮です',
    date: '2024/06/30',
    time: '14:28',
    content: 'もう飽きたこの話題',
    board: '垢版',
    type: '大砲',
    idCode: '6C/Ty2sa0',
    commentCount: '11',
    media: [],
  },
  {
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
  },
]);
</script>

<style scoped>
.no-padding-list {
  padding: 0;
}
</style>
