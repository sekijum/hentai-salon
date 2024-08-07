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
          title: 'タグ一覧',
          disabled: false,
          href: '/admin/tags',
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
      :items="tags"
      :items-length="totalCount"
      :hover="true"
      item-value="id"
      @update:options="fetchTags"
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
          <td>{{ item.threadCount }}</td>
          <td>{{ $formatDate(item.createdAt) }}</td>
          <td>
            <v-icon size="small" @click="() => deleteTag(item.id)">mdi-delete</v-icon>
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

interface ITag {
  id: number;
  name: string;
  threadCount: number;
  createdAt: string;
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
const tags = ref<ITag[]>([]);

const headers = [
  { title: 'ID', align: 'start', sortable: true, key: 'id' },
  { title: '名前', key: 'name' },
  { title: 'スレッド数', key: 'threadCount', sortable: false },
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

async function fetchTags(params: { page: number; itemsPerPage: number; sortBy: SortBy[]; search: string }) {
  let order = params.sortBy.length ? params.sortBy[0].order : null;
  let sort = params.sortBy.length ? params.sortBy[0].key : null;
  sort = sort === 'createdAt' ? 'created_at' : sort;

  const response = await $api.get<ICollection<ITag>>('/admin/tags', {
    params: {
      offset: (params.page - 1) * params.itemsPerPage,
      limit: params.itemsPerPage,
      sort,
      order,
      keyword: params.search,
    },
  });
  tags.value = response.data.data ? response.data.data : [];
  totalCount.value = response.data.totalCount;
}

async function deleteTag(tagId: number) {
  if (confirm('本当に削除しますか？')) {
    try {
      await $api.delete(`/admin/tags/${tagId}`);
      await fetchTags({ page: 1, itemsPerPage: itemsPerPage.value, sortBy: sortBy.value, search: '' });
    } catch (err) {
      alert('通信中にエラーが発生しました');
    }
  }
}

onMounted(() => fetchTags({ page: 1, itemsPerPage: itemsPerPage.value, sortBy: sortBy.value, search: '' }));
</script>
