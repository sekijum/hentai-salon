<template>
  <div v-if="comment">
    <PageTitle title="コメント詳細" />

    <v-btn block @click="() => router.push(`/threads/${comment?.thread.id}`)">
      <v-icon class="menu-icon">mdi-arrow-left</v-icon>{{ comment.thread.title }}
    </v-btn>

    <v-divider />

    <CommentList :comments="[comment]" :commentLimit="commentLimit" :threadId="comment.thread.id" @replied="fetchComment" />

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

    <ThreadItem filter="related-by-history" title="関連" :isInfiniteScroll="true" />
  </div>
</template>

<script setup lang="ts">
import CommentList from '~/components/comment/CommentList.vue';
import PageTitle from '~/components/PageTitle.vue';
import Pagination from '~/components/Pagination.vue';
import ThreadItem from '~/components/thread/ThreadItem.vue';
import type { IThreadComment } from '~/types/thread-comment';

const route = useRoute();
const router = useRouter();
const nuxtApp = useNuxtApp();
const { getCommentLimit } = useStorage();

const { $api } = nuxtApp;

const commentLimit = getCommentLimit();
const comment = ref<IThreadComment>();

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
  await fetchComment();

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
        property: 'og:url',
        content: location.href,
      },
    ],
  });
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
