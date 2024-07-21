<template>
  <div class="text-center">
    <v-snackbar v-model="snackbar.isSnackbar">
      {{ snackbar.text }}
    </v-snackbar>
  </div>
</template>

<script setup lang="ts">
const snackbar = useState('isSnackbar', () => {
  return { isSnackbar: false, text: null };
});
</script>
