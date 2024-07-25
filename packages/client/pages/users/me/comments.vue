<template>
  <div v-if="comment">
    <PageTitle title="マイレス" />

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
import type { IListResource } from '~/types/list-resource';
import CommentItem from '~/components/comment/CommentItem.vue';
import Pagination from '~/components/Pagination.vue';

definePageMeta({ middleware: ['logged-in-access-only'] });

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { $api } = nuxtApp;
const { getCommentLimit } = useStorage();

const menuItems = [
  { title: 'マイスレ', clicked: () => router.push('/users/me/threads'), icon: 'mdi-note' },
  { title: 'ユーザー情報', clicked: () => router.push('/users/me'), icon: 'mdi-account' },
  { title: 'マイレス', clicked: () => router.push('/users/me/comments'), icon: 'mdi-comment' },
];

const comment = ref<IListResource<IThreadComment>>();
const commentLimit = getCommentLimit();

onMounted(async () => {
  await fetchComments();
});

async function fetchComments() {
  const response = await $api.get<IListResource<IThreadComment>>('/users/me/comments', {
    params: { offset: route.query.offset, limit: commentLimit },
  });

  comment.value = response.data;
}
</script>
