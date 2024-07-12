<template>
  <div>
    <PageTitle title="コメント返信" />

    <v-divider></v-divider>

    <CommentItem :comment="comment" />

    <CommentForm formTitle="返信する" />

    <br />

    <v-divider></v-divider>

    <div id="comment-top" />
    <CommentList :comments="comments" />
    <div id="comment-bottom" />

    <Pagination />

    <br />

    <CommentForm formTitle="返信する" />

    <!-- 上にスクロールするFAB -->
    <v-btn icon large color="primary" class="fab fab-top" @click="scrollToCommentTop">
      <v-icon>mdi-arrow-up</v-icon>
    </v-btn>

    <!-- 下にスクロールするFAB -->
    <v-btn icon large color="primary" class="fab fab-bottom" @click="scrollToCommentBottom">
      <v-icon>mdi-arrow-down</v-icon>
    </v-btn>
  </div>
</template>

<script setup lang="ts">
import Header from '~/components/Header.vue';
import ThreadList from '~/components/thread/ThreadList.vue';
import CommentList from '~/components/comment/CommentList.vue';
import CommentForm from '~/components/comment/CommentForm.vue';
import Menu from '~/components/Menu.vue';
import PageTitle from '~/components/PageTitle.vue';
import Pagination from '~/components/Pagination.vue';
import MediaGallery from '~/components/MediaGallery.vue';
import { useRouter, useRoute } from 'vue-router';

const router = useRouter();
const route = useRoute();
const keyword = ref(route.query.keyword ?? '');
const tab = ref(route.query.tab ?? 'comments');

const menuItems = [
  {
    title: 'コメント一覧',
    navigate: () => router.replace({ query: {} }),
    icon: 'mdi-fire',
  },
  {
    title: 'メディア',
    navigate: () => router.replace({ query: { tab: 'media' } }),
    icon: 'mdi-update',
  },
];

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
