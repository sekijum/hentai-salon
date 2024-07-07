<template>
  <v-list class="no-padding-list">
    <CommentItem v-for="(comment, idx) in comments" :key="comment.id" :comment="comment" :idx="idx" />
  </v-list>
</template>

<script setup lang="ts">
import type { TThreadComment } from '~/types/thread';
import CommentItem from '~/components/comment/CommentItem.vue';

defineProps<{ comments: TThreadComment[] }>();
</script>

<style scoped>
.no-padding-list {
  padding: 0;
}
</style>
