<template>
  <v-container fluid v-if="user?.id">
    <v-breadcrumbs
      :items="[
        {
          title: 'ダッシュボード',
          disabled: false,
          href: '/admin',
        },
        {
          title: 'ユーザー一覧',
          disabled: false,
          href: '/admin/users',
        },
        {
          title: user.name,
          disabled: true,
        },
      ]"
    />

    <Form @submit="submit" :validation-schema="schema" v-slot="{ meta, errors }">
      <div class="field mb-2">
        <v-text-field v-model="user.id" label="ID" type="text" variant="outlined" density="compact" disabled />
      </div>

      <div class="field mb-2">
        <v-text-field v-model="user.createdAt" label="登録日" type="text" variant="outlined" density="compact" />
      </div>

      <div class="field mb-2">
        <Field v-model="form.name" name="name" v-slot="{ errors }">
          <v-text-field
            v-model="form.name"
            label="名前"
            type="text"
            variant="outlined"
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <div class="field mb-2">
        <Field v-model="form.email" name="email" v-slot="{ errors }">
          <v-text-field
            v-model="form.email"
            label="メールアドレス"
            type="email"
            variant="outlined"
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <div class="field mb-2">
        <Field v-model="form.role" name="role" v-slot="{ errors }">
          <v-select
            v-model="form.role"
            :items="roleList"
            :item-props="roleProps"
            label="権限"
            variant="outlined"
            density="compact"
            :item-title="'text'"
            :item-value="'value'"
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
  </v-container>
</template>

<script setup lang="ts">
import { Form, Field } from 'vee-validate';
import * as yup from 'yup';

definePageMeta({
  layout: 'admin',
  middleware: ['admin-access-only'],
});

interface IUser {
  id: number;
  name: string;
  role: number;
  roleLabel: string;
  status: number;
  statusLabel: string;
  email: string;
  profileLink: string;
  createdAt: string;
  updatedAt: string;
}

const router = useRouter();
const route = useRoute();
const { $api, $formatDate } = useNuxtApp();
const user = ref<IUser>();

const form = ref({
  name: '',
  email: '',
  role: 0,
  status: 0,
});

const roleList = [
  { text: '会員', value: 0, subtitle: 'システムの通常ユーザー' },
  { text: '管理者', value: 1, subtitle: 'システム全体の管理を行う' },
];

function roleProps(item: (typeof roleList)[0]) {
  return {
    title: item.text,
    subtitle: item.subtitle,
  };
}

const statusList = [
  { text: '有効', value: 0, subtitle: '現在アクティブなユーザー' },
  { text: '退会済', value: 1, subtitle: 'ユーザーが自主的に退会した状態' },
  { text: '凍結', value: 2, subtitle: 'アカウントが一時的に使用できない状態' },
  { text: '無効', value: 3, subtitle: 'アカウントが無効化されている状態' },
];

function statusProps(item: (typeof statusList)[0]) {
  return {
    title: item.text,
    subtitle: item.subtitle,
  };
}

const schema = yup.object({
  name: yup.string().required('必須項目です'),
  email: yup.string().email('有効なメールアドレスを入力してください').required('必須項目です'),
  role: yup.number().required('必須項目です'),
  status: yup.number().required('必須項目です'),
});

async function submit() {
  if (confirm('ユーザー情報が編集しますか？')) {
    try {
      await $api.put(`/admin/users/${user?.value?.id}`, form.value);
      await fetchUser();
    } catch (err) {
      alert('通信中にエラーが発生しました');
    }
  }
}

async function fetchUser() {
  const { data } = await $api.get<IUser>(`/admin/users/${route.params.userId}`);
  user.value = data;
  form.value = {
    name: data.name,
    email: data.email,
    role: data.role,
    status: data.status,
  };
}

onMounted(() => {
  fetchUser();
});
</script>
