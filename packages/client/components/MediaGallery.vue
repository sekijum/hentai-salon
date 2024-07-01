<template>
  <v-infinite-scroll height="100%" :items="items" :onLoad="load">
    <v-row class="pa-0 ma-0">
      <v-col v-for="n in items" :key="n" class="d-flex child-flex pa-0 ma-0" cols="6">
        <v-img
          :lazy-src="`https://picsum.photos/10/6?image=${n * 5 + 10}`"
          :src="`https://picsum.photos/500/300?image=${n * 5 + 10}`"
          aspect-ratio="1"
          class="bg-grey-lighten-2"
          cover
          @click="navigateToExternalPath(`/comments/${n}`)"
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
</template>

<script setup lang="ts">
import { ref } from 'vue';

const items = ref(Array.from({ length: 30 }, (k, v) => v + 1));

async function api() {
  return new Promise(resolve => {
    setTimeout(() => {
      resolve(Array.from({ length: 10 }, (k, v) => v + items.value.at(-1) + 1));
    }, 1000);
  });
}

async function load({ done }) {
  const res = await api();
  items.value.push(...res);
  done('ok');
}

function navigateToExternalPath(path: string) {
  const domain = window.location.origin;
  window.open(`${domain}${path}`, '_blank');
}
</script>

<style scoped>
.v-col {
  padding: 2px !important; /* 画像同士の間隔を狭める */
  cursor: pointer; /* マウスオーバー時にカーソルをポインターに変更 */
}
</style>
