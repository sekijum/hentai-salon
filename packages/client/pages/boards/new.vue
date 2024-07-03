<template>
  <div class="mx-2">
    <PageTitle title="板作成" />

    <v-divider></v-divider>

    <v-form @submit.prevent="submitForm">
      <v-text-field v-model="form.title" label="タイトル" required></v-text-field>

      <v-textarea v-model="form.description" label="説明" required></v-textarea>

      <v-file-input
        v-model="form.thumbnailFile"
        label="サムネイルを選択"
        multiple
        show-size
        truncate-length="25"
        prepend-icon=""
        variant="outlined"
        dense
        hide-details
        accept="image/*"
        :rules="[rules.file]"
      >
        <template v-slot:loader>
          <v-progress-linear
            :active="custom"
            :color="color"
            :model-value="progress"
            height="2"
            indeterminate
          ></v-progress-linear> </template
      ></v-file-input>

      <br />
      <v-btn type="submit" color="primary" block>作成</v-btn>
      <p class="note">＊反映には時間が掛かる場合があります＊</p>
    </v-form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const form = ref({
  title: '',
  description: '',
  thumbnailFile: null as File | null,
  isNotifyOnComment: false,
  tags: [] as string[], // Added for tags
});

const value = ref('');
const custom = ref(false);

const progress = computed(() => Math.min(100, value.value.length * 10));
const color = computed(() => ['error', 'warning', 'success'][Math.floor(progress.value / 40)]);
const tagItems = ['foo', 'bar', 'fizz', 'buzz']; // Example tag items, adjust as needed

const rules = {
  required: (value: string) => !!value || 'Required.',
  max50: (value: string) => value.length <= 50 || 'Max 50 characters',
  max255: (value: string) => value.length <= 255 || 'Max 255 characters',
  file: (value: File | null) => !value || value.size < 2000000 || 'File size should be less than 2 MB', // Example file size validation
};

const submitForm = () => {
  custom.value = !custom.value;
  const formData = new FormData();
  formData.append('title', form.value.title);
  if (form.value.description) formData.append('description', form.value.description);
  if (form.value.thumbnailFile) formData.append('thumbnailFile', form.value.thumbnailFile);
  formData.append('isNotifyOnComment', String(form.value.isNotifyOnComment));
  form.value.tags.forEach(tag => formData.append('tags[]', tag)); // Append tags to the form data

  // ここでフォーム送信のロジックを追加します。例えば:
  // axios.post('/your-endpoint', formData)
  console.log('Form data:', form.value);
};
</script>

<style scoped>
.note {
  font-size: 12px;
  color: grey;
  text-align: center;
  margin-top: 8px;
}
</style>
