<template>
  <div>
    <template v-for="(comment, idx) in sortedComments" :key="comment.id">
      <CommentItem :comment="comment" :commentLimit="commentLimit" :threadId="threadId" @replied="emit('replied')" />
      <template v-if="(idx + 1) % 5 === 0">
        <Ad :content="getRandomContent()" />
        <v-divider class="border-opacity-75"></v-divider>
      </template>
    </template>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { IThreadComment } from '~/types/thread-comment';
import CommentItem from '~/components/comment/CommentItem.vue';
import Ad from '~/components/Ad.vue';

const emit = defineEmits(['replied']);
const { getCommentOrder } = useStorage();

const props = defineProps<{
  comments: IThreadComment[];
  commentLimit: number;
  threadId: number;
  adContents: string[];
}>();

const sortedComments = computed(() => {
  return getCommentOrder() === 'asc' ? [...props.comments].reverse() : props.comments;
});

function getRandomContent() {
  const randomIndex = Math.floor(Math.random() * props.adContents.length);
  return props.adContents[randomIndex];
}
</script>

<style scoped></style>
