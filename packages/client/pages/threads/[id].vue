<template>
  <div>
    <PageTitle :title="item.title" />

    <v-divider></v-divider>

    <v-chip-group active-class="primary--text" column>
      <v-chip
        size="x-small"
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
import CommentList from '~/components/comment/CommentList.vue';
import CommentForm from '~/components/comment/CommentForm.vue';
import Menu from '~/components/Menu.vue';
import PageTitle from '~/components/PageTitle.vue';
import Pagination from '~/components/Pagination.vue';
import MediaGallery from '~/components/MediaGallery.vue';
import { useRouter, useRoute } from 'vue-router';
import type { TThread } from '~/types/thread';

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { $api } = nuxtApp;

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

const threads = ref<TThread>();

const item = ref({
  title: 'ラーメン店主異例の訴え',
  subtitle: '食事中にイヤホンやめて',
  link: '/news/1',
  comments: 12,
  board: 'ニュース',
});

const comments = ref([]);

const usernames = ['警備員', '名無しさん@恐縮です', '住民A', '住民B', '観察者'];
const contents = [
  '俺も恋人じゃないけど触らせてもらおう',
  'もう結婚するしか逃げ道はない',
  'もう飽きたこの話題',
  '歌舞伎俳優も未成年の舞妓と疑惑があっても伝統芸能の嗜みでOKのナゾのルールがあったみたいなものかw',
  'この話はもうやめよう',
];
const boards = ['垢版', 'ニュース', 'スポーツ', 'エンタメ', '一般'];
const types = ['大砲', '小銃', '拳銃'];
const idCodes = ['+Q6HhT800', 'X+z5zrq00', '6C/Ty2sa0', 'mswGlHt', '3H4JrKd9'];
const mediaUrls = [
  'https://placehold.jp/300x300.png',
  'https://drive.google.com/thumbnail?id=1atuQlT_wuPT73fwo4x1fZsGc2_ErVNqs&sz=w670',
  'https://www.w3schools.com/html/mov_bbb.mp4',
];

for (let i = 1; i <= 100; i++) {
  comments.value.push({
    id: i,
    username: usernames[i % usernames.length],
    level: Math.floor(Math.random() * 20) + 1,
    rank: ['新芽', '成長', '熟成', '老舗'][i % 4],
    date: `2024/06/${String((i % 30) + 1).padStart(2, '0')}`,
    time: `${String(Math.floor(Math.random() * 24)).padStart(2, '0')}:${String(Math.floor(Math.random() * 60)).padStart(
      2,
      '0',
    )}`,
    content: contents[i % contents.length],
    board: boards[i % boards.length],
    type: types[i % types.length],
    idCode: idCodes[i % idCodes.length],
    commentCount: String(Math.floor(Math.random() * 20)),
    replyTo: 20,
    media:
      i % 3 === 0
        ? [
            { type: 'image', url: mediaUrls[0] },
            { type: 'video/mp4', url: mediaUrls[2], thumbnail: mediaUrls[0] },
          ]
        : i % 3 === 1
        ? [{ type: 'image', url: mediaUrls[1] }]
        : [],
  });
}

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

onMounted(async () => {
  await fetchThreads();
});

async function fetchThreads() {
  const threadId = route.params.id;
  const response = await $api.get<TThread>(`/threads/${threadId}`);
  threads.value = response.data;
  console.log(threads.value);
}
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
