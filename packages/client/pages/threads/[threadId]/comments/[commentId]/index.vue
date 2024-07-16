<template>
  <div v-if="threadComment">
    <PageTitle title="コメント詳細" />

    <v-divider />

    <CommentList :comments="[threadComment]" :commentLimit="commentLimit" :threadId="threadComment.thread.id" />

    <v-divider />
    <Pagination :totalCount="threadComment.replies.totalCount" :limit="commentLimit" />
    <v-divider />

    <div id="comment-top" />
    <CommentList
      v-if="threadComment.replies.data"
      :comments="threadComment?.replies.data"
      :commentLimit="commentLimit"
      :threadId="threadComment.thread.id"
    />
    <div id="comment-bottom" />

    <Pagination :totalCount="threadComment.replies.totalCount" :limit="commentLimit" />

    <v-btn icon large color="primary" class="fab fab-top" @click="scrollToCommentTop">
      <v-icon>mdi-arrow-up</v-icon>
    </v-btn>

    <v-btn icon large color="primary" class="fab fab-bottom" @click="scrollToCommentBottom">
      <v-icon>mdi-arrow-down</v-icon>
    </v-btn>

    <v-divider />

    <!-- <Pagination :totalCount="threadComment.replies.totalReplies" :limit="commentLimit" /> -->
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
import type { IThreadComment } from '~/types/thread-comment';

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { setThreadViewHistory, getThreadViewHistory } = useStorage();

const { $api } = nuxtApp;

const commentLimit = 100;
const isLoading = ref(true);
const threadsByHistory = ref<IThread[]>([]);
const threadsByPopular = ref<IThread[]>([]);
const threadComment = ref<IThreadComment>();

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
  await fetchThreads();
  await fetchComment();
});

async function fetchComment() {
  isLoading.value = true;
  const threadId = route.params.threadId;
  const commentId = route.params.commentId;
  const response = await $api.get<IThreadComment>(`/threads/${threadId}/comments/${commentId}`, {
    params: {
      limit: route.query.limit || commentLimit,
      offset: route.query.offset,
    },
  });
  threadComment.value = response.data;
  isLoading.value = false;
}

async function fetchThreads() {
  const queryCriteria = ['popularity'];
  if (getThreadViewHistory().length) {
    queryCriteria.push('history');
  }
  const response = await $api.get<{ threadsByHistory: IThread[]; threadsByPopular: IThread[] }>('/threads/', {
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
    fetchComment();
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
