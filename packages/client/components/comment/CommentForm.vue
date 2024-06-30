<template>
  <v-card>
    <v-card-title>書き込み</v-card-title>
    <v-card-text>
      <v-form @submit.prevent="submitForm">
        <v-text-field v-model="name" label="名前(省略可)" outlined></v-text-field>
        <v-text-field v-model="email" label="E-mail(省略可)" outlined></v-text-field>
        <v-textarea v-model="comment" label="コメント" rows="4" outlined></v-textarea>
        <input type="file" multiple @change="handleFileChange" style="display: none" ref="fileInput" />

        <v-file-input v-model="files" label="ファイルを選択" multiple>
          <template v-slot:selection="{ fileNames }">
            <template v-for="fileName in fileNames" :key="fileName">
              <v-chip class="file-chip" color="primary" size="small" label>
                {{ fileName }}
              </v-chip>
            </template>
          </template>
        </v-file-input>

        <v-btn type="submit" color="primary" block class="mt-3">書き込みをする</v-btn>
      </v-form>
    </v-card-text>

    <ModalMedia :dialog="dialog" :mediaItems="dialogMedia" @update:dialog="dialog = $event" />
  </v-card>
</template>

<script setup>
import { ref } from 'vue';
import ModalMedia from '~/components/ModalMedia.vue';

const name = ref('');
const email = ref('');
const comment = ref('');
const files = ref([]);
const previews = ref([]);
const fileInput = ref(null);
const dialog = ref(false);
const dialogMedia = ref([]);

const submitForm = () => {
  console.log('名前:', name.value);
  console.log('E-mail:', email.value);
  console.log('コメント:', comment.value);
  console.log('ファイル:', files.value);
  // ここにフォーム送信のロジックを追加します
};

const triggerFileInput = () => {
  fileInput.value.click();
};

const handleFileChange = event => {
  const selectedFiles = Array.from(event.target.files);
  files.value.push(...selectedFiles);
  previewFiles();
};

const previewFiles = () => {
  previews.value = files.value.map(file => {
    const url = URL.createObjectURL(file);
    console.log(file);
    const preview = { url, type: file.type };
    if (file.type.startsWith('video/')) {
      preview.thumbnail = 'https://via.placeholder.com/300'; // サムネイル画像のURLを指定してください
    }
    return preview;
  });
};

const removeFile = index => {
  files.value.splice(index, 1);
  previews.value.splice(index, 1);
};

const openDialog = () => {
  dialogMedia.value = previews.value;
  dialog.value = true;
};
</script>

<style scoped>
.v-card {
  padding: 16px;
}

.v-card-title {
  font-weight: bold;
}

.file-chip {
  margin-right: 8px;
}
</style>
