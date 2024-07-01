<template>
  <v-sheet>
    <v-row align="center" justify="center">
      <v-col cols="4">
        <v-btn block text class="menu-button">
          <v-icon class="menu-icon">mdi-format-list-bulleted</v-icon>
          関連順
        </v-btn>
      </v-col>
      <v-col cols="4">
        <v-btn block text class="menu-button">
          <v-icon class="menu-icon">mdi-fire</v-icon>
          人気
        </v-btn>
      </v-col>
      <v-col cols="4">
        <v-btn block text class="menu-button">
          <v-icon class="menu-icon">mdi-update</v-icon>
          閲覧履歴
        </v-btn>
      </v-col>
      <v-col cols="4">
        <v-btn block text class="menu-button">
          <v-icon class="menu-icon">mdi-earth</v-icon>
          閲覧順
        </v-btn>
      </v-col>
      <v-col cols="4">
        <v-btn block text class="menu-button">
          <v-icon class="menu-icon">mdi-new-box</v-icon>
          新着順
        </v-btn>
      </v-col>
      <v-col cols="4">
        <v-btn block text class="menu-button">
          <v-icon class="menu-icon">mdi-cog</v-icon>
          コメント数
        </v-btn>
      </v-col>
    </v-row>
  </v-sheet>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';

const router = useRouter();

const navigateTo = link => {
  router.push(link);
};
</script>

<style scoped>
.menu-button {
  align-items: center;
  border: none;
  box-shadow: none;
}
</style>
