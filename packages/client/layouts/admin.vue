<template>
  <v-app>
    <v-navigation-drawer permanent expand-on-hover rail>
      <v-list>
        <v-list-item prepend-avatar="/no-image.jpg" :subtitle="payload.user?.email" :title="payload.user?.name" />
      </v-list>

      <v-divider />

      <v-list density="compact" nav>
        <v-list-item
          prepend-icon="mdi-account-group"
          title="ユーザー管理"
          @click="navigate('/admin/users')"
          value="member-management"
        />
        <v-list-item
          prepend-icon="mdi-view-dashboard"
          title="板管理"
          @click="navigate('/admin/boards')"
          value="board-management"
        />
        <v-list-item
          prepend-icon="mdi-text-box-multiple-outline"
          title="スレ管理"
          @click="navigate('/admin/threads')"
          value="thread-management"
        />
        <v-list-item
          prepend-icon="mdi-email"
          title="問い合わせ"
          @click="navigate('/admin/contacts')"
          value="inquiry-management"
        />
        <v-list-item
          prepend-icon="mdi-arrow-right-bold"
          title="クライアント画面へ"
          @click="navigate('/')"
          value="client-view"
        />
        <v-list-item prepend-icon="mdi-database" title="adminer" @click="openAdminer" value="client-view" />
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <NuxtPage />
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
const router = useRouter();
const config = useRuntimeConfig();

function navigate(to: string) {
  router.push(to);
}

const nuxtApp = useNuxtApp();
const { payload } = nuxtApp;

function openAdminer() {
  open(config.public.adminerUrl, '_blank');
}
</script>

<style scoped></style>
