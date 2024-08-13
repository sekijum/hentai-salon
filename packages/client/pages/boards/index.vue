<template>
  <div>
    <PageTitle title="板一覧" />

    <v-divider />

    <br />

    <v-row>
      <v-col v-for="(board, index) in boards" :key="index" cols="6">
        <v-card @click="() => router.push(`/threads?filter=board&boardId=${board.id}`)">
          <v-card-title class="board-title">{{ board.title }}</v-card-title>
        </v-card>
      </v-col>
    </v-row>

    <br />
  </div>
</template>

<script setup lang="ts">
import PageTitle from '~/components/PageTitle.vue';
import type { IBoard } from '~/types/board';

const router = useRouter();
const nuxtApp = useNuxtApp();

const { $api } = nuxtApp;

const boards = ref<IBoard[]>([]);

onMounted(async () => {
  await fetchBoards();
});

async function fetchBoards() {
  const response = await $api.get<IBoard[]>('/boards');
  boards.value = response.data;
}

useHead({
  title: '変態サロン | 板一覧',
});
</script>

<style scoped>
.board-title {
  font-size: 0.75rem;
  text-align: center;
}
</style>
