<template>
  <v-container fluid>
    <v-breadcrumbs :items="['HOME', 'ユーザー']"></v-breadcrumbs>
    <v-data-table-server
      v-model:items-per-page="itemsPerPage"
      :headers="headers"
      :items="serverItems"
      :items-length="totalItems"
      :loading="loading"
      :hover="true"
      item-value="id"
      @update:options="loadItems"
      items-per-page-text="表示行数"
      density="compact"
    >
      <template #top>
        <v-toolbar flat>
          <v-toolbar-title>ユーザー一覧</v-toolbar-title>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>
          <v-dialog v-model="dialog" max-width="500px">
            <template #activator="{ props }">
              <v-btn class="mb-2" color="primary" dark v-bind="props">作成</v-btn>
            </template>
            <v-card>
              <v-card-title>
                <span class="text-h5">{{ formTitle() }}</span>
              </v-card-title>
              <v-card-text>
                <v-container>
                  <v-row>
                    <v-col cols="12">
                      <v-text-field v-model="editedItem.name" label="名前"></v-text-field>
                    </v-col>
                    <v-col cols="12">
                      <v-text-field v-model="editedItem.email" label="メール"></v-text-field>
                    </v-col>
                    <v-col cols="12">
                      <v-text-field v-model="editedItem.status" label="ステータス"></v-text-field>
                    </v-col>
                    <v-col cols="12">
                      <v-text-field v-model="editedItem.role" label="権限"></v-text-field>
                    </v-col>
                  </v-row>
                </v-container>
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue-darken-1" text @click="close">キャンセル</v-btn>
                <v-btn color="blue-darken-1" text @click="save">保存</v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
          <v-dialog v-model="dialogDelete" max-width="500px">
            <v-card>
              <v-card-title class="text-h5">本当に削除しますか？</v-card-title>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue-darken-1" text @click="closeDelete">キャンセル</v-btn>
                <v-btn color="blue-darken-1" text @click="deleteItemConfirm">OK</v-btn>
                <v-spacer></v-spacer>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-toolbar>
      </template>
      <template #item.createdAt="{ item }">
        {{ $formatDate(item.createdAt) }}
      </template>
      <template #item.actions="{ item }">
        <v-icon class="me-2" size="small" @click="editItem(item)">mdi-pencil</v-icon>
        <v-icon size="small" @click="deleteItem(item)">mdi-delete</v-icon>
      </template>
      <template #no-data>
        <v-btn color="primary" @click="loadItems">リセット</v-btn>
      </template>
    </v-data-table-server>
  </v-container>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter, useNuxtApp } from '#app';
import type { ICollection } from '~/types/collection';

definePageMeta({
  layout: 'admin',
  middleware: ['admin-access-only'],
});

interface IUser {
  id: number;
  name: string;
  role: number;
  roleLabel: string;
  status: number;
  statusLabel: string;
  email: string;
  profileLink: string;
  createdAt: string;
  updatedAt: string;
}

const router = useRouter();
const nuxtApp = useNuxtApp();
const { $api, $formatDate } = nuxtApp;

const dialog = ref(false);
const dialogDelete = ref(false);
const editedIndex = ref(-1);
const editedItem = ref<IUser>({
  id: 0,
  name: '',
  role: 0,
  roleLabel: '',
  status: 0,
  statusLabel: '',
  email: '',
  profileLink: '',
  createdAt: '',
  updatedAt: '',
});
const defaultItem = {
  id: 0,
  name: '',
  role: 0,
  roleLabel: '',
  status: 0,
  statusLabel: '',
  email: '',
  profileLink: '',
  createdAt: '',
  updatedAt: '',
};

const headers = [
  { title: 'ID', align: 'start', sortable: true, key: 'id' },
  { title: '名前', key: 'name' },
  { title: 'メール', key: 'email' },
  { title: '権限', key: 'roleLabel' },
  { title: 'ステータス', key: 'statusLabel' },
  { title: '登録日', key: 'createdAt' },
  { title: '操作', key: 'actions', sortable: false },
];

const itemsPerPage = ref(10);
const totalItems = ref(0);
const serverItems = ref<IUser[]>([]);
const loading = ref(false);

function formTitle() {
  return editedIndex.value === -1 ? '新規作成' : '編集';
}

function editItem(item: IUser) {
  editedIndex.value = serverItems.value.indexOf(item);
  editedItem.value = { ...item };
  dialog.value = true;
}

function deleteItem(item: IUser) {
  editedIndex.value = serverItems.value.indexOf(item);
  editedItem.value = { ...item };
  dialogDelete.value = true;
}

function deleteItemConfirm() {
  serverItems.value.splice(editedIndex.value, 1);
  closeDelete();
}

function close() {
  dialog.value = false;
  editedItem.value = { ...defaultItem };
  editedIndex.value = -1;
}

function closeDelete() {
  dialogDelete.value = false;
  editedItem.value = { ...defaultItem };
  editedIndex.value = -1;
}

function save() {
  if (editedIndex.value > -1) {
    Object.assign(serverItems.value[editedIndex.value], editedItem.value);
  } else {
    serverItems.value.push(editedItem.value);
  }
  close();
}

async function loadItems(params: { page: number; itemsPerPage: number; sortBy: [] }) {
  loading.value = true;
  console.log({
    page: params.page,
    limit: params.itemsPerPage,
    sort: params.sortBy.length ? params.sortBy[0].key : null,
    order: params.sortBy.length ? params.sortBy[0].order : null,
    offset: (params.page - 1) * params.itemsPerPage,
  });
  const response = await $api.get<ICollection<IUser>>('/admin/users', {
    params: {
      offset: (params.page - 1) * params.itemsPerPage,
      limit: params.itemsPerPage,
      sort: params.sortBy.length ? params.sortBy[0].key : null,
      order: params.sortBy.length ? params.sortBy[0].order : null,
    },
  });
  serverItems.value = response.data.data;
  totalItems.value = response.data.totalCount;
  loading.value = false;
}

onMounted(() => loadItems({ page: 1, itemsPerPage: itemsPerPage.value, sortBy: [] }));
</script>

<style scoped>
/* 必要に応じてスタイルを追加 */
</style>
