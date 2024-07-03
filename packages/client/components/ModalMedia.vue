<template>
  <v-dialog
    :model-value="dialog"
    fullscreen
    hide-overlay
    transition="dialog-bottom-transition"
    @update:model-value="updateDialog"
  >
    <v-card class="media-card">
      <v-toolbar dense>
        <v-btn icon @click="closeDialog">
          <v-icon>mdi-close</v-icon>
        </v-btn>
        <v-toolbar-title>Media</v-toolbar-title>
      </v-toolbar>
      <div class="media-container">
        <template v-if="mediaItem.type === 'video/mp4'">
          <VideoPlayer :src="mediaItem.url" />
        </template>
        <template v-else>
          <v-img :src="mediaItem.url" class="media-image"></v-img>
        </template>
      </div>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from 'vue';
import VideoPlayer from '~/components/VideoPlayer.vue';

const props = defineProps({
  dialog: Boolean,
  mediaItem: Object,
});

const emit = defineEmits(['close', 'update:dialog']);

const closeDialog = () => {
  emit('update:dialog', false);
  emit('close');
};

const updateDialog = value => {
  emit('update:dialog', value);
};
</script>

<style scoped>
.media-card {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh; /* Full height to ensure it doesn't cut off */
  margin: 0;
}

.media-container {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: calc(100vh - 64px); /* Adjust height minus toolbar height */
}

.media-image {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
}
</style>
