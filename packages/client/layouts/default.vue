<template>
  <v-app>
    <v-container class="main-container">
      <Header />
      <main><slot /></main>
    </v-container>
  </v-app>
</template>

<script setup>
import MenuSection from '~/components/MenuSection.vue';
</script>

<style scoped>
.main-container {
  border: 1px solid #e0e0e0;
  max-width: 100%;
  overflow-x: hidden;
  padding: 0;
}
</style>
