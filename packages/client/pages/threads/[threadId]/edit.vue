<template>
  <div class="mx-2">
    <PageTitle title="スレ編集" />

    <v-divider />

    <br />

    <Form @submit="submit" :validation-schema="schema" v-slot="{ meta }">
      <div class="field">
        <v-text-field v-model="thread.title" label="タイトル" variant="outlined" density="compact" readonly />
      </div>

      <div class="field">
        <Field name="description" v-model="form.description" v-slot="{ errors }">
          <v-textarea
            v-model="form.description"
            label="説明"
            variant="outlined"
            density="compact"
            dense
            single-line
            :error-messages="errors"
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
          dense
          single-line
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
import type { IThread } from '~/types/thread';

definePageMeta({ middleware: ['logged-in-access-only'] });

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { fetchListPresignedUrl, uploadFilesToS3 } = useActions();

const { $api, payload } = nuxtApp;

const tagSuggestions = ref<string[]>([]);

const thumbnailFile = ref<File>();

const form = ref<{
  description: string;
  thumbnailUrl: string | null;
  tagNameList: string[];
}>({
  description: '',
  thumbnailUrl: null as string | null,
  tagNameList: [],
});

onMounted(async () => {
  await fetchTagSuggestions();
  await fetchThread();
});

const thread = ref<IThread>({
  id: 0,
  title: '',
  description: '',
  thumbnailUrl: '',
  tagNameList: [],
  commentCount: 0,
  userId: 0,
  comments: { totalCount: 0, limit: 0, offset: 0, data: [] },
  attachments: { totalCount: 0, limit: 0, offset: 0, data: [] },
  isLiked: false,
});

const schema = yup.object({});

async function fetchTagSuggestions() {
  const response = await $api.get<string[]>('/tags/name');
  tagSuggestions.value = response.data;
}

function handleThumbnailChange(event: Event) {
  const input = event.target as HTMLInputElement;
  if (input.files && input.files[0]) {
    thumbnailFile.value = input.files[0];
  }
}

async function submit() {
  if (confirm('スレッドを編集しますか？')) {
    try {
      if (thumbnailFile.value) {
        const presignedUrls = await fetchListPresignedUrl([thumbnailFile.value.name]);
        const thumbnailUrl = await uploadFilesToS3(presignedUrls[0], thumbnailFile.value);
        form.value.thumbnailUrl = thumbnailUrl;
      }
      await $api.put<IThread>(`/threads/${route.params.threadId}`, form.value);
    } catch (err) {
      alert('通信中にエラーが発生しました');
    }
  }
}

async function fetchThread() {
  const threadId = route.params.threadId;
  const response = await $api.get<IThread>(`/threads/${threadId}`);
  if (response.data.userId !== payload?.user?.id) router.push('/');
  thread.value = response.data;
  form.value = {
    description: thread.value.description,
    thumbnailUrl: thread.value.thumbnailUrl,
    tagNameList: thread.value.tagNameList,
  };
}

useHead({
  title: '変態サロン | スレ編集',
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
