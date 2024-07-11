<template>
  <div>
    <p class="form-title">{{ title ?? '書き込み' }}</p>

    <Form @submit="submit" :validation-schema="schema" v-slot="{ meta }">
      <div class="field">
        <Field name="guestName" v-slot="{ field, errorMessage }">
          <v-text-field
            v-model="form.guestName"
            v-bind="field"
            label="名前(省略可)"
            variant="outlined"
            hide-details
            counter
            single-line
            clearable
            dense
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="content" v-slot="{ field, errorMessage }">
          <v-textarea
            v-model="form.content"
            v-bind="field"
            label="コメント"
            rows="3"
            variant="outlined"
            hide-details
            counter
            dense
            single-line
            clearable
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <div class="field">
        <v-file-input
          label="ファイルを選択"
          show-size
          truncate-length="25"
          prepend-icon=""
          variant="outlined"
          dense
          hide-details
          multiple
          accept="image/*"
          density="compact"
          chips
        />
      </div>

      <v-btn class="clear-button" block @click="clearForm">クリア</v-btn>
      <v-btn type="submit" class="submit-button" block :disabled="!meta?.valid">書き込みをする</v-btn>
      <p class="note">＊書き込み反映には時間が掛かる場合があります＊</p>
    </Form>
  </div>
</template>

<script setup lang="ts">
import { Form, Field, ErrorMessage } from 'vee-validate';
import * as yup from 'yup';

const props = defineProps<{ title?: string; parentCommentId?: number }>();

const nuxtApp = useNuxtApp();
const router = useRouter();
const route = useRoute();
const { $api } = nuxtApp;

const form = ref({
  guestName: '',
  content: '',
});

const schema = yup.object({
  guestName: yup.string().optional(),
  content: yup.string().required('コメントは必須項目です'),
  files: yup.array().of(yup.mixed()).optional(),
});

const clearForm = () => {
  if (confirm('クリアしますか？')) {
    form.value.guestName = '';
    form.value.content = '';
  }
};

const handleFileChange = event => {};

async function submit() {
  if (confirm('本当に書き込みますか？')) {
    try {
      const threadId = parseInt(route.params.id.toString(), 10);
      if (props.parentCommentId) {
        await $api.post(`/threads/${threadId}/comments/${props.parentCommentId}/reply`, {
          ...form.value,
          ...{ parentCommentId: props.parentCommentId },
        });
        alert('返信しました。');
      } else {
        await $api.post(`/threads/${threadId}/comments/`, form.value);
        alert('書き込みました。');
      }
      router.go(0);
    } catch (error) {
      console.error('通信中にエラーが発生しました:', error);
    }
  }
}
</script>

<style scoped>
.form-title {
  font-size: 1.2em;
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
  margin-top: 10px;
  margin-bottom: 10px;
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
