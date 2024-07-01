<template>
  <v-row class="pa-0 ma-0">
    <v-col v-for="n in 1000" :key="n" class="d-flex child-flex pa-0 ma-0" cols="4">
      <v-img
        :lazy-src="`https://picsum.photos/10/6?image=${n * 5 + 10}`"
        :src="`https://picsum.photos/500/300?image=${n * 5 + 10}`"
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
</template>

<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router';

const router = useRouter();
</script>

<style scoped>
.v-col {
  padding: 2px !important; /* 画像同士の間隔を狭める */
  cursor: pointer; /* マウスオーバー時にカーソルをポインターに変更 */
}
</style>
