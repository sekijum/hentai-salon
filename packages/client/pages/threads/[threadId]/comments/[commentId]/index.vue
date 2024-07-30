<template>
  <div v-if="comment">
    <PageTitle title="コメント詳細" />

    <v-divider />

    <CommentList
      :comments="[comment]"
      :commentLimit="commentLimit"
      :threadId="comment.thread.id"
      @replied="fetchComment"
    />

    <v-divider />
    <Pagination :totalCount="comment.replies.totalCount" :limit="commentLimit" />
    <v-divider />

    <div id="comment-top" />
    <CommentList
      v-if="comment.replies.data"
      :comments="comment?.replies.data"
      :commentLimit="commentLimit"
      :threadId="comment.thread.id"
      @replied="fetchComment"
    />
    <div id="comment-bottom" />

    <Pagination :totalCount="comment.replies.totalCount" :limit="commentLimit" />

    <v-btn icon large color="primary" class="fab fab-top" @click="scrollToCommentTop">
      <v-icon>mdi-arrow-up</v-icon>
    </v-btn>

    <v-btn icon large color="primary" class="fab fab-bottom" @click="scrollToCommentBottom">
      <v-icon>mdi-arrow-down</v-icon>
    </v-btn>

    <v-divider />

    <ThreadList
      v-if="threads.threadsByRelated.length"
      filter="related"
      title="関連"
      :items="threads?.threadsByRelated"
      :isInfiniteScroll="true"
      :threadLimit="threadLimit"
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

const route = useRoute();
const nuxtApp = useNuxtApp();
const { getThreadViewHistory, getCommentLimit } = useStorage();

const { $api } = nuxtApp;

const commentLimit = getCommentLimit();
const threadLimit = 10;
const comment = ref<IThreadComment>();

const threads = ref<{ threadsByRelated: IThread[] }>({ threadsByRelated: [] });

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
  comment.value = response.data;
}

async function fetchThreads() {
  await Promise.all(
    ['related'].map(async filter => {
      const response = await $api.get<IThread[]>('/threads', {
        params: {
          filter,
          threadIds: getThreadViewHistory(),
          limit: threadLimit,
        },
      });
      if (filter === 'related') {
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

useHead({
  title: comment.value?.content,
  meta: [
    { name: 'description', content: comment.value?.content },
    {
      property: 'og:title',
      content: comment.value?.content,
    },
    {
      property: 'og:description',
      content: comment.value?.content,
    },
    {
      property: 'og:image',
      content: comment.value?.attachments[0].url || '',
    },
    {
      property: 'og:url',
      content: location.href,
    },
  ],
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
