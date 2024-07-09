<template>
  <div class="comment-item" :id="'comments-' + comment.idx">
    <v-list-item class="comment-list-item">
      <div class="comment-header">
        <div class="comment-header-text">
          {{ comment.idx }}
          <nuxt-link to="/" class="username-link">{{
            comment?.guestName || '名無し' || 'ログインユーザー名表示'
          }}</nuxt-link>
        </div>
      </div>
      <div v-if="comment.parentCommentIdx" class="reply-indication">
        <nuxt-link
          class="reply-link"
          :to="{
            path: toParentComment(comment.parentCommentIdx).path,
            query: toParentComment(comment.parentCommentIdx).query,
            hash: toParentComment(comment.parentCommentIdx).hash,
          }"
          target="_blank"
          rel="noopener noreferrer"
        >
          >> {{ comment.parentCommentIdx }}
        </nuxt-link>
      </div>
      <v-list-item-title class="comment-content">
        <a :href="'#comment-' + comment.id" class="comment-anchor">{{ comment.content }}</a>
      </v-list-item-title>
      <template v-if="comment.attachments && comment.attachments.length">
        <v-row dense class="attachment-row">
          <v-chip
            v-for="(attachment, index) in comment.attachments"
            :key="index"
            class="attachment-chip d-flex justify-center align-center my-1 mx-1"
            @click="openModalMedia(attachment)"
            variant="outlined"
            size="x-small"
          >
            <v-icon left>
              <template v-if="attachment.type === 'Video'">mdi-video</template>
              <template v-else>mdi-image</template>
            </v-icon>
            {{ attachment.url }}
          </v-chip>
        </v-row>
      </template>
      <!-- <template v-if="comment.attachments && comment.attachments.length">
        <v-row dense>
          <v-col
            cols="3"
            v-for="(attachment, index) in comment.attachments"
            :key="index"
            class="media-col"
            @click="openModalMedia(attachment)"
          >
            <div class="media-item-wrapper">
              <template v-if="attachment.type === 'Video'">
                <video :src="attachment.url" class="media-item" muted @loadeddata="onVideoLoad">
                  <source :src="attachment.url" type="video/mp4" />
                </video>
                <v-icon size="40" class="play-icon">mdi-play-circle</v-icon>
              </template>
              <template v-else>
                <v-img :src="attachment.url" class="media-item" contain>
                  <template v-slot:placeholder>
                    <v-row align="center" class="fill-height ma-0" justify="center">
                      <v-progress-circular color="grey-lighten-5" indeterminate></v-progress-circular>
                    </v-row>
                  </template>
                </v-img>
              </template>
            </div>
          </v-col>
        </v-row>
      </template> -->
      <div class="interaction-section">
        <v-row dense>
          <v-col cols="6">
            <v-icon small @click="toggleReplyForm">{{ showReplyForm ? 'mdi-close' : 'mdi-reply' }}</v-icon>
          </v-col>
          <v-col cols="6" class="interaction-right">
            <nuxt-link :to="'/comments/' + comment.id + '/replies'" class="interaction-link">
              <v-icon small>mdi-comment</v-icon>
              <span class="interaction-text">{{ comment.totalReplies }}</span>
            </nuxt-link>
            <nuxt-link :to="'/comments/' + comment.id" class="interaction-link id-link">
              <span class="interaction-text">ID: {{ comment.id }}</span>
            </nuxt-link>
          </v-col>
        </v-row>
      </div>
      <div v-if="showReplyForm" class="reply-form">
        <CommentForm :formTitle="'返信 >> ' + comment.id" @submit="submitReply" />
      </div>
    </v-list-item>
    <v-divider />

    <MediaModal
      v-if="selectedAttachment"
      :dialog="isOpenMediaModalOpen"
      :type="selectedAttachment.type"
      :url="selectedAttachment.url"
      @close="closeModalMedia"
    />
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router';
import MediaModal from '~/components/MediaModal.vue';
import CommentForm from '~/components/comment/CommentForm.vue';
import type { IThreadComment } from '~/types/thread';
import type { IThreadCommentAttachment } from '~/types/thread-comment-attachment';

const props = defineProps<{ comment: IThreadComment; commentLimit: number; threadId: number }>();

const route = useRoute();

const isOpenMediaModalOpen = ref(false);
const selectedAttachment = ref<IThreadCommentAttachment | null>();
const showReplyForm = ref(false);

function toggleReplyForm() {
  showReplyForm.value = !showReplyForm.value;
}

function submitReply() {
  console.log('返信を送信');
  alert('返信しました。');
  toggleReplyForm();
}

function openModalMedia(attachment: IThreadCommentAttachment) {
  selectedAttachment.value = attachment;
  isOpenMediaModalOpen.value = true;
}

function closeModalMedia() {
  selectedAttachment.value = null;
  isOpenMediaModalOpen.value = false;
}

function onVideoLoad(event) {
  const video = event.target;
  video.currentTime = 1;
  video.pause();
}

function toParentComment(parentCommentIdx: number): {
  path: string;
  query: { offset: number; limit: number };
  hash: string;
} {
  const limit = route.query.limit ? parseInt(route.query.limit as string, 10) : props.commentLimit;
  const newOffset = Math.floor((parentCommentIdx - 1) / limit) * limit;
  return {
    path: `/threads/${props.threadId}`,
    query: { offset: newOffset, limit },
    hash: `#comments-${parentCommentIdx}`,
  };
}
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

.attachment-chip {
  cursor: pointer;
}

.attachment-row {
  flex-direction: column;
}

.reply-form {
  margin-top: 16px;
}
</style>
