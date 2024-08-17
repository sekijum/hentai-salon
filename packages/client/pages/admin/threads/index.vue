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
          title: 'スレッド一覧',
          disabled: false,
          href: '/admin/threads',
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
              <strong>status:</strong> 特定のスレッドステータスを持つスレッドを検索します。以下の例を参照してください:
              <ul>
                <li>公開中のスレッドを検索: <code>status:0</code></li>
                <li>保留中のスレッドを検索: <code>status:1</code></li>
                <li>アーカイブされたスレッドを検索: <code>status:2</code></li>
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
      :items="threads"
      :items-length="totalCount"
      :hover="true"
      item-value="id"
      @update:options="fetchThreads"
      items-per-page-text="表示行数"
      density="compact"
      must-sort
      no-data-text="検索結果は0件です"
      v-model:search="search"
    >
      <template #item="{ item }">
        <tr>
          <td>{{ item.id }}</td>
          <td>{{ item.title }}</td>
          <td>{{ item.board.title }}</td>
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
            <v-icon @click="router.push(`/admin/threads/${item.id}`)">mdi-pencil</v-icon>
            <v-icon @click="() => deleteThread(item.id)">mdi-delete</v-icon>
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

interface IThread {
  id: number;
  userId: number;
  boardId: number;
  title: string;
  description: string;
  status: number;
  statusLabel: string;
  ipAddress: string;
  thumbnailUrl: string;
  updatedAt: string;
  createdAt: string;
  board: {
    createdAt: string;
    description: string;
    id: number;
    status: number;
    statusLabel: string;
    thumbnailUrl: string;
    title: string;
    updatedAt: string;
    userId: number;
  };
}

const route = useRoute();
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
const threads = ref<IThread[]>([]);

const headers = [
  { title: 'ID', align: 'start', sortable: true, key: 'id' },
  { title: '板', key: 'boardId', sortable: false },
  { title: 'タイトル', key: 'title' },
  { title: 'ステータス', key: 'statusLabel' },
  { title: '登録日', key: 'createdAt' },
  { title: '操作', key: 'actions', sortable: false },
];

const statusList = [
  { text: '公開', value: 0, subtitle: 'スレッドが現在公開されている状態' },
  { text: '保留', value: 1, subtitle: 'スレッドが現在保留されている状態' },
  { text: 'アーカイブ', value: 2, subtitle: 'スレッドがアーカイブされている状態' },
];

function statusProps(item: (typeof statusList)[0]) {
  return {
    title: item.text,
    subtitle: item.subtitle,
  };
}

async function fetchThreads(params: { page: number; itemsPerPage: number; sortBy: SortBy[]; search: string }) {
  let order = params.sortBy.length ? params.sortBy[0].order : null;
  let sort = params.sortBy.length ? params.sortBy[0].key : null;
  sort = sort === 'roleLabel' ? 'role' : sort === 'statusLabel' ? 'status' : sort === 'createdAt' ? 'created_at' : sort;

  const response = await $api.get<ICollection<IThread>>('/admin/threads', {
    params: {
      offset: (params.page - 1) * params.itemsPerPage,
      limit: params.itemsPerPage,
      sort,
      order,
      keyword: params.search,
    },
  });
  threads.value = response.data.data ? response.data.data : [];
  totalCount.value = response.data.totalCount;
}

async function updateStatus(threadId: number, status: number) {
  await $api.patch(`/admin/threads/${threadId}/status`, {
    status,
  });
}

async function deleteThread(threadId: number) {
  if (confirm('本当に削除しますか？')) {
    try {
      await $api.delete(`/admin/threads/${threadId}`);
      await fetchThreads({ page: 1, itemsPerPage: itemsPerPage.value, sortBy: sortBy.value, search: '' });
    } catch (err) {
      alert('通信中にエラーが発生しました');
    }
  }
}

onMounted(() => fetchThreads({ page: 1, itemsPerPage: itemsPerPage.value, sortBy: sortBy.value, search: '' }));
</script>
