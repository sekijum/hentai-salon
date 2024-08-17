<template>
  <div>
    <PageTitle title="パスワードを忘れた方" />

    <v-divider />

    <br />

    <p class="mb-4 text-muted">ご登録されたメールアドレスにパスワード再設定のご案内が送信されます。</p>

    <Form @submit="submit" :validation-schema="schema" class="mx-2 mb-2" v-slot="{ meta, errors }">
      <div class="field mb-2">
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

      <v-btn type="submit" color="primary" block :disabled="!meta.valid">パスワードをリセットする</v-btn>
    </Form>
  </div>
</template>

<script setup lang="ts">
import { Form, Field } from 'vee-validate';
import PageTitle from '~/components/PageTitle.vue';
import * as yup from 'yup';

definePageMeta({ middleware: ['unauthentication-only'] });

const nuxtApp = useNuxtApp();
const { $api } = nuxtApp;
const form = ref({ email: '' });

const schema = yup.object({
  email: yup.string().email('有効なメールアドレスを入力してください').required('必須項目です'),
});

async function submit() {
  try {
    await $api.post('/forgot-password', { email: form.value.email });
    alert(`メールを確認してください${form.value.email}にパスワード再設定のメールを送信しました`);
  } catch (err) {
    alert('通信中にエラーが発生しました');
  }
}

useHead({
  title: '変態サロン | パスワードを忘れた方',
});
</script>
