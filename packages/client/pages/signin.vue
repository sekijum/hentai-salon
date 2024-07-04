<template>
  <div>
    <PageTitle title="サインイン" />
    <v-divider></v-divider>

    <Menu :items="menuItems" />

    <br />

    <v-form @submit.prevent="signin" ref="formRef" class="mx-2">
      <v-text-field
        label="メールアドレス"
        v-model="form.email"
        type="email"
        :rules="[rules.required, rules.email]"
        required
        density="compact"
        variant="outlined"
      ></v-text-field>

      <v-text-field
        label="パスワード"
        v-model="form.password"
        type="password"
        :rules="[rules.required, rules.min(6)]"
        required
        density="compact"
        variant="outlined"
      ></v-text-field>

      <v-btn type="submit" color="primary" block>サインイン</v-btn>
    </v-form>
  </div>
</template>

<script setup lang="ts">
import PageTitle from '~/components/PageTitle.vue';
import Menu from '~/components/Menu.vue';
import Storage from '~/utils/storage';
import api from '~/utils/api';

const router = useRouter();
const menuItems = [
  { title: 'サインイン', navigate: () => router.push('/signin'), icon: 'mdi-login' },
  { title: 'サインアップ', navigate: () => router.push('/signup'), icon: 'mdi-account-plus' },
];

const form = ref({
  email: '',
  password: '',
});
const formRef = ref();

const rules = {
  required: (value: string) => !!value || '必須項目です',
  email: (value: string) => /.+@.+\..+/.test(value) || '有効なメールアドレスを入力してください',
  min: (length: number) => (value: string) => value.length >= length || `${length}文字以上で入力してください`,
};

async function signin() {
  if (formRef.value.validate()) {
    try {
      const credentials = { email: form.value.email, password: form.value.password };
      const response = await api.post('/signin', credentials);

      const token = response.headers['authorization'].split(' ')[1];
      console.log(token);

      Storage.setItem('access_token', token);

      router.push('/');
    } catch (error) {
      console.error('ログイン中にエラーが発生しました:', error);
    }
  }
}
</script>

<style scoped></style>
