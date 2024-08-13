<template>
  <div>
    <PageTitle title="問い合わせ" />

    <v-divider />

    <br />

    <Form @submit="submit" :validation-schema="schema" class="mx-2 mb-2" v-slot="{ meta, errors }">
      <div class="field">
        <Field name="email" v-model="form.email" v-slot="{ errors }">
          <v-text-field
            v-model="form.email"
            label="メールアドレス"
            type="email"
            variant="outlined"
            density="compact"
            dense
            single-line
            :error-messages="errors"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="subject" v-model="form.subject" v-slot="{ errors }">
          <v-text-field
            v-model="form.subject"
            label="件名"
            variant="outlined"
            density="compact"
            dense
            single-line
            :error-messages="errors"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="message" v-model="form.message" v-slot="{ errors }">
          <v-textarea
            v-model="form.message"
            label="内容"
            variant="outlined"
            density="compact"
            dense
            single-line
            :error-messages="errors"
          />
        </Field>
      </div>

      <v-btn type="submit" color="primary" block :disabled="!meta.valid">送信</v-btn>
    </Form>
  </div>
</template>

<script setup lang="ts">
import { Form, Field } from 'vee-validate';
import PageTitle from '~/components/PageTitle.vue';
import * as yup from 'yup';

const nuxtApp = useNuxtApp();

const { $api } = nuxtApp;

const form = ref({
  email: '',
  subject: '',
  message: '',
});

const schema = yup.object({
  email: yup.string().email('有効なメールアドレスを入力してください').nullable(),
  subject: yup.string().required('必須項目です'),
  message: yup.string().required('必須項目です'),
});

async function submit(_: typeof form.value, { resetForm }: { resetForm: () => void }) {
  if (confirm('送信しますか？')) {
    try {
      await $api.post('/contact', form.value);
      resetForm();
    } catch (err) {
      alert('通信中にエラーが発生しました');
    }
  }
}

useHead({
  title: '変態サロン | 問い合わせ',
});
</script>
