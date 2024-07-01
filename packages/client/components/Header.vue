<template>
  <div>
    <v-row>
      <v-col cols="3" class="d-flex justify-start align-center">
        <v-sheet class="menu-item" @click="router.back()">
          <v-icon class="menu-icon">mdi-arrow-left</v-icon>
          <span class="menu-title">戻る</span>
        </v-sheet>
      </v-col>
      <v-col cols="6" class="d-flex justify-center align-center">
        <h1 class="font-weight-regular header-title" @click="router.push('/')">変態サロン</h1>
      </v-col>
      <v-col cols="3" class="d-flex justify-end align-center">
        <v-sheet class="menu-item" @click="router.push('/setting')">
          <v-icon class="menu-icon">mdi-cog</v-icon>
          <span class="menu-title">設定</span>
        </v-sheet>
      </v-col>
    </v-row>
    <v-divider></v-divider>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';

const router = useRouter();
</script>

<style scoped>
.menu-item {
  display: flex;
  align-items: center;
  padding: 8px;
  cursor: pointer;
  border: none;
  box-shadow: none;
}

.menu-icon {
  margin-right: 8px;
}

.header-title {
  cursor: pointer;
}
</style>
