<template>
  <v-row justify="center">
    <v-dialog v-model="isMenuModal" fullscreen hide-overlay transition="dialog-bottom-transition">
      <v-card>
        <v-toolbar dark>
          <v-btn icon dark @click="isMenuModal = false">
            <v-icon>mdi-close</v-icon>
          </v-btn>
          <v-toolbar-title>メニュー</v-toolbar-title>
          <v-spacer></v-spacer>
        </v-toolbar>
        <Menu :items="menuItems" />
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script setup lang="ts">
import Menu from '~/components/Menu.vue';

const isMenuModal = useState('isMenuModal', () => false);

const router = useRouter();
const route = useRoute();

const menuItems = [
  { title: 'お知らせ', clicked: () => router.push('/'), icon: 'mdi-update' },
  { title: 'スレ一覧', clicked: () => router.push('/threads'), icon: 'mdi-new-box' },
  { title: '板一覧', clicked: () => router.push('/boards'), icon: 'mdi-format-list-bulleted' },
  { title: '設定', clicked: () => router.push('/setting'), icon: 'mdi-cog' },
  { title: 'サインイン', clicked: () => router.push('/signin'), icon: 'mdi-login' },
  { title: 'サインアップ', clicked: () => router.push('/signup'), icon: 'mdi-account-plus' },
  { title: 'サインアウト', clicked: () => router.push('/signup'), icon: 'mdi-logout' },
  { title: 'スレ作成', clicked: () => router.push('/threads/new'), icon: 'mdi-forum' },
  { title: '板作成', clicked: () => router.push('/boards/new'), icon: 'mdi-forum' },
  { title: '管理画面', clicked: () => router.push('/admin'), icon: 'mdi-forum' },
  { title: 'いいね', clicked: () => router.push('/admin'), icon: 'mdi-forum' },
  { title: 'ユーザー', clicked: () => router.push('/admin'), icon: 'mdi-forum' },
];

function closeMenuModal() {
  isMenuModal.value = false;
}

watch(
  () => route.path,
  () => closeMenuModal(),
);
</script>
