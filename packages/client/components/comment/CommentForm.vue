<template>
  <div>
    <p class="form-title">{{ formTitle }}</p>
    <v-form @submit.prevent="submitForm" class="form">
      <v-row dense class="no-gutters">
        <v-col cols="12">
          <v-text-field v-model="name" label="名前(省略可)" variant="outlined" dense hide-details></v-text-field>
        </v-col>
      </v-row>
      <v-textarea v-model="comment" label="コメント" rows="4" variant="outlined" dense hide-details></v-textarea>
      <input type="file" multiple @change="handleFileChange" style="display: none" ref="fileInput" />

      <v-file-input
        v-model="files"
        label="ファイルを選択"
        multiple
        show-size
        truncate-length="25"
        prepend-icon=""
        variant="outlined"
        dense
        hide-details
      >
        <template v-slot:selection="{ fileNames }">
          <v-chip v-for="fileName in fileNames" :key="fileName" class="me-2" color="primary" size="small" label>
            {{ fileName }}
          </v-chip>
        </template>
      </v-file-input>

      <v-btn class="clear-button" block @click="clearForm">クリア</v-btn>
      <v-btn type="submit" class="submit-button" block>書き込みをする</v-btn>
      <p class="note">＊書き込み反映には時間が掛かる場合があります＊</p>
    </v-form>
  </div>
</template>

<script setup lang="ts">
import { ref, defineEmits, defineProps } from 'vue';

const emit = defineEmits(['submit', 'clear']);

const props = defineProps({
  formTitle: {
    type: String,
    default: '書き込み',
  },
});

const name = ref('');
const comment = ref('');
const files = ref([]);
const fileInput = ref(null);

const submitForm = () => {
  console.log('名前:', name.value);
  console.log('コメント:', comment.value);
  console.log('ファイル:', files.value);
  emit('submit');
};

const triggerFileInput = () => {
  fileInput.value.click();
};

const clearForm = () => {
  name.value = '';
  comment.value = '';
  files.value = [];
  emit('clear');
};

const handleFileChange = event => {
  const selectedFiles = Array.from(event.target.files);
  files.value.push(...selectedFiles);
};
</script>

<style scoped>
.form-title {
  font-size: 1.2em;
}

.v-row.no-gutters {
  margin-left: 0;
  margin-right: 0;
}

.v-row.no-gutters > .v-col {
  padding-left: 0;
  padding-right: 0;
}

.v-text-field,
.v-textarea,
.v-file-input {
  margin-bottom: 0px !important;
}

.v-text-field input,
.v-textarea textarea {
  font-size: 12px;
}

.clear-button,
.submit-button {
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-bottom: 0px;
}

.clear-button {
  background-color: #f0f0f0;
  color: black;
  margin-bottom: 8px;
}

.submit-button {
  background-color: #007bff;
  color: white;
}

.submit-button:hover {
  background-color: #0056b3;
}

.clear-button:hover {
  background-color: #e0e0e0;
}

.note {
  font-size: 12px;
  color: grey;
  text-align: center;
  margin-top: 8px;
}
</style>
