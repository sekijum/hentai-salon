<template>
  <v-list class="no-padding-list">
    <CommentItem
      v-for="comment in comments"
      :key="comment.id"
      :comment="comment"
      :commentLimit="commentLimit"
      :threadId="threadId"
    />
  </v-list>
</template>

<script setup lang="ts">
import type { IThreadComment } from '~/types/thread';
import CommentItem from '~/components/comment/CommentItem.vue';

defineProps<{ comments: IThreadComment[]; commentLimit: number; threadId: number }>();
</script>

<style scoped>
.no-padding-list {
  padding: 0;
}
</style>
