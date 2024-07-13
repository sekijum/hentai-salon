<template>
  <v-app>
    <v-navigation-drawer permanent expand-on-hover rail>
      <v-list>
        <v-list-item
          prepend-avatar="https://randomuser.me/api/portraits/women/85.jpg"
          subtitle="sandra_a88@gmail.com"
          title="Sandra Adams"
        ></v-list-item>
      </v-list>

      <v-divider />

      <v-list density="compact" nav>
        <v-list-item
          prepend-icon="mdi-account-group"
          title="ユーザー管理"
          @click="navigate('/admin/users')"
          value="member-management"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-tag"
          title="タグ管理"
          @click="navigate('/admin/tags')"
          value="tag-management"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-forum"
          title="スレッド管理"
          @click="navigate('/admin/threads')"
          value="thread-management"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-view-dashboard"
          title="板管理"
          @click="navigate('/admin/boards')"
          value="board-management"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-arrow-right-bold"
          title="クライアント画面へ"
          @click="navigate('/')"
          value="board-management"
        ></v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <NuxtPage />
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
const router = useRouter();
function navigate(to: string) {
  router.push(to);
}
</script>

<style scoped>
/* 必要に応じてスタイルを追加 */
</style>
