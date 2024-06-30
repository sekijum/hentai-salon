<template>
  <div class="comment-item">
    <v-list-item class="comment-list-item">
      <v-list-item-content class="comment-list-item-content">
        <div class="d-flex justify-space-between comment-header">
          <div class="comment-header-text">
            {{ idx }} <a :href="'/user/' + comment.username" class="username-link">{{ comment.username }}</a
            >[Lv.{{ comment.level }}]{{ comment.rank }} {{ comment.date }} {{ comment.time }}
          </div>
          <div class="reply-section">
            <v-icon small>mdi-reply</v-icon>
            <span class="reply-text">返信</span>
          </div>
        </div>
        <v-list-item-title class="comment-content">{{ comment.content }}</v-list-item-title>
      </v-list-item-content>
    </v-list-item>
    <v-divider></v-divider>
  </div>
</template>

<script setup>
defineProps({
  idx: Number,
  comment: Object,
});
</script>

<style scoped>
.comment-item {
  font-size: 14px;
}

.comment-list-item {
  padding-left: 5px !important;
  padding-right: 5px !important;
}

.comment-header-text {
  flex: 1;
  word-wrap: break-word;
}

.reply-section {
  display: flex;
  align-items: center;
}

.username-link {
  color: inherit;
  text-decoration: underline;
}

.comment-content {
  font-size: 16px;
  white-space: normal;
  word-wrap: break-word;
}
</style>
