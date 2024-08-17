<template>
  <div>
    <h1>変態サロン管理画面</h1>
  </div>
</template>

<script setup>
definePageMeta({
  layout: 'admin',
  middleware: ['admin-access-only'],
});
</script>

<style scoped></style>
