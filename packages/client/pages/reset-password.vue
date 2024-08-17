<template>
  <div>
    <PageTitle title="新しいパスワードの設定" />

    <v-divider />

    <br />

    <Form @submit="submit" :validation-schema="schema" class="mx-2 mb-2" v-slot="{ meta, errors }">
      <div class="field mb-2">
        <Field name="password" v-model="form.password" v-slot="{ errors }">
          <v-text-field
            v-model="form.password"
            label="パスワード"
            type="password"
            variant="outlined"
            density="compact"
            dense
            single-line
            :error-messages="errors"
          />
        </Field>
      </div>

      <div class="field mb-2">
        <Field name="confirmPassword" v-model="form.confirmPassword" v-slot="{ errors }">
          <v-text-field
            v-model="form.confirmPassword"
            label="パスワード確認"
            type="password"
            variant="outlined"
            density="compact"
            dense
            single-line
            :error-messages="errors"
          />
        </Field>
      </div>

      <v-btn type="submit" color="primary" block :disabled="!meta.valid">設定する</v-btn>
    </Form>
  </div>
</template>

<script setup lang="ts">
import { Form, Field } from 'vee-validate';
import PageTitle from '~/components/PageTitle.vue';
import * as yup from 'yup';

definePageMeta({ middleware: ['unauthentication-only'] });

const nuxtApp = useNuxtApp();
const router = useRouter();
const route = useRoute();
const { $api } = nuxtApp;
const form = ref({ password: '', confirmPassword: '' });
const token = ref(route.query.token);

const schema = yup.object({
  password: yup.string().min(6, '6文字以上で入力してください').required('必須項目です'),
  confirmPassword: yup
    .string()
    .oneOf([yup.ref('password')], 'パスワードが一致しません')
    .required('必須項目です'),
});

async function submit() {
  if (confirm('パスワードを変更しますか？')) {
    try {
      await $api.patch('/reset-password', { password: form.value.password, token: token.value });
      router.push('/signin');
    } catch (err) {
      alert('通信中にエラーが発生しました');
    }
  }
}

onMounted(async () => {
  try {
    await $api.post('/verify-reset-password-token', { token: token.value });
  } catch (err) {
    alert('通信中にエラーが発生しました');
  }
});
</script>
