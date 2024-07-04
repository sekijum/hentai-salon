<template>
  <v-bottom-sheet :model-value="dialog" @update:model-value="updateDialog" max-height="100%">
    <v-card class="media-card">
      <v-toolbar dense class="toolbar">
        <v-toolbar-title>Media</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn icon @click="closeDialog">
          <v-icon>mdi-close</v-icon>
        </v-btn>
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
  </v-bottom-sheet>
</template>

<script setup lang="ts">
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
  justify-content: flex-start; /* Ensure content starts from the top */
  align-items: center;
  height: 100%;
  margin: 0;
  overflow: hidden; /* Prevent overflow issues */
}

.toolbar {
  width: 100%; /* Ensure toolbar takes full width */
  position: sticky;
  top: 0;
  z-index: 1; /* Ensure it stays on top */
  background-color: white; /* Ensure it has a background */
}

.media-container {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%; /* Full height to ensure it doesn't cut off */
}

.media-image {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
}
</style>
