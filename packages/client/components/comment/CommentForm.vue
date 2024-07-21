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
            :readonly="payload.isLoggedIn"
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

      <v-row class="no-gutters">
        <v-col cols="6">
          <v-btn class="clear-button" block @click="clearForm">クリア</v-btn>
        </v-col>
        <v-col cols="6">
          <template v-if="canCommentState || payload.isLoggedIn">
            <v-btn type="submit" class="submit-button" block :disabled="!meta?.valid">書き込みをする</v-btn>
          </template>
          <template v-else>
            <v-btn type="submit" class="submit-button" block disabled>
              {{ remainingTimeText }}後に書き込み可能です
            </v-btn>
          </template>
        </v-col>
      </v-row>

      <p class="note">＊書き込み反映には時間が掛かる場合があります＊</p>
    </Form>
  </div>
</template>

<script setup lang="ts">
import * as yup from 'yup';
import { Form, Field } from 'vee-validate';

interface Attachment {
  url: string;
  displayOrder: number;
  type: 'Video' | 'Image';
}

interface FormState {
  guestName: string;
  content: string;
  attachments: Attachment[];
}

const props = defineProps<{ title?: string; parentCommentId?: number; showReplyForm?: boolean }>();

const nuxtApp = useNuxtApp();
const router = useRouter();
const route = useRoute();
const { uploadFilesToImgur } = useActions();
const { setLastCommentTime, canComment, timeUntilNextComment } = useStorage();

const { $api, payload } = nuxtApp;

const fileInput = ref<InstanceType<typeof HTMLInputElement>>();
const attachmentFiles = ref<File[] | null>(null);
const canCommentState = ref(canComment());
const remainingTime = ref<{ minutes: number; seconds: number } | null>(null);

const emit = defineEmits(['submit']);

const snackbar = useState('isSnackbar', () => {
  return { isSnackbar: false, text: '' };
});

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
      fileInput.value.value = '';
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
        fileInput.value.value = '';
      }
    } else {
      attachmentFiles.value = files;
    }
  }
}

async function submit(): Promise<void> {
  if (confirm('本当に書き込みますか？')) {
    try {
      if (attachmentFiles.value && attachmentFiles.value.length > 0) {
        const uploadedAttachments = await uploadFilesToImgur(attachmentFiles.value);
        form.value.attachments = uploadedAttachments;
      }

      if (props.parentCommentId) {
        await $api.post(`/threads/${route.params.threadId}/comments/${props.parentCommentId}/reply`, {
          ...form.value,
          parentCommentId: props.parentCommentId,
        });
      } else {
        await $api.post(`/threads/${route.params.threadId}/comments/`, form.value);
      }
      if (!payload.isLoggedIn) setLastCommentTime();
      snackbar.value.isSnackbar = true;
      snackbar.value.text = '書き込みに成功しました。';
      form.value.content = '';
      attachmentFiles.value = null;
      if (fileInput.value) {
        fileInput.value.value = '';
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

.v-row.no-gutters {
  margin-right: 0;
  margin-left: 0;
}

.v-row.no-gutters > .v-col {
  padding-right: 0;
  padding-left: 0;
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
