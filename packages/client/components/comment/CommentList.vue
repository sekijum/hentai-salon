<template>
  <CommentItem
    v-for="comment in sortedComments"
    :key="comment.id"
    :comment="comment"
    :commentLimit="commentLimit"
    :threadId="threadId"
    @replied="emit('replied')"
  />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { IThreadComment } from '~/types/thread-comment';
import CommentItem from '~/components/comment/CommentItem.vue';

const emit = defineEmits(['replied']);
const { getCommentOrder } = useStorage();

const props = defineProps<{
  comments: IThreadComment[];
  commentLimit: number;
  threadId: number;
}>();

// getCommentOrderに基づいてcommentsを逆順に並び替える
const sortedComments = computed(() => {
  return getCommentOrder() === 'asc' ? [...props.comments].reverse() : props.comments;
});
</script>

<style scoped></style>
