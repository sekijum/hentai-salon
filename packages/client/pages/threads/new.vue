<template>
  <div class="mx-2">
    <PageTitle title="スレ作成" />

    <v-divider />

    <br />

    <Form @submit="submit" :validation-schema="schema" v-slot="{ meta }">
      <div class="field">
        <Field name="boardId" v-slot="{ field, errorMessage }">
          <v-select
            v-model="form.boardId"
            v-bind="field"
            label="板"
            item-value="id"
            item-text="title"
            :items="boardSuggestions"
            variant="outlined"
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="title" v-slot="{ field, errorMessage }">
          <v-text-field
            v-model="form.title"
            v-bind="field"
            label="タイトル"
            variant="outlined"
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="description" v-slot="{ field, errorMessage }">
          <v-textarea
            v-model="form.description"
            v-bind="field"
            label="説明"
            variant="outlined"
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
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
          clearable
          multiple
          :items="tagSuggestions"
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
  await fetchTagSuggestions();
  await fetchBoardSuggestions();
});

const schema = yup.object({
  title: yup.string().required('必須項目です'),
  boardId: yup.string().required('必須項目です'),
});

async function fetchTagSuggestions() {
  const response = await $api.get<string[]>('/tags/name');
  tagSuggestions.value = response.data;
}

async function fetchBoardSuggestions() {
  const response = await $api.get<IBoard[]>('/boards');

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
      alert('スレッドを作成しました。');
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
