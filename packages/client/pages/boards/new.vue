<template>
  <div class="mx-2">
    <PageTitle title="板作成" />

    <v-divider />

    <Form @submit="submit" :validation-schema="schema" v-slot="{ meta }">
      <div class="field">
        <Field name="title" v-model="form.title" v-slot="{ errors }">
          <v-text-field
            v-model="form.title"
            label="タイトル"
            variant="outlined"
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="description" v-model="form.description" v-slot="{ errors }">
          <v-textarea
            v-model="form.description"
            label="説明"
            variant="outlined"
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <v-btn type="submit" color="primary" block :disabled="!meta?.valid">作成</v-btn>
      <p class="note">＊反映には時間が掛かる場合があります＊</p>
    </Form>
  </div>
</template>

<script setup lang="ts">
import PageTitle from '~/components/PageTitle.vue';
import { Form, Field } from 'vee-validate';
import * as yup from 'yup';

definePageMeta({ middleware: ['admin-access-only'] });

const router = useRouter();
const nuxtApp = useNuxtApp();

const { $api } = nuxtApp;

const form = ref({
  title: '',
  description: '',
});

const schema = yup.object({
  title: yup.string().required('必須項目です'),
});

async function submit() {
  try {
    if (confirm('板を作成しますか？')) {
      await $api.post('/admin/boards', form.value);
      router.push('/');
    }
  } catch (error) {
    alert('通信中にエラーが発生しました');
  }
}

useHead({
  title: '変態サロン | 板作成',
});
</script>

<style scoped>
.note {
  font-size: 12px;
  color: grey;
  text-align: center;
  margin-top: 8px;
}
</style>
