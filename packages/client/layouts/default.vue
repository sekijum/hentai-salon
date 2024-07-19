<template>
  <v-app class="main-container">
    <Header />

    <main><NuxtPage /></main>

    <MenuModal />
  </v-app>
</template>

<script setup lang="ts">
import MenuModal from '~/components/MenuModal.vue';
import Header from '~/components/Header.vue';
</script>

<style scoped>
.main-container {
  border: 1px solid #e0e0e0;
  max-width: 100%;
  overflow-x: hidden;
  padding: 0;
}
</style>
