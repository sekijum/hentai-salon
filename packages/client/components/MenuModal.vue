<template>
  <v-dialog v-model="isMenuModal" fullscreen hide-overlay transition="dialog-bottom-transition">
    <v-card flat>
      <v-toolbar density="compact">
        <v-toolbar-title>メニュー</v-toolbar-title>

        <v-spacer></v-spacer>

        <v-btn icon @click="isMenuModal = false">
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </v-toolbar>

      <Menu :items="guestMenuItems" />
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import Menu from '~/components/Menu.vue';

const isMenuModal = useState('isMenuModal', () => false);

const config = useRuntimeConfig();
const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { $storage } = useNuxtApp();

const { payload } = nuxtApp;

const guestMenuItems = computed(() => {
  let items: { title: string; clicked: () => any; icon: string }[] = [
    { title: 'スレ一覧', clicked: () => router.push('/threads'), icon: 'mdi-format-list-bulleted' },
    { title: '板一覧', clicked: () => router.push('/boards'), icon: 'mdi-view-list' },
    { title: '設定', clicked: () => router.push('/setting'), icon: 'mdi-cog' },
  ];

  if (!payload.isLoggedIn) {
    items = items.concat([
      { title: 'サインイン', clicked: () => router.push('/signin'), icon: 'mdi-login' },
      { title: 'サインアップ', clicked: () => router.push('/signup'), icon: 'mdi-account-plus' },
    ]);
  } else {
    items = items.concat([
      {
        title: 'サインアウト',
        clicked: () => {
          $storage.removeItem('access_token');
          router.go(0);
        },
        icon: 'mdi-logout',
      },
    ]);
  }

  if (payload.isMember || payload.isAdmin) {
    items = items.concat([
      { title: 'スレ作成', clicked: () => router.push('/threads/new'), icon: 'mdi-pencil' },
      { title: 'ユーザー情報', clicked: () => router.push('/users/me'), icon: 'mdi-account' },
      { title: 'マイスレ', clicked: () => router.push('/users/me/threads'), icon: 'mdi-note' },
      { title: 'マイレス', clicked: () => router.push('/users/me/comments'), icon: 'mdi-comment' },
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

function closeMenuModal() {
  isMenuModal.value = false;
}

watch(
  () => route.path,
  () => closeMenuModal(),
);
</script>
