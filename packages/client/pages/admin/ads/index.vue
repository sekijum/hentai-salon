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
          title: '問い合わせ一覧',
          disabled: false,
          href: '/admin/ads',
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
            <li><strong>id:</strong> 特定の広告IDで検索します。例: <code>id:123</code></li>
            <li>
              <strong>active:</strong> 特定のアクティブを持つ広告を検索します。以下の例を参照してください:
              <ul>
                <li>非アクティブな広告を検索: <code>active:0</code></li>
                <li>アクティブな広告を検索: <code>active:1</code></li>
              </ul>
            </li>
            <li>また、名前やメールアドレスを部分一致で検索することも可能です。</li>
          </ul>
        </v-expansion-panel-text>
      </v-expansion-panel>
    </v-expansion-panels>

    <v-divider class="border-opacity-0" :thickness="20" />

    <v-row>
      <v-col>
        <v-btn @click="() => router.push('/admin/ads/create')" type="submit"> 新規登録 </v-btn>
      </v-col>
    </v-row>

    <v-divider class="border-opacity-0" :thickness="20" />

    <v-data-table-server
      v-model:sort-by="sortBy"
      v-model:items-per-page="itemsPerPage"
      :headers="headers"
      :items="ads"
      :items-length="totalCount"
      :hover="true"
      item-value="id"
      @update:options="fetchAds"
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
          <td>
            <Ad :content="item.content" />
          </td>
          <td>
            <v-select
              v-model="item.isActive"
              :items="activeList"
              :item-props="activeProps"
              label="ステータス"
              variant="outlined"
              density="compact"
              hide-details
              counter
              dense
              single-line
              :item-title="'text'"
              :item-value="'value'"
              @update:modelValue="updateActive(item.id, item.isActive)"
            />
          </td>
          <td>
            <v-icon class="me-2" size="small" @click="router.push(`/admin/ads/${item.id}`)">mdi-pencil</v-icon>
            <v-icon size="small" @click="() => deleteAd(item.id)">mdi-delete</v-icon>
          </td>
        </tr>
      </template>
    </v-data-table-server>
  </v-container>
</template>

<script setup lang="ts">
import type { ICollection } from '~/types/collection';
import type { IAd } from '~/types/ad';
import Ad from '~/components/Ad.vue';

definePageMeta({
  layout: 'admin',
  middleware: ['admin-access-only'],
});

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
const ads = ref<IAd[]>([]);

const headers = [
  { title: 'ID', align: 'start', sortable: true, key: 'id' },
  { title: '内容', key: 'content', sortable: false },
  { title: 'アクティブ', key: 'isActive' },
  { title: '操作', key: 'actions', sortable: false },
];

const activeList = [
  { text: 'アクティブ', value: 1, subtitle: '現在この広告は有効です' },
  { text: '非アクティブ', value: 0, subtitle: '現在この広告は無効です' },
];

function activeProps(item: (typeof activeList)[0]) {
  return {
    title: item.text,
    subtitle: item.subtitle,
  };
}

async function fetchAds(params: { page: number; itemsPerPage: number; sortBy: SortBy[]; search: string }) {
  let order = params.sortBy.length ? params.sortBy[0].order : null;
  let sort = params.sortBy.length ? params.sortBy[0].key : null;
  sort = sort === 'isActive' ? 'is_active' : sort === 'createdAt' ? 'created_at' : sort;

  const response = await $api.get<ICollection<IAd>>('/admin/ads', {
    params: {
      offset: (params.page - 1) * params.itemsPerPage,
      limit: params.itemsPerPage,
      sort,
      order,
      keyword: params.search,
    },
  });
  ads.value = response.data.data ? response.data.data : [];
  totalCount.value = response.data.totalCount;
}

async function deleteAd(adId: number) {
  if (confirm('本当に削除しますか？')) {
    try {
      await $api.delete(`/admin/ads/${adId}`);
      await fetchAds({ page: 1, itemsPerPage: itemsPerPage.value, sortBy: sortBy.value, search: '' });
    } catch (err) {
      alert('通信中にエラーが発生しました');
    }
  }
}

async function updateActive(adId: number, isActive: number) {
  await $api.patch(`/admin/ads/${adId}/is-active`, {
    isActive,
  });
}

onMounted(() => fetchAds({ page: 1, itemsPerPage: itemsPerPage.value, sortBy: sortBy.value, search: '' }));
</script>
