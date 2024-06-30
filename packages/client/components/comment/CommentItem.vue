<template>
  <div class="comment-item">
    <v-list-item class="comment-list-item">
      <v-list-item-content class="comment-list-item-content">
        <div class="d-flex justify-space-between comment-header">
          <div class="comment-header-text">
            {{ idx }} <a :href="'/user/' + comment.username" class="username-link">{{ comment.username }}</a
            >[Lv.{{ comment.level }}]{{ comment.rank }} {{ comment.date }} {{ comment.time }}
          </div>
          <div class="interaction-section">
            <v-icon small>mdi-thumb-up</v-icon>
            <v-icon small>mdi-reply</v-icon>
            <span class="reply-text">返信</span>
          </div>
        </div>
        <v-list-item-title class="comment-content">{{ comment.content }}</v-list-item-title>
        <template v-if="comment.media && comment.media.length">
          <v-row dense>
            <v-col cols="3" v-for="(media, index) in comment.media" :key="index" class="media-col">
              <div v-if="media.type === 'video/mp4'" class="media-item video-thumbnail" @click="openDialog(comment.media)">
                <v-img :src="media.thumbnail" class="media-item"></v-img>
                <v-icon size="40" class="play-icon">mdi-play-circle</v-icon>
              </div>
              <v-img v-else :src="media.url" class="media-item" @click="openDialog(comment.media)"></v-img>
            </v-col>
          </v-row>
        </template>
      </v-list-item-content>
    </v-list-item>
    <v-divider></v-divider>

    <ModalMedia :dialog="dialog" :mediaItems="comment.media" @close="() => (dialog = false)" />
  </div>
</template>

<script setup>
import { ref } from 'vue';
import ModalMedia from '~/components/ModalMedia.vue';

defineProps({
  idx: Number,
  comment: Object,
});

const dialog = ref(false);

const openDialog = media => {
  dialog.value = true;
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
  color: blue;
}

.interaction-section {
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

.media-col {
  padding: 1px;
}

.media-item {
  width: 100%;
  height: 150px;
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

.modal-media {
  width: 100%;
  margin: auto;
  object-fit: contain;
}
</style>
