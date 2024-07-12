<template>
  <v-overlay :model-value="isLoading" class="align-center justify-center" scroll-strategy="none">
    <v-progress-circular color="primary" size="150" indeterminate>{{ title }}</v-progress-circular>
  </v-overlay>
</template>
<script setup lang="ts">
defineProps<{ isLoading: boolean; title?: string }>();
</script>

<style scoped></style>
