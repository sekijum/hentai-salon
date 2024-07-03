<template>
  <v-infinite-scroll height="100%" :items="items" :onLoad="load">
    <v-row class="pa-0 ma-0">
      <v-col v-for="item in items" :key="item.id" class="d-flex child-flex pa-0 ma-0" cols="6">
        <v-img
          :lazy-src="item.lazySrc"
          :src="item.src"
          aspect-ratio="1"
          class="bg-grey-lighten-2"
          cover
          @click="navigateToExternalPath(`/comments/${item.id}`)"
        >
          <template v-slot:placeholder>
            <v-row align="center" class="fill-height ma-0" justify="center">
              <v-progress-circular color="grey-lighten-5" indeterminate></v-progress-circular>
            </v-row>
          </template>
          <v-icon v-if="item.type === 'video/mp4'" size="60" class="play-icon">mdi-play-circle</v-icon>
          <div v-if="item.type === 'video/mp4'" class="time-label">{{ formatTime(item.duration) }}</div>
        </v-img>
      </v-col>
    </v-row>
  </v-infinite-scroll>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const items = ref([
  ...Array.from({ length: 20 }, (k, v) => ({
    id: v + 1,
    type: 'image',
    lazySrc: `https://picsum.photos/10/6?image=${(v + 1) * 5 + 10}`,
    src: `https://picsum.photos/500/300?image=${(v + 1) * 5 + 10}`,
  })),
  {
    id: 21,
    type: 'video/mp4',
    lazySrc: 'https://picsum.photos/10/6?image=1000',
    src: 'https://picsum.photos/500/300?image=1000',
    videoSrc: 'https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4',
    duration: 120, // 動画の時間を秒単位で指定
  },
]);

async function api() {
  return new Promise(resolve => {
    setTimeout(() => {
      resolve(
        Array.from({ length: 10 }, (k, v) => ({
          id: v + items.value.at(-1).id + 1,
          type: (v + items.value.at(-1).id + 1) % 5 === 0 ? 'video/mp4' : 'image',
          lazySrc: `https://picsum.photos/10/6?image=${(v + items.value.at(-1).id + 1) * 5 + 10}`,
          src: `https://picsum.photos/500/300?image=${(v + items.value.at(-1).id + 1) * 5 + 10}`,
          ...((v + items.value.at(-1).id + 1) % 5 === 0 && {
            thumbnail: 'https://picsum.photos/500/300?image=1000',
            videoSrc: 'https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4',
            duration: 150, // 動画の時間を秒単位で指定
          }),
        })),
      );
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

function formatTime(seconds: number): string {
  const minutes = Math.floor(seconds / 60);
  const remainingSeconds = seconds % 60;
  return `${minutes}:${remainingSeconds.toString().padStart(2, '0')}`;
}
</script>

<style scoped>
.v-col {
  padding: 2px !important; /* 画像同士の間隔を狭める */
  cursor: pointer; /* マウスオーバー時にカーソルをポインターに変更 */
}

.play-icon {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  color: white;
  font-size: 60px;
}

.time-label {
  position: absolute;
  bottom: 8px;
  right: 8px;
  background-color: rgba(0, 0, 0, 0.7);
  color: white;
  padding: 2px 4px;
  border-radius: 3px;
  font-size: 14px;
}
</style>
