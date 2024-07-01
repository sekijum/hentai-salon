<template>
  <div>
    <PageTitle title="板一覧" />

    <v-divider></v-divider>

    <v-autocomplete
      label="板名で検索"
      :items="['California', 'Colorado', 'Florida', 'Georgia', 'Texas', 'Wyoming']"
    ></v-autocomplete>

    <v-row>
      <v-col
        v-for="(board, index) in filteredBoards"
        :key="index"
        cols="12"
        class="board-item"
        style="flex: 0 0 33.3333%; max-width: 33.3333%"
      >
        <v-card @click="navigateToBoard(board.link)" class="board-card">
          <v-img :src="board.image" aspect-ratio="1" class="board-image"></v-img>
          <v-card-title class="board-title">{{ board.title }}</v-card-title>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import PageTitle from '~/components/PageTitle.vue';

const search = ref('');
const boards = ref([
  { title: 'よく見る板', image: 'https://via.placeholder.com/80', link: '/board/1' },
  { title: '風俗・業界', image: 'https://via.placeholder.com/80', link: '/board/2' },
  { title: '地域', image: 'https://via.placeholder.com/80', link: '/board/3' },
  { title: 'エロ表現・創作', image: 'https://via.placeholder.com/80', link: '/board/4' },
  { title: 'エロ文化', image: 'https://via.placeholder.com/80', link: '/board/5' },
  { title: 'エロメディア', image: 'https://via.placeholder.com/80', link: '/board/6' },
  { title: 'えっちな生活', image: 'https://via.placeholder.com/80', link: '/board/7' },
  { title: '画像・動画', image: 'https://via.placeholder.com/80', link: '/board/8' },
  { title: '案内・雑談', image: 'https://via.placeholder.com/80', link: '/board/9' },
  { title: 'その他', image: 'https://via.placeholder.com/80', link: '/board/10' },
  { title: 'エロゲー', image: 'https://via.placeholder.com/80', link: '/board/11' },
  { title: 'WorldWide', image: 'https://via.placeholder.com/80', link: '/board/12' },
]);

const router = useRouter();

const navigateToBoard = link => {
  router.push(link);
};

const filteredBoards = computed(() => {
  return boards.value.filter(board => board.title.includes(search.value));
});
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
