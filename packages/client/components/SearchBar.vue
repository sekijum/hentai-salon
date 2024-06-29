<template>
  <div class="search-bar">
    <v-autocomplete label="話題を検索" :items="['California', 'Colorado', 'Florida', 'Georgia', 'Texas', 'Wyoming']"></v-autocomplete>
  </div>
</template>

<script setup></script>

<style scoped>
.search-bar {
  margin-top: 16px;
}
</style>
