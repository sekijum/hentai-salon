<template>
  <v-bottom-sheet :model-value="!!url" @update:model-value="closeDialog" max-height="100%">
    <v-card class="media-card">
      <v-toolbar dense class="toolbar">
        <v-toolbar-title>
          <nuxt-link :to="to" target="_blank" rel="noopener noreferrer">{{ 'Media' }}</nuxt-link>
        </v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn icon @click="closeDialog">
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </v-toolbar>
      <div class="media-container">
        <template v-if="type === 'video'">
          <VideoPlayer :src="url" />
        </template>
        <template v-else>
          <v-img :src="url" class="media-image" referrerpolicy="no-referrer" />
        </template>
      </div>
    </v-card>
  </v-bottom-sheet>
</template>

<script setup lang="ts">
import VideoPlayer from '~/components/VideoPlayer.vue';

defineProps<{ type: 'video' | 'image'; url: string; to?: string }>();

const emit = defineEmits(['close']);

function closeDialog() {
  emit('close');
}
</script>

<style scoped>
.media-card {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
  height: 100%;
  margin: 0;
  overflow: hidden;
}

.toolbar {
  width: 100%;
  position: sticky;
  top: 0;
  z-index: 1;
  background-color: white;
}

.media-container {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
}

.media-image {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
}
</style>
