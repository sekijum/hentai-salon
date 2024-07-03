<template>
  <div>
    <PageTitle title="設定" />

    <v-divider></v-divider>

    <div v-for="item in menuItems" :key="item.name" :class="{ highlight: item.type === 'header' }" class="bordered-row">
      <v-row class="align-center">
        <v-col cols="6" class="d-flex justify-center align-center">{{ item.name }}</v-col>
        <v-col cols="6" class="d-flex justify-center align-center">
          <div v-if="item.key" class="d-flex justify-center align-center">
            <v-btn v-if="item.key === 'board-history-delete'" small class="spaced-button">削除</v-btn>
            <v-btn v-else-if="item.key === 'thread-history-delete'" small class="spaced-button">削除</v-btn>
            <v-switch
              v-else-if="item.key === 'display-latest-50-replies-only'"
              flat
              small
              class="centered-switch"
            ></v-switch>
            <v-switch
              v-else-if="item.key === 'dark-mode'"
              flat
              small
              class="centered-switch"
              @change="toggleTheme"
            ></v-switch>
          </div>
        </v-col>
      </v-row>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import PageTitle from '~/components/PageTitle.vue';
import { useNuxtApp } from '#app';
const { $toggleTheme } = useNuxtApp();

const toggleTheme = () => {
  $toggleTheme();
};
const menuItems = [
  { name: '表示', type: 'header' },
  { name: 'ダークモード', key: 'dark-mode' },
  { name: '板', type: 'header' },
  { name: '閲覧履歴', key: 'board-history-delete' },
  { name: 'スレッド', type: 'header' },
  { name: '閲覧履歴', key: 'thread-history-delete' },
  { name: '最新レス50レスのみ表示', key: 'display-latest-50-replies-only' },
  { name: 'コメント', type: 'header' },
  { name: '昇順/降順', key: 'display-latest-50-replies-only' },
];
</script>

<style scoped>
.highlight {
  background-color: #f0f0f0;
}
.bordered-row {
  border-top: 1px solid #ddd;
  border-bottom: 1px solid #ddd;
}
.align-center {
  align-items: center;
}
.centered-switch {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}
.spaced-button {
  margin-top: 8px;
  margin-bottom: 8px;
}
</style>
