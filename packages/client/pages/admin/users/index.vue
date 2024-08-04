<template>
  <v-container fluid>
    <v-breadcrumbs
      :items="[
        {
          title: 'ダッシュボード',
          disabled: false,
          href: '/admin',
        },
        {
          title: 'ユーザー一覧',
          disabled: false,
          href: '/admin/users',
        },
      ]"
    />

    <v-text-field
      v-model="search"
      density="compact"
      label="検索"
      prepend-inner-icon="mdi-magnify"
      variant="solo-filled"
      flat
      hide-details
      single-line
    />

    <v-divider class="border-opacity-0" :thickness="20" />

    <v-expansion-panels>
      <v-expansion-panel title="検索の使い方の説明">
        <v-expansion-panel-text>
          <ul>
            <li><strong>id:</strong> 特定のユーザーIDで検索します。例: <code>id:123</code></li>
            <li>
              <strong>role:</strong> 特定の権限を持つユーザーを検索します。以下の例を参照してください:
              <ul>
                <li>会員を検索: <code>role:0</code></li>
                <li>管理者を検索: <code>role:1</code></li>
              </ul>
            </li>
            <li>
              <strong>status:</strong> 特定のステータスを持つユーザーを検索します。以下の例を参照してください:
              <ul>
                <li>有効なユーザーを検索: <code>status:0</code></li>
                <li>退会済のユーザーを検索: <code>status:1</code></li>
                <li>凍結されたユーザーを検索: <code>status:2</code></li>
                <li>無効化されたユーザーを検索: <code>status:3</code></li>
              </ul>
            </li>
            <li>また、名前やメールアドレスを部分一致で検索することも可能です。</li>
          </ul>
        </v-expansion-panel-text>
      </v-expansion-panel>
    </v-expansion-panels>

    <v-divider class="border-opacity-0" :thickness="20" />

    <v-data-table-server
      v-model:sort-by="sortBy"
      v-model:items-per-page="itemsPerPage"
      :headers="headers"
      :items="users"
      :items-length="totalCount"
      :hover="true"
      item-value="id"
      @update:options="fetchUsers"
      items-per-page-text="表示行数"
      density="compact"
      must-sort
      no-data-text="検索結果は0件です"
      v-model:search="search"
    >
      <template #item="{ item }">
        <tr>
          <td>{{ item.id }}</td>
          <td>{{ item.name }}</td>
          <td>{{ item.email }}</td>
          <td>{{ item.roleLabel }}</td>
          <td>
            <v-select
              v-model="item.status"
              :items="statusList"
              :item-props="statusProps"
              label="ステータス"
              variant="outlined"
              density="compact"
              hide-details
              counter
              dense
              single-line
              :item-title="'text'"
              :item-value="'value'"
              @update:modelValue="updateStatus(item.id, item.status)"
            />
          </td>
          <td>{{ $formatDate(item.createdAt) }}</td>
          <td>
            <v-icon class="me-2" size="small" @click="router.push(`/admin/users/${item.id}`)">mdi-pencil</v-icon>
          </td>
        </tr>
      </template>
    </v-data-table-server>
  </v-container>
</template>

<script setup lang="ts">
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

type SortBy = {
  key: string;
  order: 'asc' | 'desc';
};

const initSortBy: SortBy = {
  key: 'id',
  order: 'desc',
};
const search = ref('');
const sortBy = ref<SortBy[]>([initSortBy]);
const itemsPerPage = ref(20);
const totalCount = ref(0);
const users = ref<IUser[]>([]);

const headers = [
  { title: 'ID', align: 'start', sortable: true, key: 'id' },
  { title: '名前', key: 'name' },
  { title: 'メール', key: 'email' },
  { title: '権限', key: 'roleLabel' },
  { title: 'ステータス', key: 'statusLabel' },
  { title: '登録日', key: 'createdAt' },
  { title: '操作', key: 'actions', sortable: false },
];

const statusList = [
  { text: '有効', value: 0, subtitle: '現在アクティブなユーザー' },
  { text: '退会済', value: 1, subtitle: 'ユーザーが自主的に退会した状態' },
  { text: '凍結', value: 2, subtitle: 'アカウントが一時的に使用できない状態' },
  { text: '無効', value: 3, subtitle: 'アカウントが無効化されている状態' },
];

function statusProps(item: (typeof statusList)[0]) {
  return {
    title: item.text,
    subtitle: item.subtitle,
  };
}

async function fetchUsers(params: { page: number; itemsPerPage: number; sortBy: SortBy[]; search: string }) {
  let order = params.sortBy.length ? params.sortBy[0].order : null;
  let sort = params.sortBy.length ? params.sortBy[0].key : null;
  sort = sort === 'roleLabel' ? 'role' : sort === 'statusLabel' ? 'status' : sort === 'createdAt' ? 'created_at' : sort;

  const response = await $api.get<ICollection<IUser>>('/admin/users', {
    params: {
      offset: (params.page - 1) * params.itemsPerPage,
      limit: params.itemsPerPage,
      sort,
      order,
      keyword: params.search,
    },
  });
  users.value = response.data.data ? response.data.data : [];
  totalCount.value = response.data.totalCount;
}

async function updateStatus(userId: number, status: number) {
  await $api.patch(`/admin/users/${userId}/status`, {
    status,
  });
}

onMounted(() => fetchUsers({ page: 1, itemsPerPage: itemsPerPage.value, sortBy: sortBy.value, search: '' }));
</script>
