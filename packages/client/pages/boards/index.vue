<template>
  <div>
    <PageTitle title="板一覧" />

    <v-divider></v-divider>

    <br />

    <v-row class="">
      <v-col v-for="(board, index) in filteredBoards" :key="index" cols="4" class="board-item">
        <v-card @click="() => router.push(`/boards/${board.id}`)" class="board-card">
          <v-img :src="getImageSrc(board.thumbnailUrl)" aspect-ratio="1" class="board-image">
            <template v-slot:placeholder>
              <v-row align="center" class="fill-height ma-0" justify="center">
                <v-progress-circular color="grey-lighten-5" indeterminate></v-progress-circular>
              </v-row>
            </template>
          </v-img>
          <v-card-title class="board-title">{{ board.title }}</v-card-title>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script setup lang="ts">
import PageTitle from '~/components/PageTitle.vue';
import type { TBoard } from '~/types/board';

const nuxtApp = useNuxtApp();
const { $api } = nuxtApp;

const boards = ref<TBoard[]>([]);

const search = ref('');

const router = useRouter();

const filteredBoards = computed(() => {
  return boards.value.filter(board => board.title.includes(search.value));
});

onMounted(async () => {
  await fetchBoards();
});

async function fetchBoards() {
  const response = await $api.get<TBoard[]>('/boards');
  boards.value = response.data;
}

function getImageSrc(thumbnailUrl: string) {
  return thumbnailUrl ? thumbnailUrl : '/no-image.jpg';
}
</script>

<style scoped>
.board-image {
  width: 100%;
}

.board-title {
  font-size: 0.75rem;
  text-align: center;
}
</style>
