<template>
  <v-list class="no-padding-list">
    <CommentItem
      v-for="(comment, idx) in comments"
      :key="comment.id"
      :comment="comment"
      :commentLimit="commentLimit"
      :threadId="threadId"
    />
  </v-list>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router';
import type { IThreadComment } from '~/types/thread';
import CommentItem from '~/components/comment/CommentItem.vue';

defineProps<{ comments: IThreadComment[]; commentLimit: number; threadId: number }>();

const route = useRoute();

function calculateIndex(idx: number): number {
  const limit = Number(route.query.limit) || 0;
  const offset = Number(route.query.offset) || 0;
  return offset + idx + 1;
}
</script>

<style scoped>
.no-padding-list {
  padding: 0;
}
</style>
