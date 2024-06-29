<template>
  <v-card>
    <v-card-title>設定</v-card-title>

    <v-divider></v-divider>

    <v-data-table
      :headers="headers"
      :items="[
        { name: '板', type: 'header' },
        { name: '閲覧履歴', key: 'board-history-delete' },
        { name: 'スレッド', type: 'header' },
        { name: '閲覧履歴', key: 'thread-history-delete' },
        { name: '最新レス50レスのみ表示', key: 'display-latest-50-replies-only' },
      ]"
      hide-default-footer
      hide-default-header
    >
      <template v-slot:item="{ item }">
        <tr :class="{ highlight: item.type === 'header' }">
          <td>{{ item.name }}</td>
          <td>
            <div v-if="item.key" class="d-flex align-center">
              <v-btn v-if="item.key === 'board-history-delete'" small>削除</v-btn>
              <v-btn v-else-if="item.key === 'thread-history-delete'" small>削除</v-btn>
              <v-switch v-else-if="item.key === 'display-latest-50-replies-only'" flat small></v-switch>
            </div>
          </td>
        </tr>
      </template>
    </v-data-table>

    <v-divider></v-divider>
  </v-card>
</template>

<script setup>
import { ref } from 'vue';
</script>

<style scoped>
.d-flex {
  display: flex;
  align-items: center;
}

.align-center {
  align-items: center;
}

.highlight {
  background-color: #f0f0f0; /* グレー背景色 */
}
</style>
