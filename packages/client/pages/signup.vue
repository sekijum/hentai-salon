<template>
  <div class="mx-2">
    <PageTitle title="サインアップ" />

    <v-divider></v-divider>

    <v-form @submit.prevent="signup">
      <v-text-field label="名前(コメントの表示名になります)" v-model="name" type="text" required></v-text-field>

      <v-text-field label="メールアドレス" v-model="email" type="email" required></v-text-field>

      <v-text-field label="パスワード" v-model="password" type="password" required></v-text-field>

      <v-file-input
        v-model="avatar"
        label="アバターを選択"
        multiple
        show-size
        truncate-length="25"
        prepend-icon=""
        variant="outlined"
        dense
        hide-details
        accept="image/*"
      />

      <br />

      <v-btn type="submit" color="primary" block>サインアップ</v-btn>
    </v-form>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import PageTitle from '~/components/PageTitle.vue';

const name = ref('');
const email = ref('');
const password = ref('');
const avatar = ref();

const signup = () => {
  // サインアップ処理
  const userSignupRequest = {
    name: name.value,
    email: email.value,
    password: password.value,
    avatar: avatar.value || null,
  };
  console.log(userSignupRequest);
};
</script>

<style scoped></style>
