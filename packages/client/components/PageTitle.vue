<template>
  <h1 class="font-weight-regular page-title">{{ title }}</h1>
</template>

<script setup lang="ts">
defineProps({
  title: String,
});
</script>

<style scoped>
.page-title {
  font-size: 1rem;
  text-align: center;
}
</style>
