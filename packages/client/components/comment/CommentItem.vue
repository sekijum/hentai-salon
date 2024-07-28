<template>
  <div class="comment-item">
    <v-list-item class="comment-list-item">
      <div class="comment-header">
        <div class="comment-header-text">
          <template v-if="comment?.user?.id && comment?.user?.profileLink">
            <nuxt-link :to="comment.user.profileLink" class="username-link" target="_blank" rel="noopener">
              {{ username() }}
            </nuxt-link>
          </template>
          <template v-else>
            {{ username() }}
          </template>
          {{ $formatDate(comment.createdAt) }}
        </div>
      </div>
      <div v-if="comment.parentCommentId" class="reply-indication">
        <template v-if="comment.parentCommentId">
          <nuxt-link
            class="reply-link"
            target="_blank"
            rel="noopener"
            :to="`/threads/${threadId}/comments/${comment.parentCommentId}`"
          >
            >> {{ comment.parentCommentId }}
          </nuxt-link>
        </template>
        <template v-else-if="comment.parentCommentId">
          <nuxt-link
            class="reply-link"
            target="_blank"
            rel="noopener"
            :to="`/threads/${threadId}/comments/${comment.parentCommentId}`"
          >
            >> {{ comment.parentCommentId }}
          </nuxt-link>
        </template>
      </div>
      <v-list-item-title class="comment-content">
        <OEmbedContent :text="comment.content" />
      </v-list-item-title>
      <template v-if="comment.attachments && comment.attachments.length">
        <v-row dense>
          <v-col
            cols="3"
            v-for="(attachment, index) in comment.attachments"
            :key="index"
            class="media-col"
            @click="openModalMedia(attachment)"
          >
            <div class="media-item-wrapper">
              <v-img
                :src="`${attachment.url}?w=5`"
                :srcset="`${attachment.url}?w=5 5w, ${attachment.url}?w=5 5w`"
                class="media-item"
                contain
                referrerpolicy="no-referrer"
                :alt="comment.content"
              >
                <template v-slot:placeholder>
                  <v-row align="center" class="fill-height ma-0" justify="center">
                    <v-progress-circular color="grey-lighten-5" indeterminate></v-progress-circular>
                  </v-row>
                </template>
              </v-img>
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
            <nuxt-link
              :to="`/threads/${threadId}/comments/${comment.id}`"
              class="interaction-link"
              target="_blank"
              rel="noopener"
            >
              <v-icon small>mdi-comment</v-icon>
              <span class="interaction-text">{{ comment.replyCount }}</span>
            </nuxt-link>
            <v-icon class="interaction-link id-link" small @click="toggleLike">
              {{ isLiked ? 'mdi-thumb-up' : 'mdi-thumb-up-outline' }}
            </v-icon>
            <nuxt-link
              :to="`/threads/${threadId}/comments/${comment.id}`"
              class="interaction-link id-link"
              target="_blank"
              rel="noopener"
            >
              <span class="interaction-text">ID: {{ comment.id }}</span>
            </nuxt-link>
          </v-col>
        </v-row>
      </div>
      <div v-if="showReplyForm" class="reply-form">
        <CommentForm
          :threadId="threadId"
          :title="'返信 >> ' + comment.id"
          :parentCommentId="comment.id"
          @submit="
            () => {
              toggleReplyForm();
              emit('replied');
            }
          "
          :showReplyForm="showReplyForm"
        />
      </div>
    </v-list-item>
    <v-divider />

    <MediaModal
      v-if="selectedAttachment"
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
import OEmbedContent from '~/components/OEmbedContent.vue';
import type { IThreadComment } from '~/types/thread';
import type { IThreadCommentAttachment } from '~/types/thread-comment-attachment';

const props = defineProps<{ comment: IThreadComment; commentLimit: number; threadId: number }>();
const nuxtApp = useNuxtApp();
const route = useRoute();

const { $formatDate, $api, payload } = nuxtApp;

const selectedAttachment = ref<IThreadCommentAttachment | null>();
const showReplyForm = ref(false);
const emit = defineEmits(['replied']);
const isLiked = ref(props.comment.isLiked);

function toggleReplyForm() {
  showReplyForm.value = !showReplyForm.value;
}

function openModalMedia(attachment: IThreadCommentAttachment) {
  selectedAttachment.value = attachment;
}

function closeModalMedia() {
  selectedAttachment.value = null;
}

function username() {
  const name = props.comment?.user?.name || props.comment?.guestName || '匿名';
  return `${name}`;
}

async function toggleLike() {
  if (!payload.isLoggedIn) {
    alert('ログインしてください。');
    return;
  }
  if (isLiked.value) {
    await $api.post(`/threads/${props.threadId}/comments/${props.comment.id}/unlike`);
    isLiked.value = false;
  } else {
    await $api.post(`/threads/${props.threadId}/comments/${props.comment.id}/like`);
    isLiked.value = true;
  }
}
</script>

<style scoped>
.comment-item {
  font-size: 12px;
}

.comment-list-item {
  padding-left: 5px !important;
  padding-right: 5px !important;
}

.comment-header-text {
  flex: 1;
  word-wrap: break-word;
  font-size: 12px;
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
  margin-left: 15px;
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

.media-item-wrapper {
  position: relative;
  width: 100%;
  padding-bottom: 100%;
  background-color: black;
}

.media-item {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}
</style>
