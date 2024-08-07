<template>
  <v-container fluid v-if="board?.id">
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
        {
          title: board.title,
          disabled: true,
        },
      ]"
    />

    <Form @submit="submit" :validation-schema="schema" v-slot="{ meta, errors }">
      <div class="field mb-2">
        <v-text-field v-model="board.id" label="ID" type="text" variant="outlined" density="compact" readonly />
      </div>

      <div class="field mb-2">
        <v-text-field
          v-model="board.createdAt"
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

      <div class="field mb-2">
        <Field v-model="form.description" name="description" v-slot="{ errors }">
          <v-text-field
            v-model="form.description"
            label="説明"
            type="text"
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

      <v-img v-if="board.thumbnailUrl" :width="300" aspect-ratio="16/9" cover :src="board.thumbnailUrl" />

      <div class="field mt-2 mb-2">
        <v-file-input
          label="サムネイルを選択"
          show-size
          truncate-length="25"
          prepend-icon=""
          variant="outlined"
          dense
          hide-details
          accept="image/*"
          density="compact"
          @change="handleThumbnailChange"
        />
      </div>

      <v-btn type="submit" color="primary" block :disabled="!meta.valid" class="mt-5">保存</v-btn>
    </Form>
  </v-container>
</template>

<script setup lang="ts">
import { Form, Field } from 'vee-validate';
import * as yup from 'yup';

definePageMeta({
  layout: 'admin',
  middleware: ['admin-access-only'],
});

interface IBoard {
  id: number;
  userId: number;
  title: string;
  description: string;
  status: number;
  statusLabel: string;
  thumbnailUrl: string;
  updatedAt: string;
  createdAt: string;
}

const router = useRouter();
const route = useRoute();
const { $api, $formatDate } = useNuxtApp();
const board = ref<IBoard>();
const thumbnailFile = ref<File | null>(null);
const { fetchListPresignedUrl, uploadFilesToS3 } = useActions();

const form = ref({
  title: '',
  description: '',
  status: 0,
  thumbnailUrl: '',
});

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

function handleThumbnailChange(event: Event) {
  const input = event.target as HTMLInputElement;
  if (input.files && input.files[0]) {
    thumbnailFile.value = input.files[0];
  }
}

const schema = yup.object({
  title: yup.string().required('必須項目です'),
  description: yup.string().optional(),
  status: yup.number().required('必須項目です'),
});

async function submit() {
  if (confirm('板情報を更新しますか？')) {
    try {
      if (thumbnailFile.value) {
        const presignedUrls = await fetchListPresignedUrl([thumbnailFile.value.name]);
        const thumbnailUrl = await uploadFilesToS3(presignedUrls[0], thumbnailFile.value);
        form.value.thumbnailUrl = thumbnailUrl;
      }
      await $api.put(`/admin/boards/${board?.value?.id}`, form.value);
      await fetchBoards();
    } catch (err) {
      alert('通信中にエラーが発生しました');
    }
  }
}

async function fetchBoards() {
  const { data } = await $api.get<IBoard>(`/admin/boards/${route.params.boardId}`);
  board.value = data;
  form.value = {
    title: data.title,
    description: data.description,
    status: data.status,
    thumbnailUrl: data.thumbnailUrl,
  };
}

onMounted(() => {
  fetchBoards();
});
</script>
