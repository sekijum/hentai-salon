<template>
  <div v-for="(item, index) in items" :key="item.id" class="d-flex align-center p-2 item-row">
    <v-row>
      <v-col cols="3" class="d-flex align-center">
        <div class="fixed-image mr-1">
          <v-img :src="getImageSrc(item.thumbnailUrl)" class="image"></v-img>
        </div>
      </v-col>
      <v-col cols="9" class="d-flex flex-column justify-center">
        <p class="item-title" @click="() => router.push(`/threads/${item.id}`)">
          {{ item.title }}
        </p>
        <div class="item-details">
          <small>
            <span @click="() => router.push(`/threads?filter=board&boardId=${item?.board?.id}`)">
              {{ item.board?.title }}
            </span>
            /
            <span @click="() => router.push(`/threads/${item.id}`)">
              <v-icon small color="grey">mdi-comment</v-icon>
              <span>{{ item.commentCount }}</span>
            </span>
          </small>
        </div>
      </v-col>
    </v-row>
  </div>
</template>

<script setup lang="ts">
import type { IThread } from '~/types/thread';

const nuxtApp = useNuxtApp();
const router = useRouter();

const props = defineProps<{
  items: IThread[];
}>();

function getImageSrc(thumbnailUrl: string) {
  return thumbnailUrl ? thumbnailUrl : '/no-image.jpg';
}
</script>

<style scoped>
.fixed-image {
  width: 100px;
  height: 100px;
  flex-shrink: 0;
}

.fixed-image .image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.item-row {
  border-top: 1px solid #ccc;
  border-bottom: 1px solid #ccc;
}

.item-title {
  margin: 0;
  display: flex;
  align-items: center;
  height: 100%;
}

.item-details {
  display: flex;
  justify-content: flex-end;
  align-items: flex-end;
  flex-grow: 1;
  text-align: right;
}

.v-row,
.v-col {
  margin: 0;
  padding: 0;
}
</style>
