<template>
  <v-sheet>
    <v-row align="center" justify="center" class="m-0">
      <v-col cols="3" class="p-0">
        <v-sheet class="menu-item" @click="showAll"> 全部 </v-sheet>
      </v-col>
      <v-col cols="3" class="p-0">
        <v-sheet class="menu-item" @click="showPrev"> 前100 </v-sheet>
      </v-col>
      <v-col cols="3" class="p-0">
        <v-sheet class="menu-item" @click="showNext"> 次100 </v-sheet>
      </v-col>
      <v-col cols="3" class="p-0">
        <v-sheet class="menu-item" @click="showLatest"> 最新50 </v-sheet>
      </v-col>
    </v-row>
    <v-row align="center" justify="center" class="m-0">
      <template v-for="(item, idx) in displayedItems" :key="idx">
        <v-col cols="3" class="p-0">
          <v-sheet class="menu-item" @click="navigateTo(item.to)">
            <v-icon class="menu-icon">{{ item.icon }}</v-icon>
            <span class="menu-title">{{ item.title }}</span>
          </v-sheet>
        </v-col>
      </template>
    </v-row>
  </v-sheet>
</template>

<script setup>
const props = defineProps({
  items: {
    type: Array,
    default: () => [],
  },
});

import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const currentPage = ref(1);
const itemsPerPage = ref(100);

const navigateTo = link => {
  router.push(link);
};

const showAll = () => {
  itemsPerPage.value = props.items.length;
  currentPage.value = 1;
};

const showPrev = () => {
  if (currentPage.value > 1) {
    currentPage.value--;
  }
};

const showNext = () => {
  if (currentPage.value * itemsPerPage.value < props.items.length) {
    currentPage.value++;
  }
};

const showLatest = () => {
  itemsPerPage.value = 50;
  currentPage.value = Math.ceil(props.items.length / itemsPerPage.value);
};

const startIdx = computed(() => (currentPage.value - 1) * itemsPerPage.value);
const endIdx = computed(() => startIdx.value + itemsPerPage.value);

const displayedItems = computed(() => props.items.slice(startIdx.value, endIdx.value));
</script>

<style scoped>
.menu-item {
  align-items: center;
  border: 1px solid #ccc; /* ボーダーを追加 */
  box-shadow: none;
  width: 100%; /* カラムと同じ幅に設定 */
  margin: 0; /* マージンを0に設定 */
  padding: 16px; /* 適切なパディングを追加 */
  display: flex;
  justify-content: center;
  cursor: pointer;
  text-align: center; /* テキストを中央揃えにする */
}

.menu-icon {
  margin-right: 8px; /* アイコンとテキストの間にスペースを追加 */
}

.menu-title {
  font-size: 12px; /* 文字サイズを小さくする */
  max-width: calc(100% - 32px); /* タイトルの最大幅をアイコンの幅に合わせて固定 */
  overflow: hidden; /* オーバーフローを隠す */
  text-overflow: ellipsis; /* テキストが溢れた場合に省略記号を表示 */
  white-space: nowrap; /* テキストを一行にする */
}

.v-col {
  padding: 0 !important; /* v-colのパディングを強制的に0に設定 */
}

.v-row {
  margin: 0 !important; /* v-rowのマージンを強制的に0に設定 */
}
</style>
