<template>
  <div>
    <PageTitle title="サインイン" />

    <v-divider />

    <Menu :items="menuItems" />

    <br />

    <Form @submit="submit" :validation-schema="schema" class="mx-2 mb-2" v-slot="{ meta, errors }">
      <div class="field">
        <Field name="email" v-slot="{ field, errorMessage }">
          <v-text-field
            v-model="form.email"
            v-bind="field"
            label="メールアドレス"
            type="email"
            variant="outlined"
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="password" v-slot="{ field, errorMessage }">
          <v-text-field
            v-model="form.password"
            v-bind="field"
            label="パスワード"
            type="password"
            variant="outlined"
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <v-btn type="submit" color="primary" block :disabled="!meta.valid" class="mt-5">サインイン</v-btn>
    </Form>
  </div>
</template>

<script setup lang="ts">
import { Form, Field, ErrorMessage } from 'vee-validate';
import * as yup from 'yup';
import PageTitle from '~/components/PageTitle.vue';
import Menu from '~/components/Menu.vue';
import Storage from '~/utils/storage';

const router = useRouter();
const nuxtApp = useNuxtApp();
const api = nuxtApp.$api;

const form = ref({
  email: '',
  password: '',
});

const menuItems = [
  { title: 'サインイン', navigate: () => router.push('/signin'), icon: 'mdi-login' },
  { title: 'サインアップ', navigate: () => router.push('/signup'), icon: 'mdi-account-plus' },
];

const schema = yup.object({
  email: yup.string().email('有効なメールアドレスを入力してください').required('必須項目です'),
  password: yup.string().min(6, '6文字以上で入力してください').required('必須項目です'),
});

async function submit() {
  try {
    const credentials = { email: form.value.email, password: form.value.password };
    const response = await api.post('/signin', credentials);

    const authHeader = response.headers.authorization;
    const token = authHeader.split(' ')[1];
    Storage.setItem('access_token', token);
    router.push('/');
  } catch (error) {
    console.error('通信中にエラーが発生しました:', error);
  }
}
</script>
