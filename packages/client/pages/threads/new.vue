<template>
  <div class="mx-2">
    <PageTitle title="スレ作成" />

    <v-divider />

    <br />

    <Form @submit="submit" :validation-schema="schema" v-slot="{ meta, errors }">
      <div class="field">
        <Field name="title" v-slot="{ field, errorMessage }">
          <v-text-field
            v-bind="field"
            label="タイトル"
            variant="outlined"
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="description" v-slot="{ field, errorMessage }">
          <v-textarea
            v-bind="field"
            label="説明"
            variant="outlined"
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <div class="field">
        <v-combobox
          v-model="form.tags"
          chips
          small-chips
          label="タグ"
          variant="outlined"
          density="compact"
          clearable
          multiple
          :items="['California', 'Colorado', 'Florida', 'Georgia', 'Texas', 'Wyoming']"
        />
      </div>

      <div class="field">
        <Field name="thumbnail" v-slot="{ field, errorMessage }">
          <v-file-input
            v-bind="field"
            label="サムネイルを選択"
            show-size
            truncate-length="25"
            prepend-icon=""
            variant="outlined"
            dense
            hide-details
            accept="image/*"
            :error-messages="errorMessage ? [errorMessage] : []"
            density="compact"
          />
        </Field>
      </div>
    </Form>

    <v-btn type="submit" color="primary" block :disabled="!meta?.valid" class="mt-5">作成</v-btn>
    <p class="note">＊反映には時間が掛かる場合があります＊</p>
  </div>
</template>

<script setup lang="ts">
import { Form, Field, ErrorMessage } from 'vee-validate';
import * as yup from 'yup';
import PageTitle from '~/components/PageTitle.vue';

const router = useRouter();
const nuxtApp = useNuxtApp();
const api = nuxtApp.$api;

const thumbnailFile = new FormData();

const form = ref({
  title: '',
  description: '',
  thumbnailUrl: null,
  tags: [],
});

const schema = yup.object({
  title: yup.string().required('必須項目です'),
  description: yup.string().required('必須項目です'),
});

async function submit() {
  try {
    const response = await api.post('/threads', form.value);

    router.push('/');
  } catch (error) {
    console.error('通信中にエラーが発生しました:', error);
  }
}
</script>

<style scoped>
.note {
  font-size: 12px;
  color: grey;
  text-align: center;
  margin-top: 8px;
}
</style>
