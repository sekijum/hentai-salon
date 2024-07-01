<template>
  <div class="mx-2">
    <PageTitle title="スレ作成" />

    <v-divider></v-divider>

    <v-form @submit.prevent="submitForm">
      <v-text-field
        v-model="form.title"
        label="タイトル"
        :rules="[rules.required, rules.max50]"
        required
      ></v-text-field>

      <v-textarea v-model="form.description" label="説明" :rules="[rules.max255]"></v-textarea>

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
      />

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
});

const rules = {
  required: (value: string) => !!value || 'Required.',
  max50: (value: string) => value.length <= 50 || 'Max 50 characters',
  max255: (value: string) => value.length <= 255 || 'Max 255 characters',
  file: (value: File | null) => !value || value.size < 2000000 || 'File size should be less than 2 MB', // Example file size validation
};

const submitForm = () => {
  const formData = new FormData();
  formData.append('title', form.value.title);
  if (form.value.description) formData.append('description', form.value.description);
  if (form.value.thumbnailFile) formData.append('thumbnailFile', form.value.thumbnailFile);
  formData.append('isNotifyOnComment', String(form.value.isNotifyOnComment));

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
