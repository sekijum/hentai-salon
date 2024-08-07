<template>
  <v-app class="main-container">
    <Header />

    <v-main><NuxtPage /></v-main>

    <MenuModal />
  </v-app>
</template>

<script setup lang="ts">
import Header from '~/components/Header.vue';
import MenuModal from '~/components/MenuModal.vue';
</script>

<style scoped>
.main-container {
  overflow: hidden;
}
</style>
