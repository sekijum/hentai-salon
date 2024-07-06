<template>
  <div>
    <PageTitle title="サインアップ" />

    <v-divider />

    <Menu :items="menuItems" />

    <br />

    <Form @submit="submit" :validation-schema="schema" class="mx-2 mb-2" v-slot="{ meta, errors }">
      <div class="field">
        <Field name="name" v-slot="{ field, errorMessage }">
          <v-text-field
            v-model="form.name"
            v-bind="field"
            label="名前(コメントの表示名になります)"
            type="text"
            variant="outlined"
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

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

      <div class="field">
        <Field name="confirmPassword" v-slot="{ field, errorMessage }">
          <v-text-field
            v-bind="field"
            label="パスワード確認"
            type="text"
            variant="outlined"
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <div class="field">
        <v-file-input
          v-bind="field"
          show-size
          truncate-length="25"
          prepend-icon=""
          label="アバターを選択"
          variant="outlined"
          density="compact"
          hide-details
          accept="image/*"
          :error-messages="errorMessage ? [errorMessage] : []"
        />
      </div>

      <v-btn type="submit" color="primary" block :disabled="!meta.valid" class="mt-5">サインアップ</v-btn>
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

const avatarFile = new FormData();

const form = ref({
  name: '',
  email: '',
  password: '',
  avatarUrl: null,
});

const menuItems = [
  { title: 'サインイン', navigate: () => router.push('/signin'), icon: 'mdi-login' },
  { title: 'サインアップ', navigate: () => router.push('/signup'), icon: 'mdi-account-plus' },
];

const schema = yup.object({
  name: yup.string().required('必須項目です'),
  email: yup.string().email('有効なメールアドレスを入力してください').required('必須項目です'),
  password: yup.string().min(6, '6文字以上で入力してください').required('必須項目です'),
  confirmPassword: yup
    .string()
    .oneOf([yup.ref('password'), null], 'パスワードが一致しません')
    .required('必須項目です'),
  avatar: yup.mixed().nullable(),
});

async function submit() {
  try {
    const response = await api.post('/signup', form.value);

    router.push('/');
  } catch (error) {
    console.error('通信中にエラーが発生しました:', error);
  }
}
</script>
