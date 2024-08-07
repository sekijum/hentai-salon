<template>
  <v-container fluid v-if="contact?.id">
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
          href: '/admin/contacts',
        },
        {
          title: contact.subject,
          disabled: true,
        },
      ]"
    />

    <Form @submit="submit" :validation-schema="schema" v-slot="{ meta, errors }">
      <div class="field mb-2">
        <v-text-field v-model="contact.id" label="ID" type="text" variant="outlined" density="compact" readonly />
      </div>

      <div class="field mb-2">
        <v-text-field
          v-model="contact.createdAt"
          label="登録日"
          type="text"
          variant="outlined"
          density="compact"
          readonly
        />
      </div>

      <div class="field mb-2">
        <v-text-field
          v-model="contact.email"
          label="メールアドレス"
          type="email"
          variant="outlined"
          density="compact"
          readonly
        />
      </div>

      <div class="field mb-2">
        <v-text-field v-model="contact.subject" label="件名" variant="outlined" density="compact" readonly />
      </div>

      <div class="field mb-2">
        <v-textarea v-model="contact.message" label="内容" variant="outlined" density="compact" readonly />
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
  </v-container>
</template>

<script setup lang="ts">
import { Form, Field } from 'vee-validate';
import * as yup from 'yup';

definePageMeta({
  layout: 'admin',
  middleware: ['admin-access-only'],
});

interface IContact {
  id: number;
  email: string;
  subject: string;
  message: string;
  status: number;
  statusLabel: string;
  createdAt: string;
  updatedAt: string;
}
const router = useRouter();
const route = useRoute();
const { $api, $formatDate } = useNuxtApp();
const contact = ref<IContact>();

const form = ref({
  status: 0,
});

const statusList = [
  { text: '未対応', value: 0, subtitle: '問い合わせがまだ対応されていない状態' },
  { text: '対応中', value: 1, subtitle: '問い合わせが対応中である状態' },
  { text: '完了', value: 2, subtitle: '問い合わせが完了した状態' },
];

function statusProps(item: (typeof statusList)[0]) {
  return {
    title: item.text,
    subtitle: item.subtitle,
  };
}

const schema = yup.object({
  status: yup.number().required('必須項目です'),
});

async function submit() {
  try {
    await $api.patch(`/admin/contacts/${contact?.value?.id}/status`, form.value);
    await fetchContact();
  } catch (err) {
    alert('通信中にエラーが発生しました');
  }
}

async function fetchContact() {
  const { data } = await $api.get<IContact>(`/admin/contacts/${route.params.contactId}`);
  contact.value = data;
  form.value = {
    status: data.status,
  };
}

onMounted(() => {
  fetchContact();
});
</script>
