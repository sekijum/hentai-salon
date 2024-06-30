<template>
  <div class="comment-item">
    <v-list-item class="comment-list-item">
      <v-list-item-content class="comment-list-item-content">
        <div class="comment-header">
          <div class="comment-header-text">
            {{ idx }} <router-link :to="'/user/' + comment.username" class="username-link">{{ comment.username }}</router-link
            >[Lv.{{ comment.level }}]{{ comment.rank }} {{ comment.date }} {{ comment.time }}
          </div>
        </div>
        <v-list-item-title class="comment-content">{{ comment.content }}</v-list-item-title>
        <template v-if="comment.media && comment.media.length">
          <v-row dense>
            <v-col cols="3" v-for="(media, index) in comment.media" :key="index" class="media-col">
              <div v-if="media.type === 'video/mp4'" class="media-item video-thumbnail" @click="() => (dialog = true)">
                <v-img :src="media.thumbnail" class="media-item"></v-img>
                <v-icon size="40" class="play-icon">mdi-play-circle</v-icon>
              </div>
              <v-img v-else :src="media.url" class="media-item" @click="() => (dialog = true)"></v-img>
            </v-col>
          </v-row>
        </template>
        <div class="interaction-section">
          <v-row dense>
            <v-col cols="6">
              <router-link :to="'/reply/' + comment.id" class="interaction-link">
                <v-icon small>mdi-reply</v-icon>
              </router-link>
            </v-col>
            <v-col cols="6" class="interaction-right">
              <router-link :to="'/comment/' + comment.id" class="interaction-link">
                <v-icon small>mdi-comment</v-icon>
                <span class="interaction-text">{{ comment.commentCount }}</span>
              </router-link>
              <span class="interaction-divider"></span>
              <router-link :to="'/comment/' + comment.id" class="interaction-link">
                <span class="interaction-text">ID: {{ comment.id }}</span>
              </router-link>
            </v-col>
          </v-row>
        </div>
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

.interaction-section {
  display: flex;
  align-items: center;
}

.interaction-link {
  display: flex;
  align-items: center;
  text-decoration: none;
  color: grey; /* 色をグレーに設定 */
}

.interaction-text {
  color: grey; /* 色をグレーに設定 */
}

.interaction-divider {
  width: 16px; /* コメント数とコメントIDの間にスペースを追加 */
}

.interaction-right {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}

.comment-content {
  font-size: 14px;
  white-space: normal;
  word-wrap: break-word;
}

.media-item {
  width: 100%;
  height: auto;
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
