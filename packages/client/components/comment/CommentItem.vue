<template>
  <div class="comment-item" :id="'comments-' + idx">
    <v-list-item class="comment-list-item">
      <div class="comment-header">
        <div class="comment-header-text">
          {{ idx }}
          <router-link to="/" class="username-link">{{
            comment?.guestName || '名無し' || 'ログインユーザー名表示'
          }}</router-link>
        </div>
      </div>
      <div v-if="comment.parentCommentID" class="reply-indication">
        <router-link :to="'#comments-' + comment.parentCommentID" class="reply-link">
          >>{{ comment.parentCommentID }}
        </router-link>
      </div>
      <v-list-item-title class="comment-content">
        <a :href="'#comment-' + comment.id" class="comment-anchor">{{ comment.content }}</a>
      </v-list-item-title>
      <template v-if="comment.attachments && comment.attachments.length">
        <v-row dense>
          <v-col cols="3" v-for="(attachment, index) in comment.attachments" :key="index" class="media-col">
            <div class="media-item-wrapper" @click="() => openDialog(attachment)">
              <v-img :src="attachment.type === 'Video' ? attachment.url : attachment.url" class="media-item">
                <template v-slot:placeholder>
                  <v-row align="center" class="fill-height ma-0" justify="center">
                    <v-progress-circular color="grey-lighten-5" indeterminate></v-progress-circular>
                  </v-row>
                </template>
              </v-img>
              <v-icon v-if="attachment.type === 'Video'" size="40" class="play-icon">mdi-play-circle</v-icon>
            </div>
          </v-col>
        </v-row>
      </template>
      <div class="interaction-section">
        <v-row dense>
          <v-col cols="6">
            <v-icon small @click="toggleReplyForm">{{ showReplyForm ? 'mdi-close' : 'mdi-reply' }}</v-icon>
          </v-col>
          <v-col cols="6" class="interaction-right">
            <router-link :to="'/comments/' + comment.id + '/replies'" class="interaction-link">
              <v-icon small>mdi-comment</v-icon>
              <span class="interaction-text">{{ comment.totalReplies }}</span>
            </router-link>
            <router-link :to="'/comments/' + comment.id" class="interaction-link id-link">
              <span class="interaction-text">ID: {{ comment.id }}</span>
            </router-link>
          </v-col>
        </v-row>
      </div>
      <div v-if="showReplyForm" class="reply-form">
        <CommentForm :formTitle="'返信 >> ' + comment.id" @submit="submitReply" @clear="clearReplyForm" />
      </div>
    </v-list-item>
    <v-divider></v-divider>

    <ModalMedia :dialog="dialog" :mediaItem="selectedMedia" @close="closeDialog" @update:dialog="dialog = $event" />
  </div>
</template>

<script setup lang="ts">
import ModalMedia from '~/components/ModalMedia.vue';
import CommentForm from '~/components/comment/CommentForm.vue';
import type { TThreadComment } from '~/types/thread';

defineProps<{ idx: number; comment: TThreadComment }>();

const dialog = ref(false);
const selectedMedia = ref(null);
const showReplyForm = ref(false);

const toggleReplyForm = () => {
  showReplyForm.value = !showReplyForm.value;
  if (!showReplyForm.value) {
    clearReplyForm();
  }
};

const submitReply = () => {
  console.log('返信を送信');
  alert('返信しました。');
  toggleReplyForm();
};

const clearReplyForm = () => {
  console.log('返信フォームをクリア');
};

const openDialog = attachment => {
  selectedMedia.value = attachment;
  dialog.value = true;
};

const closeDialog = () => {
  dialog.value = false;
};
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
  font-size: 14px;
}

.username-link {
  color: blue;
  text-decoration: underline;
}

.reply-indication {
  background-color: #f0f0f0;
  border: 1px solid #ddd;
  border-radius: 4px;
  display: inline-block;
  padding: 2px 6px;
}

.reply-link {
  text-decoration: none;
  color: #333;
}

.comment-anchor {
  color: inherit;
  text-decoration: none;
}

.comment-anchor:hover {
  text-decoration: underline;
}

.interaction-section {
  display: flex;
  align-items: center;
}

.interaction-link {
  display: flex;
  align-items: center;
  text-decoration: none;
  color: grey;
}

.interaction-text {
  color: grey;
}

.interaction-right {
  display: flex;
  justify-content: flex-end;
}

.id-link {
  margin-left: 20px;
}

.comment-content {
  font-size: 14px;
  white-space: normal;
  word-wrap: break-word;
}

.media-item-wrapper {
  position: relative;
}

.media-item {
  width: 100px;
  height: 100px;
  object-fit: cover;
  cursor: pointer;
}

.play-icon {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  color: white;
  font-size: 40px;
}

.reply-form {
  margin-top: 16px;
}
</style>
