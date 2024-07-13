<template>
  <div>
    <PageTitle title="設定" />

    <v-divider />

    <div v-for="item in menuItems" :key="item.name" :class="{ highlight: item.type === 'header' }" class="bordered-row">
      <v-row class="align-center">
        <template v-if="item.type === 'header'">
          <div class="header-text">{{ item.name }}</div>
        </template>
        <template v-else>
          <v-col cols="6" class="d-flex justify-center align-center item-name">{{ item.name }}</v-col>
          <v-col cols="6" class="d-flex justify-center align-center">
            <div v-if="item.key" class="d-flex justify-center align-center">
              <v-btn
                v-if="item.key === 'thread-history-delete'"
                small
                class="spaced-button"
                @click="clearThreadViewHistory"
                >削除</v-btn
              >
              <v-switch
                v-else-if="item.key === 'comment-sort-order'"
                v-model="commentSortOrder"
                flat
                small
                class="centered-switch"
                @change="setCommentSortOrder(commentSortOrder ? 'desc' : 'asc')"
              ></v-switch>
              <v-switch
                v-else-if="item.key === 'comment-limit'"
                v-model="commentLimit"
                flat
                small
                class="centered-switch"
                @change="setCommentLimit(commentLimit)"
              ></v-switch>
              <v-switch
                v-else-if="item.key === 'theme'"
                v-model="theme"
                flat
                small
                class="centered-switch"
                @change="
                  () => {
                    setTheme(theme ? 'light' : 'dark'), router.go(0);
                  }
                "
              ></v-switch>
            </div>
          </v-col>
        </template>
      </v-row>
    </div>
  </div>
</template>

<script setup lang="ts">
import PageTitle from '~/components/PageTitle.vue';

const router = useRouter();

const {
  getCommentLimit,
  getTheme,
  getCommentSortOrder,
  setCommentLimit,
  setTheme,
  setCommentSortOrder,
  clearThreadViewHistory,
} = useStorage();

const theme = ref(getTheme() === 'light' ? true : false);
const commentSortOrder = ref(getCommentSortOrder() === 'desc' ? true : false);
const commentLimit = ref(getCommentLimit() ? true : false);

const menuItems = [
  { name: '表示', type: 'header' },
  { name: 'ダークモード', key: 'theme' },
  { name: 'スレッド', type: 'header' },
  { name: '閲覧履歴', key: 'thread-history-delete' },
  { name: '最新レス50レスのみ表示', key: 'comment-limit' },
  { name: 'コメント', type: 'header' },
  { name: '昇順/降順', key: 'comment-sort-order' },
];
</script>

<style scoped>
.highlight {
  background-color: #f0f0f0;
}
.bordered-row {
  border-top: 1px solid #ddd;
  border-bottom: 1px solid #ddd;
}
.align-center {
  align-items: center;
}
.centered-switch {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}
.spaced-button {
  margin-top: 8px;
  margin-bottom: 8px;
}
.text-center {
  text-align: center;
}
.item-name {
  display: flex;
  align-items: center;
  justify-content: center;
}
.header-text {
  padding: 16px;
  width: 100%;
  text-align: center;
}
</style>
