<template>
  <div v-if="thread">
    <PageTitle :title="thread.title" />

    <h2 class="font-weight-regular page-description">{{ thread.description }}</h2>

    <v-chip-group v-if="thread.tagNameList" active-class="primary--text" column>
      <v-chip size="x-small" v-for="tagName in thread.tagNameList" :key="tagName">
        {{ tagName }}
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
      <CommentForm @submit="fetchThread" :showReplyForm="true" :threadId="thread.id" />

      <v-divider />

      <div id="comment-top" />
      <CommentList
        :comments="thread?.comments.data"
        :commentLimit="commentLimit"
        :threadId="thread.id"
        @replied="fetchThread"
      />
      <div id="comment-bottom" />

      <CommentForm @submit="fetchThread" :showReplyForm="true" :threadId="thread.id" />

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
      v-if="threads.threadsByRelated?.length"
      queryCriteria="related"
      title="関連"
      :items="threads?.threadsByRelated"
      :isInfiniteScroll="true"
    />
  </div>
</template>

<script setup lang="ts">
import CommentList from '~/components/comment/CommentList.vue';
import CommentForm from '~/components/comment/CommentForm.vue';
import Menu from '~/components/Menu.vue';
import PageTitle from '~/components/PageTitle.vue';
import Pagination from '~/components/Pagination.vue';
import MediaGallery from '~/components/MediaGallery.vue';
import ThreadList from '~/components/thread/ThreadList.vue';
import type { IThread } from '~/types/thread';

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { setThreadViewHistory, getThreadViewHistory, getCommentLimit, getCommentSortOrder } = useStorage();

const { $api } = nuxtApp;

const snackbar = useState('isSnackbar', () => {
  return { isSnackbar: false, text: '' };
});

const commentLimit = getCommentLimit();
const thread = ref<IThread>();

const threads = ref<{
  threadsByRelated: IThread[];
}>({
  threadsByRelated: [],
});

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
  const threadId = route.params.threadId;
  const response = await $api.get<IThread>(`/threads/${threadId}`, {
    params: {
      limit: route.query.limit || commentLimit,
      offset: route.query.offset,
      sortOrder: getCommentSortOrder(),
    },
  });
  thread.value = response.data;
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