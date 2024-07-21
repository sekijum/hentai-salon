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
      @replied="fetchComment"
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

    <ThreadList
      v-if="threads.threadsByRelated.length"
      queryCriteria="related"
      title="関連"
      :items="threads?.threadsByRelated"
      :isInfiniteScroll="true"
    />
  </div>
</template>

<script setup lang="ts">
import CommentList from '~/components/comment/CommentList.vue';
import PageTitle from '~/components/PageTitle.vue';
import Pagination from '~/components/Pagination.vue';
import ThreadList from '~/components/thread/ThreadList.vue';
import type { IThread } from '~/types/thread';
import type { IThreadComment } from '~/types/thread-comment';

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { getThreadViewHistory, getCommentLimit } = useStorage();

const { $api } = nuxtApp;

const commentLimit = getCommentLimit();
const threadComment = ref<IThreadComment>();

const threads = ref<{
  threadsByRelated: IThread[];
}>({
  threadsByRelated: [],
});

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
  const threadId = route.params.threadId;
  const commentId = route.params.commentId;
  const response = await $api.get<IThreadComment>(`/threads/${threadId}/comments/${commentId}`, {
    params: {
      limit: route.query.limit || commentLimit,
      offset: route.query.offset,
    },
  });
  threadComment.value = response.data;
}

async function fetchThreads() {
  await Promise.all(
    ['related'].map(async queryCriteria => {
      const response = await $api.get<IThread[]>('/threads/', {
        params: {
          queryCriteria,
          threadIds: getThreadViewHistory(),
          limit: 10,
        },
      });
      if (queryCriteria === 'related') {
        threads.value.threadsByRelated = response.data;
      }
    }),
  );
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
