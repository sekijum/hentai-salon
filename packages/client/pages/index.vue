<template>
  <div>
    <v-alert
      v-if="payload.isLoggedIn"
      :text="`${payload?.user?.name}さん、あなたのアカウントにログインしています。`"
      prominent
      class="small-text text-center"
    />

    <Menu :items="guestMenuItems" />

    <ThreadItem
      filter="history"
      title="スレッド閲覧履歴"
      :clicked="() => router.push({ path: '/threads', query: { filter: 'history' } })"
      :isInfiniteScroll="false"
    />
    <ThreadItem
      filter="popularity"
      title="人気"
      :clicked="() => router.push({ path: '/threads' })"
      :isInfiniteScroll="false"
    />
    <ThreadItem filter="newest" title="新着" :isInfiniteScroll="true" />
  </div>
</template>

<script setup lang="ts">
import Menu from '~/components/Menu.vue';
import ThreadItem from '~/components/thread/ThreadItem.vue';

const config = useRuntimeConfig();
const route = useRoute();
const router = useRouter();
const nuxtApp = useNuxtApp();
const isMenuModal = useState('isMenuModal', () => false);

const { payload } = nuxtApp;

const guestMenuItems = computed(() => {
  let items: { title: string; clicked: () => any; icon: string }[] = [
    { title: 'スレ一覧', clicked: () => router.push('/threads'), icon: 'mdi-format-list-bulleted' },
    { title: '板一覧', clicked: () => router.push('/boards'), icon: 'mdi-view-list' },
    { title: '設定', clicked: () => router.push('/setting'), icon: 'mdi-cog' },
    { title: 'メニュー', clicked: () => (isMenuModal.value = true), icon: 'mdi-menu' },
  ];

  if (!payload.isLoggedIn) {
    items = items.concat([
      { title: 'サインイン', clicked: () => router.push('/signin'), icon: 'mdi-login' },
      { title: 'サインアップ', clicked: () => router.push('/signup'), icon: 'mdi-account-plus' },
    ]);
  }

  if (payload.isMember || payload.isAdmin) {
    items = items.concat([
      { title: 'スレ作成', clicked: () => router.push('/threads/new'), icon: 'mdi-pencil' },
      { title: 'マイページ', clicked: () => router.push('/users/me'), icon: 'mdi-account' },
    ]);
  }

  if (payload.isAdmin) {
    items = items.concat([
      { title: '板作成', clicked: () => router.push('/boards/new'), icon: 'mdi-plus-box' },
      { title: 'adminer', clicked: () => open(config.public.adminerUrl, '_blank'), icon: 'mdi-database' },
    ]);
  }

  return items;
});

useHead({
  title: '変態サロン',
  meta: [
    { name: 'description', content: '変態に特化したサロン。' },
    {
      property: 'og:title',
      content: '変態サロン',
    },
    {
      property: 'og:description',
      content: '変態に特化したサロン。',
    },
    {
      property: 'og:image',
      content: '/hentai-salon-logo/logo_transparent.png',
    },
    {
      property: 'og:url',
      content: location.href,
    },
  ],
});
</script>

<style scoped>
.small-text {
  font-size: 0.75rem;
}
</style>
