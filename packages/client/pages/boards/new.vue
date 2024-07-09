<template>
  <div class="mx-2">
    <PageTitle title="板作成" />

    <v-divider />

    <Form @submit="submit" :validation-schema="schema" v-slot="{ meta }">
      <div class="field">
        <Field name="title" v-slot="{ field, errorMessage }">
          <v-text-field
            v-model="form.title"
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
            v-model="form.description"
            v-bind="field"
            label="説明"
            variant="outlined"
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <div class="field">
        <v-file-input
          label="サムネイルを選択"
          show-size
          truncate-length="25"
          prepend-icon=""
          variant="outlined"
          dense
          hide-details
          accept="image/*"
          density="compact"
        />
      </div>

      <v-btn type="submit" color="primary" block :disabled="!meta?.valid" class="mt-5">作成</v-btn>
      <p class="note">＊反映には時間が掛かる場合があります＊</p>
    </Form>
  </div>
</template>

<script setup lang="ts">
import PageTitle from '~/components/PageTitle.vue';
import { Form, Field } from 'vee-validate';
import * as yup from 'yup';

const router = useRouter();
const nuxtApp = useNuxtApp();
const { $api } = nuxtApp;

const form = ref({
  title: '',
  description: '',
  thumbnailUrl: null,
});

const schema = yup.object({
  title: yup.string().required('必須項目です'),
});

async function submit() {
  try {
    await $api.post('/admin/boards', form.value);
    alert('板が正常に作成されました。');
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
