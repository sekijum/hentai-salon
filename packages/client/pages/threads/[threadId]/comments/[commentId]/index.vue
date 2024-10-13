<template>
  <div v-if="comment">
    <PageTitle title="コメント詳細" />

    <v-divider />

    <CommentList
      :comments="[comment]"
      :commentLimit="commentLimit"
      :threadId="comment.thread.id"
      @replied="fetchComment"
      :adContents="[]"
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
      :adContents="ads"
    />
    <div id="comment-bottom" />

    <Pagination :totalCount="comment.replies.totalCount" :limit="commentLimit" />

    <v-bottom-navigation color="teal" grow :fixed="true">
      <v-btn @click="() => router.push(`/threads/${comment?.thread.id}`)">
        <v-icon>mdi-arrow-left</v-icon>

        <span>{{ comment.thread.title }}</span>
      </v-btn>

      <v-btn @click="scrollToCommentBottom">
        <v-icon>mdi-arrow-down</v-icon>
      </v-btn>

      <v-btn @click="scrollToCommentTop">
        <v-icon>mdi-arrow-up</v-icon>
      </v-btn>
    </v-bottom-navigation>

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

const { $api, payload } = nuxtApp;

const commentLimit = getCommentLimit();
const comment = ref<IThreadComment>();
const ads = ref<string[]>([]);

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
  await fetchAds();

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

async function fetchAds() {
  const response = await $api.get<string[]>('/ads');
  ads.value = response.data;
}

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
