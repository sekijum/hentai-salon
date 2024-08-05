<template></template>

<script setup lang="ts">
definePageMeta({ middleware: ['unauthentication-only'] });

const nuxtApp = useNuxtApp();
const router = useRouter();
const route = useRoute();

const { $api, $storage } = nuxtApp;

onMounted(async () => {
  if (!route.query.token) {
    router.push('/');
  }
  try {
    const response = await $api.post('/verify-email-token', { token: route.query.token });
    const authHeader = response.headers.authorization;
    const accessToken = authHeader.split(' ')[1];
    $storage.setItem('access_token', accessToken);
    alert('サインインしました。');
    router.push('/');
  } catch (err) {
    alert('通信中にエラーが発生しました');
  }
});
</script>
