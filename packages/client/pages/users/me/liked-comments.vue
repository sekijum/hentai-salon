<template>
  <div v-if="comment">
    <PageTitle title="お気に入りレス" />

    <Menu :items="menuItems" />

    <Pagination :totalCount="comment?.totalCount" :limit="commentLimit" />

    <template v-for="comment in comment?.data" :key="comment.id">
      <CommentItem
        :comment="comment"
        :commentLimit="commentLimit"
        :threadId="comment.thread.id"
        @replied="fetchComments"
      />
    </template>

    <Pagination :totalCount="comment.totalCount" :limit="commentLimit" />
  </div>
</template>

<script setup lang="ts">
import PageTitle from '~/components/PageTitle.vue';
import type { IThreadComment } from '~/types/thread-comment';
import type { ICollection } from '~/types/collection';
import CommentItem from '~/components/comment/CommentItem.vue';
import Pagination from '~/components/Pagination.vue';

definePageMeta({ middleware: ['logged-in-access-only'] });

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { $api } = nuxtApp;
const { getCommentLimit } = useStorage();

const menuItems = [
  { title: 'マイスレ', clicked: () => router.push('/users/me/threads'), icon: 'mdi-file-document-multiple-outline' },
  { title: 'ユーザー情報', clicked: () => router.push('/users/me'), icon: 'mdi-account-cog-outline' },
  { title: 'マイレス', clicked: () => router.push('/users/me/comments'), icon: 'mdi-message-text-outline' },
  {
    title: 'お気に入りスレ',
    clicked: () => router.push('/users/me/liked-threads'),
    icon: 'mdi-star-box-multiple-outline',
  },
  { title: 'お気に入りレス', clicked: () => router.push('/users/me/liked-comments'), icon: 'mdi-message-star-outline' },
];

const comment = ref<ICollection<IThreadComment>>();
const commentLimit = getCommentLimit();

onMounted(async () => {
  await fetchComments();
});

async function fetchComments() {
  const response = await $api.get<ICollection<IThreadComment>>('/users/me/liked-comments', {
    params: { offset: route.query.offset, limit: commentLimit },
  });

  comment.value = response.data;
}

watch(
  () => route.query.offset,
  () => fetchComments(),
);
</script>
