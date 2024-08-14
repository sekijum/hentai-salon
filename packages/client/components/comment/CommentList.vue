<template>
  <CommentItem
    v-for="comment in comments"
    :key="comment.id"
    :comment="comment"
    :commentLimit="commentLimit"
    :threadId="threadId"
    @replied="emit('replied')"
  />
</template>

<script setup lang="ts">
import type { IThreadComment } from '~/types/thread-comment';
import CommentItem from '~/components/comment/CommentItem.vue';

const emit = defineEmits(['replied']);

const props = defineProps<{
  comments: IThreadComment[];
  commentLimit: number;
  threadId: number;
}>();
</script>

<style scoped></style>
