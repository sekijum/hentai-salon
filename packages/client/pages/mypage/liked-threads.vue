<template>
  <div>
    <PageTitle title="お気に入りスレ" />

    <Menu :items="menuItems" />

    <ThreadItem filter="liked" :isInfiniteScroll="true" />
  </div>
</template>

<script setup lang="ts">
import PageTitle from '~/components/PageTitle.vue';
import ThreadItem from '~/components/thread/ThreadItem.vue';

definePageMeta({ middleware: ['logged-in-access-only'] });

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();

const menuItems = [
  { title: 'マイスレ', clicked: () => router.push('/mypage/threads'), icon: 'mdi-file-document-multiple-outline' },
  { title: 'ユーザー情報', clicked: () => router.push('/mypage'), icon: 'mdi-account-cog-outline' },
  { title: 'マイレス', clicked: () => router.push('/mypage/comments'), icon: 'mdi-message-text-outline' },
  {
    title: 'お気に入りスレ',
    clicked: () => router.push('/mypage/liked-threads'),
    icon: 'mdi-star-box-multiple-outline',
  },
  { title: 'お気に入りレス', clicked: () => router.push('/mypage/liked-comments'), icon: 'mdi-message-star-outline' },
];
</script>
