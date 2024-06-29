<template>
  <v-sheet>
    <v-row align="center" justify="center">
      <v-col cols="4">
        <v-btn block text class="menu-button">
          <v-icon class="menu-icon">mdi-fire</v-icon>
          人気
        </v-btn>
      </v-col>
      <v-col cols="4">
        <v-btn block text class="menu-button">
          <v-icon class="menu-icon">mdi-update</v-icon>
          閲覧履歴
        </v-btn>
      </v-col>
      <v-col cols="4">
        <v-btn block text class="menu-button">
          <v-icon class="menu-icon">mdi-earth</v-icon>
          ログイン
        </v-btn>
      </v-col>
      <v-col cols="4">
        <v-btn block text class="menu-button">
          <v-icon class="menu-icon">mdi-new-box</v-icon>
          新着
        </v-btn>
      </v-col>
      <v-col cols="4">
        <v-btn block text class="menu-button">
          <v-icon class="menu-icon">mdi-format-list-bulleted</v-icon>
          板一覧
        </v-btn>
      </v-col>
      <v-col cols="4">
        <v-btn block text class="menu-button">
          <v-icon class="menu-icon">mdi-cog</v-icon>
          設定
        </v-btn>
      </v-col>
    </v-row>
  </v-sheet>
</template>

<script setup></script>

<style scoped>
.menu-button {
  align-items: center;
  border: none;
  box-shadow: none;
}
</style>
