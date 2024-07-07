<template>
  <div>
    <div v-if="title" class="section-title">
      <h2 class="font-weight-regular">{{ title }}</h2>
    </div>

    <v-data-table :headers="headers" :items="items" hide-default-footer hide-default-header class="thread-section">
      <template v-slot:item="{ item, index }">
        <div
          :class="{ alternate: index % 2 === 0 }"
          @click="() => router.push(`/threads/${item.id}`)"
          class="d-flex align-center p-2 item-row"
        >
          <div class="fixed-image mr-1">
            <v-img :src="getImageSrc(item.thumbnailUrl)" class="image"></v-img>
          </div>
          <div class="flex-grow-1">
            <p class="item-title">
              {{ truncateTitle(item.title) }}
            </p>
          </div>
          <div class="text-right mr-2">
            <small>
              {{ $formatDate(item.createdAt) }}
              <br />
              {{ item.board?.title }}
              <br />
              話題度
              {{ item.popularity }}
              <br />
              {{ item.commentCount }}
              <v-icon small>mdi-comment</v-icon>
              <br />
            </small>
          </div>
        </div>
      </template>
    </v-data-table>

    <div v-if="moreLink && items.length > 0" class="more-link" @click="() => router.push(moreLink)">
      {{ title }}をもっと見る <v-icon down>mdi-chevron-down</v-icon>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import type { TThread, TThreadList } from '~/types/thread';

const nuxtApp = useNuxtApp();
const { $formatDate } = nuxtApp;

defineProps<{
  title: String;
  items: TThread[];
  moreLink: string;
  maxItems: number;
}>();

const router = useRouter();

const truncateTitle = (title: string) => {
  return title.length > 50 ? title.slice(0, 50) + '...' : title;
};

const headers = [
  { text: '', value: 'image', width: '20%' },
  { text: 'Title', value: 'title', width: '70%' },
  { text: 'Comments/Board', value: 'commentsBoard', width: '10%' },
];

const getImageSrc = (thumbnailUrl: string) => {
  return thumbnailUrl ? thumbnailUrl : '/no-image.jpg';
};
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

.more-link {
  text-align: center;
  cursor: pointer;
  background-color: #f0f0f0;
  padding: 10px;
  text-decoration: underline;
}
</style>
