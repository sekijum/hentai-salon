<template>
  <div v-if="items.length">
    <div v-if="title" class="section-title">
      <h2 class="font-weight-regular">{{ title }}</h2>
    </div>

    <div class="thread-section">
      <template v-if="isInfiniteScroll">
        <v-infinite-scroll :onLoad="load">
          <div
            v-for="(item, index) in items"
            :key="item.id"
            :class="{ alternate: index % 2 === 0 }"
            class="d-flex align-center p-2 item-row"
          >
            <v-row>
              <v-col cols="3" class="d-flex align-center">
                <div class="fixed-image mr-1">
                  <v-img :src="getImageSrc(item.thumbnailUrl)" class="image"></v-img>
                </div>
              </v-col>
              <v-col cols="9" class="d-flex flex-column justify-center">
                <p class="item-title" @click="() => router.push(`/threads/${item.id}`)">
                  {{ truncateTitle(item.title) }}
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
          <template v-slot:empty />
        </v-infinite-scroll>
      </template>
      <template v-else>
        <div
          v-for="(item, index) in items"
          :key="item.id"
          :class="{ alternate: index % 2 === 0 }"
          class="d-flex align-center p-2 item-row"
        >
          <v-row>
            <v-col cols="3" class="d-flex align-center">
              <div class="fixed-image mr-1">
                <v-img :src="getImageSrc(item.thumbnailUrl)" class="image"></v-img>
              </div>
            </v-col>
            <v-col cols="9" class="d-flex flex-column justify-center">
              <p class="item-title" @click="() => router.push(`/threads/${item.id}`)">
                {{ truncateTitle(item.title) }}
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

        <div v-if="clicked" class="more-link" @click="clicked">
          {{ title }}をもっと見る<v-icon down>mdi-chevron-down</v-icon>
        </div>
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router';
import type { IThread } from '~/types/thread';

const nuxtApp = useNuxtApp();

const props = defineProps<{
  title?: string;
  items: IThread[];
  clicked?: () => void;
  isInfiniteScroll?: boolean;
  filter: string;
  threadLimit: number;
}>();

const router = useRouter();
const route = useRoute();

const { $api } = nuxtApp;
const { getThreadViewHistory } = useStorage();
const threadLimit = ref(props.threadLimit);
const offset = ref(0);
const items = ref<IThread[]>([...props?.items]);

function truncateTitle(title: string) {
  return title.length > 50 ? title.slice(0, 50) + '...' : title;
}

function getImageSrc(thumbnailUrl: string) {
  return thumbnailUrl ? thumbnailUrl : '/no-image.jpg';
}

async function load({ done }: { done: (status: 'loading' | 'error' | 'empty' | 'ok') => void }) {
  offset.value += threadLimit.value;
  const { canNextLoad } = await fetchLoadThreads(offset.value);
  canNextLoad ? done('ok') : done('empty');
}

async function fetchLoadThreads(offset: number) {
  console.log(threadLimit.value, offset, props.filter);
  const response = await $api.get<IThread[]>('/threads/', {
    params: {
      filter: props.filter,
      threadIds: getThreadViewHistory(),
      keyword: route.query.keyword,
      boardId: route.query.boardId,
      limit: threadLimit.value,
      offset: offset || 0,
    },
  });

  if (!response.data || response.data.length < threadLimit.value) {
    return { canNextLoad: false };
  }
  items.value = [...items.value, ...response.data];
  return { canNextLoad: true };
}

watch(
  () => route.query.keyword,
  () => {
    offset.value = 0;
    items.value = [];
    fetchLoadThreads(0);
  },
);
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

.interaction-text {
  color: grey;
}
</style>
