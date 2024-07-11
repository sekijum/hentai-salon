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
          v-model="form.tagNames"
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
import { Form, Field, ErrorMessage } from 'vee-validate';
import * as yup from 'yup';
import PageTitle from '~/components/PageTitle.vue';
import type { IBoard } from '~/types/board';

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { $api } = nuxtApp;
const { fetchListPresignedUrl, uploadFileToS3WithPresignedUrl } = useActions();

const tagSuggestions = ref<string[]>([]);
const boardSuggestions = ref<{ id: number; title: string }[]>([]);

const thumbnailFile = ref<File | null>(null);

const form = ref({
  boardId: route.query.board_id,
  title: '',
  description: '',
  thumbnailUrl: null,
  tagNames: [],
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
  try {
    const response = await $api.get<string[]>('/tags/names');
    tagSuggestions.value = response.data;
  } catch (error) {
    console.error('通信中にエラーが発生しました:', error);
  }
}

async function fetchBoardSuggestions() {
  try {
    const response = await $api.get<IBoard[]>('/boards');

    boardSuggestions.value = response.data.map(board => ({
      id: board.id,
      title: board.title,
    }));
  } catch (error) {
    console.error('通信中にエラーが発生しました:', error);
  }
}

function handleThumbnailChange(event: Event) {
  const input = event.target as HTMLInputElement;
  if (input.files && input.files[0]) {
    avatarFile.value = input.files[0];
  }
}

async function submit() {
  try {
    if (thumbnailFile.value) {
      const presignedUrls = await fetchListPresignedUrl([thumbnailFile.value.name]);
      const thumbnailUrl = await uploadFileToS3WithPresignedUrl(presignedUrls[0], thumbnailFile.value);
      form.value.thumbnailUrl = thumbnailUrl;
    }
    await $api.post('/threads', form.value);
    alert('スレッドが正常に作成されました。');
    router.push('/');
  } catch (error) {
    console.error('通信中にエラーが発生しました:', error);
  }
}
</script>

<style scoped>
.note {
  font-size: 12px;
  color: grey;
  text-align: center;
  margin-top: 8px;
}
</style>
