<template>
  <v-app>
    <v-navigation-drawer permanent expand-on-hover rail>
      <v-list>
        <v-list-item prepend-avatar="/no-image.jpg" :subtitle="payload.user?.email" :title="payload.user?.name" />
      </v-list>

      <v-divider />

      <v-list density="compact" nav>
        <v-list-item
          prepend-icon="mdi-view-dashboard"
          title="ダッシュボード"
          @click="navigate('/admin')"
          value="dashboard"
        />
        <v-list-item
          prepend-icon="mdi-account"
          title="ユーザー管理"
          @click="navigate('/admin/users')"
          value="user-management"
        />
        <v-list-item
          prepend-icon="mdi-view-list"
          title="板管理"
          @click="navigate('/admin/boards')"
          value="board-management"
        />
        <v-list-item
          prepend-icon="mdi-forum-outline"
          title="スレ管理"
          @click="navigate('/admin/threads')"
          value="thread-management"
        />
        <v-list-item
          prepend-icon="mdi-email-outline"
          title="問い合わせ"
          @click="navigate('/admin/contacts')"
          value="inquiry-management"
        />
        <v-list-item
          prepend-icon="mdi-tag-outline"
          title="タグ"
          @click="navigate('/admin/tags')"
          value="tag-management"
        />
        <v-list-item
          prepend-icon="mdi-arrow-right-bold"
          title="クライアント画面へ"
          @click="navigate('/')"
          value="client-view"
        />
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
</script>

<style scoped></style>
