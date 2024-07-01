<template>
  <v-sheet>
    <v-row align="center" justify="center" class="m-0">
      <template v-for="(item, idx) in items" :key="idx">
        <v-col :cols="columnWidth" class="p-0">
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
  items: [{ title: '', to: '', icon: '' }],
});

import { computed } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const navigateTo = link => {
  router.push(link);
};

const columnWidth = computed(() => {
  return props.items.length <= 3 ? 12 / props.items.length : 4;
});
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
