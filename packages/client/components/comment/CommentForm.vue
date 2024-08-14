<template>
  <div v-if="showReplyForm">
    <p class="form-title">{{ title ?? '書き込み' }}</p>

    <Form @submit="submit" :validation-schema="schema" v-slot="{ meta }">
      <div class="field">
        <Field name="guestName" v-model="form.guestName" v-slot="{ errors }">
          <v-text-field
            v-model="form.guestName"
            label="名前(省略可)"
            variant="outlined"
            hide-details
            counter
            single-line
            :clearable="!payload.isLoggedIn"
            dense
            density="compact"
            :error-messages="errors"
            :disabled="payload.isLoggedIn"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="content" v-model="form.content" v-slot="{ errors }">
          <v-textarea
            v-model="form.content"
            label="コメント"
            rows="3"
            variant="outlined"
            hide-details
            counter
            dense
            single-line
            clearable
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <div class="field">
        <v-file-input
          ref="fileInput"
          label="ファイルを選択"
          show-size
          prepend-icon=""
          variant="outlined"
          single-line
          dense
          multiple
          accept="image/jpeg,image/png,image/gif"
          density="compact"
          chips
          @change="handleAttachmentsChange"
        />
      </div>

      <v-row class="dense">
        <v-col cols="6">
          <v-btn class="clear-button" block @click="() => fileInput!.reset()" type="reset">クリア</v-btn>
        </v-col>
        <v-col cols="6">
          <template v-if="canCommentState || payload.isLoggedIn">
            <v-btn type="submit" class="submit-button" block :disabled="!meta?.valid" color="primary">
              書き込みをする
            </v-btn>
          </template>
          <template v-else>
            <v-btn type="submit" class="submit-button" block disabled>
              {{ remainingTimeText }}後に書き込み可能です
            </v-btn>
          </template>
        </v-col>
      </v-row>

      <p class="text-center text-subtitle-2 mt-4">＊書き込み反映には時間が掛かる場合があります＊</p>
    </Form>
  </div>
</template>

<script setup lang="ts">
import * as yup from 'yup';
import { Form, Field } from 'vee-validate';

interface Attachment {
  url: string;
  displayOrder: number;
  type: 'video' | 'image';
}

interface FormState {
  guestName: string;
  content: string;
  attachments: Attachment[];
}

const props = defineProps<{
  title?: string;
  parentCommentId?: string;
  showReplyForm?: boolean;
  threadId: number;
}>();

const nuxtApp = useNuxtApp();
const router = useRouter();
const route = useRoute();
const { fetchListPresignedUrl, uploadFilesToS3 } = useActions();
const { setLastCommentTime, canComment, timeUntilNextComment } = useStorage();

const { $api, payload } = nuxtApp;

const fileInput = ref<InstanceType<typeof HTMLInputElement>>();
const attachmentFiles = ref<File[]>([]);
const canCommentState = ref(canComment());
const remainingTime = ref<{ minutes: number; seconds: number } | null>(null);

const emit = defineEmits(['submit']);

const form = ref<FormState>({
  guestName: payload.isLoggedIn ? payload?.user?.name || '' : '',
  content: '',
  attachments: [],
});

const updateRemainingTime = () => {
  remainingTime.value = timeUntilNextComment();
  canCommentState.value = canComment();
};

const remainingTimeText = computed(() => {
  if (!remainingTime.value) {
    return '';
  }
  return `${remainingTime.value.minutes}分${remainingTime.value.seconds}秒`;
});

const schema = yup.object({
  guestName: yup.string().optional(),
  content: yup.string().required('コメントは必須項目です'),
});

function handleAttachmentsChange(event: Event): void {
  const input = event.target as HTMLInputElement;
  if (input.files) {
    const files = Array.from(input.files);
    const invalidFiles = files.filter(
      file => !['image/jpeg', 'image/png', 'image/gif'].includes(file.type) || file.size > 1 * 1024 * 1024,
    );

    if (invalidFiles.length > 0) {
      const messages: string[] = [];
      invalidFiles.forEach(file => {
        if (!['image/jpeg', 'image/png', 'image/gif'].includes(file.type)) {
          messages.push(
            `${file.name}は無効なファイルタイプです。許可されているタイプはimage/jpeg, image/png, image/gifです。`,
          );
        }
        if (file.size > 1 * 1024 * 1024) {
          messages.push(`${file.name}は1MBを超えています。ファイルサイズの最大は1MBです。`);
        }
      });

      alert(messages.join('\n'));
      attachmentFiles.value = [];
      if (fileInput.value) {
        fileInput.value.reset();
      }
      return;
    }

    if (files.length > 4) {
      alert('ファイルの最大枚数は4枚です');
      attachmentFiles.value = [];
      if (fileInput.value) {
        fileInput.value.reset();
      }
      return;
    }

    attachmentFiles.value = files;
  }
}

async function submit(_: typeof form.value, { resetForm }: { resetForm: () => void }): Promise<void> {
  if (confirm('本当に書き込みますか？')) {
    try {
      if (attachmentFiles.value && attachmentFiles.value.length > 0) {
        const presignedUrls = await fetchListPresignedUrl(attachmentFiles.value.map(file => file.name));
        const uploadedAttachments = await Promise.all(
          presignedUrls.map(async (url: string, idx: number) => {
            return uploadFilesToS3(url, attachmentFiles.value[idx]).then(uploadedUrl => ({
              url: uploadedUrl,
              displayOrder: idx,
              type: attachmentFiles.value[idx].type.startsWith('video') ? 'video' : ('image' as 'video' | 'image'),
            }));
          }),
        );

        form.value.attachments = uploadedAttachments;
      }

      if (props.parentCommentId) {
        await $api.post(`/threads/${props.threadId}/comments/${props.parentCommentId}/reply`, {
          ...form.value,
          parentCommentId: props.parentCommentId,
        });
      } else {
        await $api.post(`/threads/${props.threadId}/comments`, form.value);
      }
      if (!payload.isLoggedIn) {
        setLastCommentTime();
      }
      resetForm();
      form.value.content = '';
      form.value.attachments = [];
      attachmentFiles.value = [];
      if (fileInput.value) {
        fileInput.value.reset();
      }
      emit('submit');
    } catch (error) {
      alert('書き込み中にエラーが発生しました。');
    }
  }
}

onMounted(() => {
  updateRemainingTime();
  setInterval(updateRemainingTime, 1000); // 1秒ごとに残り時間を更新
});
</script>

<style scoped></style>
