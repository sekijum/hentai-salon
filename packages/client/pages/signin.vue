<template>
  <div class="mx-2">
    <PageTitle title="サインイン" />

    <v-divider></v-divider>

    <v-form @submit.prevent="login">
      <v-text-field label="Email Address" v-model="email" type="email" required></v-text-field>

      <v-text-field label="Password" v-model="password" type="password" required></v-text-field>

      <v-btn type="submit" color="primary" block>サインイン</v-btn>
    </v-form>
  </div>
</template>

<script setup>
import PageTitle from '~/components/PageTitle.vue';

const email = ref('');
const password = ref('');

const login = () => {
  // ログイン処理
  console.log({ email: email.value, password: password.value });
};
</script>

<style scoped></style>
