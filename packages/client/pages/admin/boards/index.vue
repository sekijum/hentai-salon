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
          title: '板一覧',
          disabled: false,
          href: '/admin/boards',
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
            <li><strong>id:</strong> 特定の板IDで検索します。例: <code>id:123</code></li>
            <li>
              <strong>status:</strong> 特定のステータスを持つボードを検索します。以下の例を参照してください:
              <ul>
                <li>公開されているボードを検索: <code>status:0</code></li>
                <li>非公開のボードを検索: <code>status:1</code></li>
                <li>保留中のボードを検索: <code>status:2</code></li>
                <li>アーカイブされたボードを検索: <code>status:3</code></li>
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
      :items="boards"
      :items-length="totalCount"
      :hover="true"
      item-value="id"
      @update:options="fetchBoards"
      items-per-page-text="表示行数"
      density="compact"
      must-sort
      no-data-text="検索結果は0件です"
      v-model:search="search"
      class="border-sm"
    >
      <template #item="{ item }">
        <tr>
          <td>{{ item.id }}</td>
          <td>{{ item.title }}</td>
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
            <v-icon class="me-2" size="small" @click="router.push(`/admin/boards/${item.id}`)">mdi-pencil</v-icon>
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

interface IBoard {
  id: 96;
  userId: 9;
  title: '掲示板4b5ba504';
  description: '掲示板の説明94836a7b';
  status: 0;
  statusLabel: '公開';
  thumbnailUrl: 'https://picsum.photos/seed/95/550/397.webp';
  updatedAt: '2024-07-27T11:35:01Z';
  createdAt: '2024-07-27T11:35:01Z';
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
const boards = ref<IBoard[]>([]);

const headers = [
  { title: 'ID', align: 'start', sortable: true, key: 'id' },
  { title: '名前', key: 'タイトル' },
  { title: 'メール', key: 'status' },
  { title: '登録日', key: 'createdAt' },
  { title: '操作', key: 'actions', sortable: false },
];

const statusList = [
  { text: '公開', value: 0, subtitle: '一般公開されています' },
  { text: '非公開', value: 1, subtitle: 'アクセスできるのは管理者のみです' },
  { text: '保留', value: 2, subtitle: 'まだ公開されていません' },
  { text: 'アーカイブ', value: 3, subtitle: 'アーカイブされ、表示されていません' },
];

function statusProps(item: (typeof statusList)[0]) {
  return {
    title: item.text,
    subtitle: item.subtitle,
  };
}

async function fetchBoards(params: { page: number; itemsPerPage: number; sortBy: SortBy[]; search: string }) {
  let order = params.sortBy.length ? params.sortBy[0].order : null;
  let sort = params.sortBy.length ? params.sortBy[0].key : null;
  sort = sort === 'statusLabel' ? 'status' : sort === 'createdAt' ? 'created_at' : sort;

  const response = await $api.get<ICollection<IBoard>>('/admin/boards', {
    params: {
      offset: (params.page - 1) * params.itemsPerPage,
      limit: params.itemsPerPage,
      sort,
      order,
      keyword: params.search,
    },
  });
  boards.value = response.data.data ? response.data.data : [];
  totalCount.value = response.data.totalCount;
}

async function updateStatus(boardId: number, status: number) {
  await $api.patch(`/admin/boards/${boardId}/status`, {
    status,
  });
}

onMounted(() => fetchBoards({ page: 1, itemsPerPage: itemsPerPage.value, sortBy: sortBy.value, search: '' }));
</script>
