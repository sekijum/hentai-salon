<template>
  <div class="mx-2">
    <PageTitle title="スレ作成" />

    <v-divider />

    <br />

    <Form @submit="submit" :validation-schema="schema" v-slot="{ meta }">
      <div class="field">
        <Field name="boardId" v-model="form.boardId" v-slot="{ errors }">
          <v-select
            v-model="form.boardId"
            label="板"
            item-value="id"
            item-text="title"
            :items="boardSuggestions"
            variant="outlined"
            density="compact"
            dense
            single-line
            :error-messages="errors"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="title" v-model="form.title" v-slot="{ errors }">
          <v-text-field v-model="form.title" label="タイトル" variant="outlined" density="compact" dense single-line :error-messages="errors" />
        </Field>
      </div>

      <div class="field">
        <Field name="description" v-model="form.description" v-slot="{ errors }">
          <v-textarea v-model="form.description" label="説明" variant="outlined" density="compact" dense single-line :error-messages="errors" />
        </Field>
      </div>

      <div class="field">
        <v-combobox
          v-model="form.tagNameList"
          chips
          small-chips
          label="タグ"
          variant="outlined"
          density="compact"
          no-data-text="一致する結果が見つかりませんでした"
          clearable
          multiple
          dense
          single-line
          :items="tagSuggestions"
          @input="fetchTagSuggestions"
        />
      </div>

      <div class="field">
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
          single-line
          @change="handleThumbnailChange"
        />
      </div>

      <v-btn type="submit" color="primary" block :disabled="!meta?.valid" class="mt-5">作成</v-btn>
      <p class="note">＊反映には時間が掛かる場合があります＊</p>
    </Form>
  </div>
</template>

<script setup lang="ts">
import * as yup from 'yup';
import { Form, Field } from 'vee-validate';
import PageTitle from '~/components/PageTitle.vue';
import type { IBoard } from '~/types/board';
import type { IThread } from '~/types/thread';

definePageMeta({ middleware: ['logged-in-access-only'] });

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { fetchListPresignedUrl, uploadFilesToS3 } = useActions();

const { $api } = nuxtApp;

const tagSuggestions = ref<string[]>([]);
const boardSuggestions = ref<{ id: number; title: string }[]>([]);

const thumbnailFile = ref<File>();

const form = ref({
  boardId: route.query.board_id,
  title: '',
  description: '',
  thumbnailUrl: null as string | null,
  tagNameList: [],
});

onMounted(async () => {
  await fetchBoardSuggestions();
});

const schema = yup.object({
  title: yup.string().required('必須項目です'),
  boardId: yup.string().required('必須項目です'),
});

async function fetchTagSuggestions(event) {
  if (!event.target.value) {
    tagSuggestions.value = [];
    return;
  }
  const response = await $api.get<string[]>('/tags/name', {
    params: {
      keyword: event.target.value,
    },
  });
  tagSuggestions.value = response.data ? response.data : [];
}

async function fetchBoardSuggestions() {
  const response = await $api.get<IBoard[]>('/boards');

  if (!response.data) {
    return;
  }
  boardSuggestions.value = response.data.map(board => ({
    id: board.id,
    title: board.title,
  }));
}

function handleThumbnailChange(event: Event) {
  const input = event.target as HTMLInputElement;
  if (input.files && input.files[0]) {
    thumbnailFile.value = input.files[0];
  }
}

async function submit() {
  try {
    if (confirm('スレッドを作成しますか？')) {
      if (thumbnailFile.value) {
        const presignedUrls = await fetchListPresignedUrl([thumbnailFile.value.name]);
        const thumbnailUrl = await uploadFilesToS3(presignedUrls[0], thumbnailFile.value);
        form.value.thumbnailUrl = thumbnailUrl;
      }
      const response = await $api.post<IThread>('/threads', form.value);
      router.push(`/threads/${response.data.id}`);
    }
  } catch (error) {
    alert('通信中にエラーが発生しました');
  }
}

useHead({
  title: '変態サロン | スレ作成',
});
</script>

<style scoped>
.note {
  font-size: 12px;
  color: grey;
  text-align: center;
  margin-top: 8px;
}
</style>
