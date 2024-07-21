<template>
  <div>
    <PageTitle title="マイレス" />

    <Menu :items="menuItems" />

    <v-infinite-scroll height="100%" :items="comments" :onLoad="loadComment">
      <template v-for="comment in comments?.data" :key="comment.id">
        <CommentItem
          :comment="comment"
          :commentLimit="commentLimit"
          :threadId="comment.thread.id"
          @replied="fetchComments"
        />
      </template>
      <template v-slot:empty>これ以上ありません</template>
    </v-infinite-scroll>
  </div>
</template>

<script setup lang="ts">
import PageTitle from '~/components/PageTitle.vue';
import type { IThreadComment } from '~/types/thread-comment';
import type { IListResource } from '~/types/list-resource';
import CommentItem from '~/components/comment/CommentItem.vue';

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { $api } = nuxtApp;
const { getCommentLimit } = useStorage();

const menuItems = [
  { title: 'ユーザー情報', clicked: () => router.push('/users/me'), icon: 'mdi-update' },
  { title: 'マイスレ', clicked: () => router.push('/users/me/threads'), icon: 'mdi-new-box' },
  { title: 'マイレス', clicked: () => router.push('/users/me/comments'), icon: 'mdi-format-list-bulleted' },
];

const comments = ref<IListResource<IThreadComment>>();
const commentLimit = getCommentLimit();
const offset = ref(0);

onMounted(async () => {
  await fetchComments();
});

async function loadComment({ done }: { done: (status: 'loading' | 'error' | 'empty' | 'ok') => void }) {
  offset.value += commentLimit;
  const { canNextLoad } = await fetchComments(offset.value);
  canNextLoad ? done('ok') : done('empty');
}

async function fetchComments(offset: number = 0) {
  const response = await $api.get<IListResource<IThreadComment>>('/users/me/comments', {
    params: { offset, limit: commentLimit },
  });

  console.log(response);
  if (!response.data.data || response.data.data.length > commentLimit) {
    return { canNextLoad: false };
  }

  if (offset) {
    response.data.data.map(item => comments.value?.data.push(item));
  } else {
    comments.value = response.data;
  }

  return { canNextLoad: true };
}
</script>
