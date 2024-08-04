<template>
  <div>
    <PageTitle title="サインイン" />

    <v-divider />

    <Menu :items="menuItems" />

    <br />

    <Form @submit="submit" :validation-schema="schema" class="mx-2 mb-2" v-slot="{ meta, errors }">
      <div class="field mb-2">
        <Field name="email" v-model="form.email" v-slot="{ errors }">
          <v-text-field
            v-model="form.email"
            label="メールアドレス"
            type="email"
            variant="outlined"
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <div class="field mb-2">
        <Field name="password" v-model="form.password" v-slot="{ errors }">
          <v-text-field
            v-model="form.password"
            label="パスワード"
            type="password"
            variant="outlined"
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <v-btn type="submit" color="primary" block :disabled="!meta.valid">サインイン</v-btn>
    </Form>

    <nuxt-link to="/forgot-password" class="d-block text-center mt-4">パスワードを忘れた方</nuxt-link>
  </div>
</template>

<script setup lang="ts">
import { Form, Field } from 'vee-validate';
import PageTitle from '~/components/PageTitle.vue';
import Menu from '~/components/Menu.vue';
import * as yup from 'yup';

definePageMeta({ middleware: ['unauthentication-only'] });

const nuxtApp = useNuxtApp();
const router = useRouter();

const { $storage, $api } = nuxtApp;

const form = ref({
  email: '',
  password: '',
});

const menuItems = [
  { title: 'サインイン', clicked: () => router.push('/signin'), icon: 'mdi-login' },
  { title: 'サインアップ', clicked: () => router.push('/signup'), icon: 'mdi-account-plus' },
];

const schema = yup.object({
  email: yup.string().email('有効なメールアドレスを入力してください').required('必須項目です'),
  password: yup.string().min(6, '6文字以上で入力してください').required('必須項目です'),
});

async function submit() {
  try {
    const credentials = { email: form.value.email, password: form.value.password };
    const response = await $api.post('/signin', credentials);

    const authHeader = response.headers.authorization;
    const token = authHeader.split(' ')[1];
    $storage.setItem('access_token', token);
    alert('サインインしました。');
    router.push('/');
  } catch (err) {
    alert('通信中にエラーが発生しました');
  }
}

useHead({
  title: '変態サロン | サインイン',
});
</script>
