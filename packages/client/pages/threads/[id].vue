<template>
  <div v-if="thread">
    <PageTitle :title="thread.title" />

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
      <CommentList :comments="thread?.comments.slice(0, 50)" />
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

const thread = ref<TThread>();

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
  thread.value = response.data;
  console.log(thread.value);
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
