<template>
  <div v-if="thread">
    <PageTitle :title="thread.title" />

    <h2 class="font-weight-regular page-description">{{ thread.description }}</h2>

    <v-chip-group v-if="thread.tags" active-class="primary--text" column>
      <v-chip size="x-small" v-for="tag in thread.tags" :key="tag">
        {{ tag }}
      </v-chip>
    </v-chip-group>

    <v-divider />

    <Menu :items="menuItems" />

    <Pagination :totalCount="thread.comments.totalCount" :limit="commentLimit" />

    <template v-if="route.query.tab === 'media'">
      <div id="media-top" />
      <MediaGallery
        v-if="thread.attachments"
        :attachments="thread.attachments"
        :commentLimit="commentLimit"
        :threadId="thread.id"
      />
      <div id="media-bottom" />

      <v-btn icon large color="primary" class="fab fab-top" @click="scrollToMediaTop">
        <v-icon>mdi-arrow-up</v-icon>
      </v-btn>
      <v-btn icon large color="primary" class="fab fab-bottom" @click="scrollToMediaBottom">
        <v-icon>mdi-arrow-down</v-icon>
      </v-btn>
    </template>
    <template v-else>
      <CommentForm />

      <v-divider />

      <div id="comment-top" />
      <CommentList :comments="thread?.comments.data" :commentLimit="commentLimit" :threadId="thread.id" />
      <div id="comment-bottom" />

      <CommentForm />

      <v-btn icon large color="primary" class="fab fab-top" @click="scrollToCommentTop">
        <v-icon>mdi-arrow-up</v-icon>
      </v-btn>

      <v-btn icon large color="primary" class="fab fab-bottom" @click="scrollToCommentBottom">
        <v-icon>mdi-arrow-down</v-icon>
      </v-btn>
    </template>

    <v-divider />

    <Pagination :totalCount="thread.comments.totalCount" :limit="commentLimit" />
    <ThreadList
      title="人気"
      :items="threadsByPopular"
      :clicked="() => router.push({ path: '/threads' })"
      :isInfiniteScroll="false"
    />
    <ThreadList
      title="閲覧履歴"
      :items="threadsByHistory"
      :clicked="() => router.push({ path: '/threads', query: { queryCriteria: ['history'] } })"
    />
  </div>

  <OverlayLoagind :isLoading="isLoading" title="読込中" />
</template>

<script setup lang="ts">
import CommentList from '~/components/comment/CommentList.vue';
import CommentForm from '~/components/comment/CommentForm.vue';
import Menu from '~/components/Menu.vue';
import PageTitle from '~/components/PageTitle.vue';
import Pagination from '~/components/Pagination.vue';
import OverlayLoagind from '~/components/OverlayLoagind.vue';
import MediaGallery from '~/components/MediaGallery.vue';
import ThreadList from '~/components/thread/ThreadList.vue';
import type { IThread } from '~/types/thread';

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { setThreadViewHistory, getThreadViewHistory } = useStorage();

const { $api } = nuxtApp;

const commentLimit = 100;
const isLoading = ref(true);
const thread = ref<IThread>();
const threadsByHistory = ref<IThread[]>([]);
const threadsByPopular = ref<IThread[]>([]);

const menuItems = [
  {
    title: 'コメント一覧',
    clicked: () => router.replace({ query: {} }),
    icon: 'mdi-fire',
  },
  {
    title: 'メディア',
    clicked: () => router.replace({ query: { tab: 'media' } }),
    icon: 'mdi-update',
  },
];

function scrollToMediaTop() {
  const mediaTop = document.getElementById('media-top');
  if (mediaTop) {
    mediaTop.scrollIntoView({ behavior: 'smooth' });
  }
}

function scrollToMediaBottom() {
  const mediaBottom = document.getElementById('media-bottom');
  if (mediaBottom) {
    mediaBottom.scrollIntoView({ behavior: 'smooth' });
  }
}

function scrollToCommentTop() {
  const commentTop = document.getElementById('comment-top');
  if (commentTop) {
    commentTop.scrollIntoView({ behavior: 'smooth' });
  }
}

function scrollToCommentBottom() {
  const commentBottom = document.getElementById('comment-bottom');
  if (commentBottom) {
    commentBottom.scrollIntoView({ behavior: 'smooth' });
  }
}

onMounted(async () => {
  await fetchThread();
  await fetchThreads();
  setThreadViewHistory(parseInt(route.params.threadId.toString(), 10));
});

async function fetchThread() {
  isLoading.value = true;
  const threadId = route.params.threadId;
  const response = await $api.get<IThread>(`/threads/${threadId}`, {
    params: {
      limit: route.query.limit || commentLimit,
      offset: route.query.offset,
    },
  });
  thread.value = response.data;
  isLoading.value = false;
}

async function fetchThreads() {
  const queryCriteria = ['popularity'];
  if (getThreadViewHistory().length) {
    queryCriteria.push('history');
  }
  const response = await $api.get<{ threadsByHistory: IThread[]; threadsByPopular: IThread[] }>('/threads', {
    params: {
      queryCriteria: queryCriteria,
      threadIds: getThreadViewHistory(),
      limit: 10,
    },
  });
  threadsByHistory.value = response.data.threadsByHistory;
  threadsByPopular.value = response.data.threadsByPopular;
}

watchEffect(() => {
  if (route.query.limit) {
    fetchThread();
  }
});
</script>

<style scoped>
.page-description {
  font-size: 1rem;
}

.fab {
  position: fixed;
  right: 16px;
}

.fab-top {
  bottom: 72px;
}

.fab-bottom {
  bottom: 16px;
}
</style>
