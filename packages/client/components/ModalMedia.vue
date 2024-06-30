<template>
  <v-dialog v-model="props.dialog" fullscreen hide-overlay transition="dialog-bottom-transition">
    <v-card>
      <v-toolbar dense>
        <v-btn icon @click="closeDialog">
          <v-icon>mdi-close</v-icon>
        </v-btn>
        <v-toolbar-title>Media</v-toolbar-title>
      </v-toolbar>
      <v-container>
        <v-row>
          <v-col v-for="(media, index) in mediaItems" :key="index" cols="12" class="media-col">
            <template v-if="media.type === 'video/mp4'">
              <video controls class="modal-media" :poster="media.thumbnail">
                <source :src="media.url" type="video/mp4" />
              </video>
            </template>
            <template v-else>
              <v-img :src="media.url" class="modal-media"></v-img>
            </template>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue';

const props = defineProps({
  dialog: Boolean,
  mediaItems: Array,
});

const emit = defineEmits(['close']);

const closeDialog = () => {
  emit('close');
};
</script>

<style scoped>
.modal-media {
  width: 100%;
  margin: auto;
  object-fit: contain;
}
</style>
