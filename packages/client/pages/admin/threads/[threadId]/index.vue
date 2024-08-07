<template>
  <v-container fluid v-if="thread?.id">
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
        {
          title: thread.title,
          disabled: true,
        },
      ]"
    />

    <v-tabs v-model="tab" grow>
      <v-tab :value="1">スレッド</v-tab>
      <v-tab :value="2">コメント</v-tab>
    </v-tabs>

    <v-window v-model="tab">
      <br />
      <v-window-item :value="1">
        <Form @submit="submit" :validation-schema="schema" v-slot="{ meta, errors }">
          <div class="field mb-2">
            <v-text-field v-model="thread.id" label="ID" type="text" variant="outlined" density="compact" readonly />
          </div>

          <div class="field mb-2">
            <v-text-field
              v-model="thread.createdAt"
              label="登録日"
              type="text"
              variant="outlined"
              density="compact"
              readonly
            />
          </div>

          <div class="field mb-2">
            <Field v-model="form.title" name="title" v-slot="{ errors }">
              <v-text-field
                v-model="form.title"
                label="タイトル"
                type="text"
                variant="outlined"
                density="compact"
                :error-messages="errors"
              />
            </Field>
          </div>

          <div class="field">
            <Field v-model="form.description" name="description" v-slot="{ errors }">
              <v-textarea
                v-model="form.description"
                label="説明"
                variant="outlined"
                density="compact"
                :error-messages="errors"
              />
            </Field>
          </div>

          <div class="field mb-2">
            <Field v-model="form.status" name="status" v-slot="{ errors }">
              <v-select
                v-model="form.status"
                :items="statusList"
                :item-props="statusProps"
                label="ステータス"
                variant="outlined"
                density="compact"
                :item-title="'text'"
                :item-value="'value'"
                :error-messages="errors"
              />
            </Field>
          </div>

          <v-btn type="submit" color="primary" block :disabled="!meta.valid">保存</v-btn>
        </Form>
      </v-window-item>
      <v-window-item :value="2">
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
                  <strong>status:</strong>
                  特定のスレッドコメントステータスを持つコメントを検索します。以下の例を参照してください:
                  <ul>
                    <li>可視状態のコメントを検索: <code>status:0</code></li>
                    <li>削除されたコメントを検索: <code>status:1</code></li>
                  </ul>
                </li>
                <li>また、内容を部分一致で検索することも可能です。</li>
              </ul>
            </v-expansion-panel-text>
          </v-expansion-panel>
        </v-expansion-panels>

        <v-divider class="border-opacity-0" :thickness="20" />

        <v-data-table-server
          v-model:sort-by="sortBy"
          v-model:items-per-page="itemsPerPage"
          :headers="headers"
          :items="comments"
          :items-length="totalCount"
          :hover="true"
          item-value="id"
          @update:options="fetchThread"
          items-per-page-text="表示行数"
          density="compact"
          must-sort
          no-data-text="検索結果は0件です"
          v-model:search="search"
        >
          <template #item="{ item }">
            <tr>
              <td>{{ item.id }}</td>
              <td>{{ item.content }}</td>
              <td>{{ $formatDate(item.createdAt) }}</td>
              <td>
                <v-icon size="small" @click="() => deleteComment(item.id)">mdi-delete</v-icon>
              </td>
            </tr>
          </template>
        </v-data-table-server>
      </v-window-item>
    </v-window>
  </v-container>
</template>

<script setup lang="ts">
import { Form, Field } from 'vee-validate';
import * as yup from 'yup';
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
  board?: {
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
  comments: ICollection<IComment>;
}

interface IComment {
  id: number;
  threadId: number;
  userId: number;
  content: string;
  status: number;
  statusLabel: string;
  createdAt: string;
  updatedAt: string;
}

type SortBy = {
  key: string;
  order: 'asc' | 'desc';
};

const initSortBy: SortBy = {
  key: 'id',
  order: 'desc',
};

const tab = ref(null);
const sortBy = ref<SortBy[]>([initSortBy]);
const itemsPerPage = ref(20);
const totalCount = ref(0);
const search = ref('');
const router = useRouter();
const route = useRoute();
const { $api, $formatDate } = useNuxtApp();
const thread = ref<IThread>();
const comments = ref<IComment[]>([]);
const form = ref({
  title: '',
  description: '',
  status: 0,
  thumbnailUrl: '',
});

const headers = [
  { title: 'ID', align: 'start', sortable: true, key: 'id' },
  { title: '内容', key: 'content' },
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

const schema = yup.object({
  title: yup.string().required('必須項目です'),
  description: yup.string().required('必須項目です'),
  status: yup.number().required('必須項目です'),
});

async function submit() {
  if (confirm('スレッド情報が更新されました。')) {
    try {
      await $api.put(`/admin/threads/${thread?.value?.id}`, form.value);
      await fetchThread({ page: 1, itemsPerPage: itemsPerPage.value, sortBy: sortBy.value, search: '' });
    } catch (err) {
      alert('通信中にエラーが発生しました');
    }
  }
}

async function deleteComment(commentId: number) {
  if (confirm('本当に削除しますか？')) {
    try {
      await $api.delete(`/admin/threads/${thread?.value?.id}/comments/${commentId}`);
      await fetchThread({ page: 1, itemsPerPage: itemsPerPage.value, sortBy: sortBy.value, search: '' });
    } catch (err) {
      alert('通信中にエラーが発生しました');
    }
  }
}

async function fetchThread(params: { page: number; itemsPerPage: number; sortBy: SortBy[]; search: string }) {
  let order = params.sortBy.length ? params.sortBy[0].order : null;
  let sort = params.sortBy.length ? params.sortBy[0].key : null;
  sort = sort === 'statusLabel' ? 'status' : sort === 'createdAt' ? 'created_at' : sort;

  const { data } = await $api.get<IThread>(`/admin/threads/${route.params.threadId}`, {
    params: {
      offset: (params.page - 1) * params.itemsPerPage,
      limit: params.itemsPerPage,
      sort,
      order,
      keyword: params.search,
    },
  });
  comments.value = data.comments.data ? data.comments.data : [];
  totalCount.value = data.comments.totalCount;
  thread.value = data;
  form.value = {
    title: data.title,
    description: data.description,
    status: data.status,
    thumbnailUrl: data.thumbnailUrl,
  };
}

onMounted(() => {
  fetchThread({ page: 1, itemsPerPage: itemsPerPage.value, sortBy: sortBy.value, search: '' });
});
</script>
