<template>
  <div>
    <PageTitle :title="item.title" />

    <v-divider></v-divider>

    <v-chip-group active-class="primary--text" column>
      <v-chip
        small
        v-for="tag in [
          'Work',
          'Home Improvement',
          'Vacation',
          'Food',
          'Drawers',
          'Shopping',
          'Art',
          'Tech',
          'Creative Writing',
        ]"
        :key="tag"
      >
        {{ tag }}
      </v-chip>
    </v-chip-group>

    <v-divider></v-divider>

    <Menu :items="menuItems" />

    <template v-if="route.query.tab === 'media'">
      <div id="media-top" />
      <MediaGallery />

      <v-btn icon large color="primary" class="fab fab-bottom" @click="scrollToMediaTop">
        <v-icon>mdi-arrow-up</v-icon>
      </v-btn>
    </template>
    <template v-else>
      <CommentForm />

      <br />

      <v-divider></v-divider>

      <div id="comment-top" />
      <CommentList :comments="comments" />
      <div id="comment-bottom" />

      <Pagination />

      <br />

      <CommentForm />

      <!-- 上にスクロールするFAB -->
      <v-btn icon large color="primary" class="fab fab-top" @click="scrollToCommentTop">
        <v-icon>mdi-arrow-up</v-icon>
      </v-btn>

      <!-- 下にスクロールするFAB -->
      <v-btn icon large color="primary" class="fab fab-bottom" @click="scrollToCommentBottom">
        <v-icon>mdi-arrow-down</v-icon>
      </v-btn>
    </template>
  </div>
</template>

<script setup lang="ts">
import Header from '~/components/Header.vue';
import ThreadTable from '~/components/thread/ThreadTable.vue';
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

const item = ref({
  title: 'ラーメン店主異例の訴え',
  subtitle: '食事中にイヤホンやめて',
  link: '/news/1',
  comments: 12,
  board: 'ニュース',
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

const scrollToMediaTop = () => {
  const mediaTop = document.getElementById('media-top');
  if (mediaTop) {
    mediaTop.scrollIntoView({ behavior: 'smooth' });
  }
};

const scrollToCommentTop = () => {
  const commentTop = document.getElementById('comment-top');
  if (commentTop) {
    commentTop.scrollIntoView({ behavior: 'smooth' });
  }
};

const scrollToCommentBottom = () => {
  const commentBottom = document.getElementById('comment-bottom');
  if (commentBottom) {
    commentBottom.scrollIntoView({ behavior: 'smooth' });
  }
};
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
