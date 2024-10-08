<template>
  <div>
    <PageTitle :title="thread.title" />

    <Content :text="thread.description" />

    <v-chip-group v-if="thread.tagNameList" active-class="primary--text" column>
      <v-chip size="x-small" v-for="(tagName, index) in thread.tagNameList" :key="`tag-${index}`">
        <nuxt-link
          :to="`/threads?filter=tags&tagNameList=${tagName}`"
          class="text-decoration-none bg-transparent"
          target="_blank"
          rel="noopener"
        >
          {{ tagName }}
        </nuxt-link>
      </v-chip>
    </v-chip-group>

    <v-divider />

    <Menu :items="menuItems" />

    <Pagination :totalCount="thread.comments.totalCount" :limit="commentLimit" />

    <template v-if="route.query.tab === 'media'">
      <div id="media-top" />
      <MediaGallery
        v-if="thread.comments.data?.length"
        :attachments="thread.attachments.data"
        :commentLimit="commentLimit"
        :threadId="thread.id"
        :adContents="ads"
      />
      <div id="media-bottom" />

      <v-bottom-navigation color="teal" grow :fixed="true">
        <v-btn @click="scrollToMediaBottom">
          <v-icon>mdi-arrow-down</v-icon>
        </v-btn>

        <template v-if="payload?.user?.id === thread.userId">
          <v-btn @click="editThread">
            <v-icon>mdi-pencil</v-icon>
          </v-btn>
        </template>

        <v-btn @click="toggleLike">
          <v-icon>{{ thread.isLiked ? 'mdi-thumb-up' : 'mdi-thumb-up-outline' }}</v-icon>
        </v-btn>

        <v-btn @click="scrollToMediaTop">
          <v-icon>mdi-arrow-up</v-icon>
        </v-btn>
      </v-bottom-navigation>
    </template>
    <template v-else>
      <CommentForm @submit="fetchThread" :showReplyForm="true" :threadId="thread.id" />

      <v-divider />

      <div id="comment-top" />
      <CommentList
        v-if="thread.comments.data?.length"
        :comments="thread?.comments.data"
        :commentLimit="commentLimit"
        :threadId="thread.id"
        @replied="fetchThread"
        :adContents="ads"
      />
      <div id="comment-bottom" />

      <CommentForm @submit="fetchThread" :showReplyForm="true" :threadId="thread.id" />

      <v-bottom-navigation color="teal" grow :fixed="true">
        <v-btn @click="scrollToCommentBottom">
          <v-icon>mdi-arrow-down</v-icon>
        </v-btn>

        <template v-if="payload?.user?.id === thread.userId">
          <v-btn @click="editThread">
            <v-icon>mdi-pencil</v-icon>
          </v-btn>
        </template>

        <v-btn @click="toggleLike">
          <v-icon>{{ thread.isLiked ? 'mdi-thumb-up' : 'mdi-thumb-up-outline' }}</v-icon>
        </v-btn>

        <v-btn @click="scrollToCommentTop">
          <v-icon>mdi-arrow-up</v-icon>
        </v-btn>
      </v-bottom-navigation>
    </template>

    <v-divider />

    <Pagination :totalCount="thread.comments.totalCount" :limit="commentLimit" />

    <SwiperGrid v-if="commentAttachments?.length" :commentAttachments="commentAttachments" />

    <ThreadItem filter="related-by-thread" title="関連スレ" :isInfiniteScroll="true" />
  </div>
</template>

<script setup lang="ts">
import CommentList from '~/components/comment/CommentList.vue';
import CommentForm from '~/components/comment/CommentForm.vue';
import Menu from '~/components/Menu.vue';
import ThreadItem from '~/components/thread/ThreadItem.vue';
import PageTitle from '~/components/PageTitle.vue';
import Pagination from '~/components/Pagination.vue';
import MediaGallery from '~/components/MediaGallery.vue';
import type { IThread } from '~/types/thread';
import type { IThreadCommentAttachmentForThread } from '~/types/thread-comment-attachment';
import Content from '~/components/Content.vue';

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { setThreadViewHistory, getThreadViewHistory, getCommentLimit } = useStorage();
const commentAttachments = ref<IThreadCommentAttachmentForThread[]>([]);
const ads = ref<string[]>([]);

const { $api, payload } = nuxtApp;

const commentLimit = getCommentLimit();

const thread = ref<IThread>({
  id: 0,
  title: '',
  description: '',
  thumbnailUrl: '',
  tagNameList: [],
  commentCount: 0,
  userId: 0,
  comments: { totalCount: 0, limit: 0, offset: 0, data: [] },
  attachments: { totalCount: 0, limit: 0, offset: 0, data: [] },
  isLiked: false,
});

const menuItems = [
  {
    title: 'コメント一覧',
    clicked: () => {
      const newQuery = { ...route.query };
      delete newQuery.tab;
      router.push({ query: newQuery });
    },
    icon: 'mdi-comment-text',
  },
  {
    title: 'メディア',
    clicked: () => {
      const newQuery = { ...route.query, tab: 'media' };
      router.push({ query: newQuery });
    },
    icon: 'mdi-folder-multiple-image',
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

async function fetchRelatedByHistoryComments() {
  const response = await $api.get<IThreadCommentAttachmentForThread[]>('/attachments', {
    params: {
      filter: 'related-by-history',
      threadIds: getThreadViewHistory(),
    },
  });
  commentAttachments.value = response.data;
}

onMounted(async () => {
  await fetchAds();
  await fetchThread();
  setThreadViewHistory(parseInt(route.params.threadId.toString(), 10));
  useHead({
    title: thread.value.title,
    meta: [
      { property: 'og:title', content: thread.value.title },
      { name: 'description', content: thread.value.description },
      { property: 'og:image', content: thread.value.thumbnailUrl || '/hentai-salon-logo/logo_transparent.png' },
      { property: 'og:url', content: location.href },
    ],
  });
  await fetchRelatedByHistoryComments();
});

async function fetchThread() {
  const threadId = route.params.threadId;
  const response = await $api.get<IThread>(`/threads/${threadId}`, {
    params: {
      limit: commentLimit,
      offset: route.query.offset,
    },
  });
  thread.value = response.data;
}

async function toggleLike() {
  if (!payload.isLoggedIn) {
    alert('ログインしてください。');
    return;
  }
  if (thread.value.isLiked) {
    await $api.post(`/threads/${thread.value.id}/unlike`);
    thread.value.isLiked = false;
  } else {
    await $api.post(`/threads/${thread.value.id}/like`);
    thread.value.isLiked = true;
  }
}

async function fetchAds() {
  const response = await $api.get<string[]>('/ads');
  ads.value = response.data;
}

const editThread = () => {
  router.push(`/threads/${thread.value.id}/edit`);
};

watch(
  () => route.query.offset,
  () => fetchThread(),
);
</script>

<style scoped>
.page-description {
  font-size: 12px;
}
</style>
