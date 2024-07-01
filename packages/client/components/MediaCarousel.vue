<template>
  <v-dialog v-model="dialog" max-width="800px">
    <v-card>
      <v-card-title>
        <v-spacer></v-spacer>
        <v-btn icon @click="closeDialog">
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </v-card-title>
      <v-card-text>
        <v-carousel v-model="currentSlide" hide-delimiters>
          <v-carousel-item v-for="(media, index) in mediaItems" :key="index">
            <v-img v-if="media.type !== 'video/mp4'" :src="media.url" class="media-carousel-img"></v-img>
            <video v-else controls class="media-carousel-img">
              <source :src="media.url" type="video/mp4" />
            </video>
          </v-carousel-item>
        </v-carousel>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script setup>
const props = defineProps({
  dialog: Boolean,
  mediaItems: Array,
  initialIndex: Number,
});

const emit = defineEmits(['close']);

const dialog = ref(props.dialog);
const currentSlide = ref(props.initialIndex);

watch(
  () => props.dialog,
  newVal => {
    dialog.value = newVal;
  },
);

const closeDialog = () => {
  dialog.value = false;
  emit('close');
};
</script>

<style scoped>
.media-carousel-img {
  width: 100%;
  height: auto;
  object-fit: contain;
}
</style>
