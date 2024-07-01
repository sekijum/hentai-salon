<template>
  <v-container fluid>
    <v-infinite-scroll :items="images" :load-more="loadMoreImages" :loader-height="40">
      <v-row class="pa-0 ma-0">
        <v-col v-for="(image, index) in images" :key="image" class="d-flex child-flex pa-0 ma-0" cols="4">
          <v-img
            :lazy-src="`https://picsum.photos/10/6?image=${image}`"
            :src="`https://picsum.photos/500/300?image=${image}`"
            aspect-ratio="1"
            class="bg-grey-lighten-2"
            cover
            @click="router.push('/comments/1')"
          >
            <template v-slot:placeholder>
              <v-row align="center" class="fill-height ma-0" justify="center">
                <v-progress-circular color="grey-lighten-5" indeterminate></v-progress-circular>
              </v-row>
            </template>
          </v-img>
        </v-col>
      </v-row>
    </v-infinite-scroll>
  </v-container>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const images = ref<number[]>([]);
let currentBatch = 0;
const batchSize = 20;

const loadMoreImages = () => {
  return new Promise(resolve => {
    setTimeout(() => {
      const newImages = [];
      for (let i = 0; i < batchSize; i++) {
        newImages.push(currentBatch * batchSize + i * 5 + 10);
      }
      images.value = [...images.value, ...newImages];
      currentBatch++;
      resolve();
    }, 1000); // Simulating a network request delay
  });
};

// Initial load
loadMoreImages();
</script>

<style scoped>
.v-col {
  padding: 2px !important; /* 画像同士の間隔を狭める */
  cursor: pointer; /* マウスオーバー時にカーソルをポインターに変更 */
}
</style>
