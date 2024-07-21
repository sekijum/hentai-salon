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
  overflow: hidden;
}
</style>
