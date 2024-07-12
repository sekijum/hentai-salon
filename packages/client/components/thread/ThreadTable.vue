<template>
  <div v-if="items.length">
    <div v-if="title" class="section-title">
      <h2 class="font-weight-regular">{{ title }}</h2>
    </div>

    <div class="thread-section">
      <div
        v-for="(item, index) in items"
        :key="item.id"
        :class="{ alternate: index % 2 === 0 }"
        @click="() => router.push(`/threads/${item.id}`)"
        class="d-flex align-center p-2 item-row"
      >
        <v-row>
          <v-col cols="3" class="d-flex align-center">
            <div class="fixed-image mr-1">
              <v-img :src="getImageSrc(item.thumbnailUrl)" class="image"></v-img>
            </div>
          </v-col>
          <v-col cols="9" class="d-flex flex-column justify-center">
            <p class="item-title">
              {{ truncateTitle(item.title) }}
            </p>
            <div class="item-details">
              <small> {{ item.board?.title }} <v-icon small>mdi-comment</v-icon> {{ item.commentCount }} </small>
            </div>
          </v-col>
        </v-row>
      </div>
    </div>

    <div v-if="navigate && items.length > 0" class="more-link" @click="navigate">
      {{ title }}をもっと見る <v-icon down>mdi-chevron-down</v-icon>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import type { IThread } from '~/types/thread';

const nuxtApp = useNuxtApp();

defineProps<{
  title: String;
  items: IThread[];
  navigate?: () => void;
}>();

const router = useRouter();

function truncateTitle(title: string) {
  return title.length > 50 ? title.slice(0, 50) + '...' : title;
}

function getImageSrc(thumbnailUrl: string) {
  return thumbnailUrl ? thumbnailUrl : '/no-image.jpg';
}
</script>

<style scoped>
.thread-section {
  cursor: pointer;
}

.section-title h2 {
  color: orange;
}

.alternate {
  background-color: #f5f5f5;
}

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

.more-link {
  text-align: center;
  cursor: pointer;
  background-color: #f0f0f0;
  padding: 10px;
  text-decoration: underline;
}
</style>
