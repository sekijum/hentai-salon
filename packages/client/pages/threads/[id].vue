<template>
  <div v-if="thread">
    <PageTitle :title="thread.title" />

    <v-divider />

    <v-chip-group active-class="primary--text" column>
      <v-chip size="x-small" v-for="tag in thread.tags" :key="tag">
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

      <v-divider />

      <div id="comment-top" />
      <CommentList :comments="thread?.comments.data" :commentLimit="commentLimit" :threadId="thread.id" />
      <div id="comment-bottom" />

      <Pagination :totalCount="thread.comments.totalCount" :limit="commentLimit" />

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
import type { IThread } from '~/types/thread';

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { $api } = nuxtApp;

const commentLimit = 50;

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

const thread = ref<IThread>();

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
  const response = await $api.get<IThread>(`/threads/${threadId}`, {
    params: {
      limit: route.query.limit || commentLimit,
      offset: route.query.offset,
    },
  });
  console.log(response);
  thread.value = response.data;
}

watchEffect(() => {
  if (route.query.limit) {
    scrollToCommentTop();
    fetchThreads();
  }
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
