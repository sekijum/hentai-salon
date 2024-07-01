<template>
  <div>
    <v-row>
      <v-col cols="3" class="d-flex justify-start align-center">
        <v-btn text class="header-button" @click="router.back()">
          <v-icon>mdi-arrow-left</v-icon>
          <div>戻る</div>
        </v-btn>
      </v-col>
      <v-col cols="6" class="d-flex justify-center align-center">
        <h1 class="header-title" @click="router.push('/')">変態サロン</h1>
      </v-col>
      <v-col cols="3" class="d-flex justify-end align-center">
        <v-btn text class="header-button" @click="router.push('/setting')">
          <v-icon>mdi-cog</v-icon>
          <div>設定</div>
        </v-btn>
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
.header-button {
  display: flex;
  flex-direction: column;
  align-items: center;
  border: none;
  box-shadow: none;
}

.header-title {
  cursor: pointer;
}
</style>
