<template>
  <div class="comment-item" :id="'comments-' + idx">
    <v-list-item class="comment-list-item">
      <v-list-item-content class="comment-list-item-content">
        <div class="comment-header">
          <div class="comment-header-text">
            {{ idx }}
            <router-link :to="'/user/' + comment.username" class="username-link">{{ comment.username }}</router-link
            >[Lv.{{ comment.level }}]{{ comment.rank }} {{ comment.date }} {{ comment.time }}
          </div>
        </div>
        <div v-if="comment.replyTo" class="reply-indication">
          <router-link :to="'/comments/' + comment.replyTo" class="reply-link">>>{{ comment.replyTo }}</router-link>
        </div>
        <v-list-item-title class="comment-content">
          <a :href="'#comment-' + comment.id" class="comment-anchor">{{ comment.content }}</a>
        </v-list-item-title>
        <template v-if="comment.media && comment.media.length">
          <v-row dense>
            <v-col cols="3" v-for="(media, index) in comment.media" :key="index" class="media-col">
              <div
                v-if="media.type === 'video/mp4'"
                class="media-item video-thumbnail"
                @click="() => openDialog(media)"
              >
                <v-img :src="media.thumbnail" class="media-item"></v-img>
                <v-icon size="40" class="play-icon">mdi-play-circle</v-icon>
              </div>
              <v-img v-else :src="media.url" class="media-item" @click="() => openDialog(media)"></v-img>
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
                <span class="interaction-text">{{ comment.commentCount }}</span>
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
      </v-list-item-content>
    </v-list-item>
    <v-divider></v-divider>

    <ModalMedia :dialog="dialog" :mediaItem="selectedMedia" @close="closeDialog" @update:dialog="dialog = $event" />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import ModalMedia from '~/components/ModalMedia.vue';
import CommentForm from '~/components/comment/CommentForm.vue';

defineProps({
  idx: Number,
  comment: Object,
});

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

const openDialog = media => {
  selectedMedia.value = media;
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

.media-item {
  width: 100px; /* Fixed width */
  height: 100px; /* Fixed height */
  object-fit: cover;
  cursor: pointer;
  position: relative;
}

.video-thumbnail {
  position: relative;
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
