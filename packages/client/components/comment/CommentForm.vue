<template>
  <div>
    <p class="form-title">{{ title ?? '書き込み' }}</p>

    <Form @submit="submit" :validation-schema="schema" v-slot="{ meta }">
      <div class="field">
        <Field name="guestName" v-slot="{ field, errorMessage }">
          <v-text-field
            v-model="form.guestName"
            v-bind="field"
            label="名前(省略可)"
            variant="outlined"
            hide-details
            counter
            single-line
            clearable
            dense
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="content" v-slot="{ field, errorMessage }">
          <v-textarea
            v-model="form.content"
            v-bind="field"
            label="コメント"
            rows="3"
            variant="outlined"
            hide-details
            counter
            dense
            single-line
            clearable
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <div class="field">
        <v-file-input
          ref="fileInput"
          label="ファイルを選択"
          show-size
          truncate-length="25"
          prepend-icon=""
          variant="outlined"
          dense
          hide-details
          multiple
          accept="image/jpeg,image/png,image/gif"
          density="compact"
          chips
          @change="handleAttachmentsChange"
        />
      </div>

      <v-btn class="clear-button" block @click="clearForm">クリア</v-btn>
      <v-btn type="submit" class="submit-button" block :disabled="!meta?.valid">書き込みをする</v-btn>
      <p class="note">＊書き込み反映には時間が掛かる場合があります＊</p>
    </Form>

    <OverlayLoagind :isLoading="isLoading" title="書き込み中" />
  </div>
</template>

<script setup lang="ts">
import { Form, Field, ErrorMessage } from 'vee-validate';
import * as yup from 'yup';
import OverlayLoagind from '~/components/OverlayLoagind.vue';

interface Attachment {
  url: string;
  displayOrder: number;
  type: 'Video' | 'Image';
}

const props = defineProps<{ title?: string; parentCommentId?: number }>();

const attachmentFiles = ref<File[] | null>(null);
const nuxtApp = useNuxtApp();
const router = useRouter();
const route = useRoute();
const { $api } = nuxtApp;
const { uploadFilesToImgur } = useActions();
const fileInput = ref<InstanceType<typeof HTMLInputElement>>();
const isLoading = ref(false);

const form = ref({
  guestName: '',
  content: '',
  attachments: [] as Attachment[],
});

const schema = yup.object({
  guestName: yup.string().optional(),
  content: yup.string().required('コメントは必須項目です'),
  attachments: yup
    .array()
    .of(
      yup.object({
        url: yup.string().required(),
        displayOrder: yup.number().required(),
        type: yup.mixed<'Video' | 'Image'>().oneOf(['Video', 'Image']).required(),
      }),
    )
    .optional(),
});

function clearForm(): void {
  if (confirm('本当にクリアしますか？')) {
    form.value.guestName = '';
    form.value.content = '';
    attachmentFiles.value = null;
    if (fileInput.value) {
      fileInput.value.reset();
    }
  }
}

function handleAttachmentsChange(event: Event): void {
  const input = event.target as HTMLInputElement;
  if (input.files) {
    const files = Array.from(input.files);
    if (files.length > 4) {
      alert('ファイルの最大枚数は4枚です');
      attachmentFiles.value = null;
      if (fileInput.value) {
        fileInput.value.reset();
      }
    } else {
      attachmentFiles.value = files;
    }
  }
}

async function submit(): Promise<void> {
  if (confirm('本当に書き込みますか？')) {
    isLoading.value = true;
    try {
      if (attachmentFiles.value && attachmentFiles.value.length > 0) {
        const uploadedAttachments = await uploadFilesToImgur(attachmentFiles.value);
        form.value.attachments = uploadedAttachments;
        console.log(uploadedAttachments);
      }

      const threadId = parseInt(route.params.id.toString(), 10);
      if (props.parentCommentId) {
        await $api.post(`/threads/${threadId}/comments/${props.parentCommentId}/reply`, {
          ...form.value,
          parentCommentId: props.parentCommentId,
        });
        alert('返信に成功しました。');
      } else {
        await $api.post(`/threads/${threadId}/comments/`, form.value);
        alert('書き込みに成功しました。');
      }
      router.go(0);
    } catch (error) {
      console.error('通信中にエラーが発生しました:', error);
      alert('コメントの送信中にエラーが発生しました。');
    }
    isLoading.value = false;
  }
}
</script>

<style scoped>
.form-title {
  font-size: 1.2em;
}

.v-text-field,
.v-textarea,
.v-file-input {
  margin-bottom: 0px !important;
}

.v-text-field input,
.v-textarea textarea {
  font-size: 12px;
}

.clear-button,
.submit-button {
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-bottom: 0px;
}

.clear-button {
  background-color: #f0f0f0;
  color: black;
  margin-top: 10px;
  margin-bottom: 10px;
}

.submit-button {
  background-color: #007bff;
  color: white;
}

.submit-button:hover {
  background-color: #0056b3;
}

.clear-button:hover {
  background-color: #e0e0e0;
}

.note {
  font-size: 12px;
  color: grey;
  text-align: center;
  margin-top: 8px;
}
</style>
