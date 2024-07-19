<template>
  <div>
    <PageTitle title="ユーザー" />

    <Menu :items="menuItems" />

    <br />

    <Form @submit="update" :validation-schema="schema" class="mx-2 mb-2" v-slot="{ meta }">
      <div class="field">
        <Field name="name" v-model="form.name" v-slot="{ errors }">
          <v-text-field
            v-model="form.name"
            label="名前(コメントの表示名になります)"
            variant="outlined"
            counter
            single-line
            dense
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="email" v-model="form.email" v-slot="{ errors }">
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

      <div class="field">
        <Field name="profileLink" v-model="form.profileLink" v-slot="{ errors }">
          <v-text-field
            v-model="form.profileLink"
            label="プロフィールリンク"
            type="url"
            variant="outlined"
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <v-btn type="submit" color="primary" block :disabled="!meta.valid">更新</v-btn>
    </Form>

    <v-divider class="border-opacity-100" />

    <br />

    <Form @submit="updatePassword" :validation-schema="passwordSchema" class="mx-2 mb-2" v-slot="{ meta }">
      <div class="field">
        <Field v-model="form.oldPassword" name="oldPassword" v-slot="{ errors }">
          <v-text-field
            v-model="form.oldPassword"
            label="現在のパスワード"
            type="password"
            variant="outlined"
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <div class="field">
        <Field v-model="form.newPassword" name="newPassword" v-slot="{ errors }">
          <v-text-field
            v-model="form.newPassword"
            label="新しいパスワード"
            type="password"
            variant="outlined"
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <v-btn type="submit" color="primary" block :disabled="!meta.valid">パスワード更新</v-btn>
    </Form>
  </div>
</template>

<script setup lang="ts">
import { Form, Field } from 'vee-validate';
import PageTitle from '~/components/PageTitle.vue';
import * as yup from 'yup';
import type { IThread } from '~/types/thread';
import type { IThreadComment } from '~/types/thread-comment';
import type { IListResource } from '~/types/list-resource';

const menuItems = [
  { title: 'ユーザー情報', clicked: () => router.push('/users/me'), icon: 'mdi-update' },
  { title: 'マイスレ', clicked: () => router.push('/users/me/threads'), icon: 'mdi-new-box' },
  { title: 'マイレス', clicked: () => router.push('/users/me/comments'), icon: 'mdi-format-list-bulleted' },
];

interface IUser {
  id: number;
  name: string;
  role: string;
  email: string;
  avatarUrl: string;
  profileLink: string;
  createdAt: string;
  updatedAt: string;
  threads: IListResource<IThread>;
  comments: IListResource<IThreadComment>;
}

const schema = yup.object({
  name: yup.string().required('必須項目です'),
  email: yup.string().email('有効なメールアドレスを入力してください').required('必須項目です'),
  profileLink: yup.string().url('有効なURLを入力してください').nullable(),
});

const passwordSchema = yup.object({
  oldPassword: yup.string().min(6, '6文字以上で入力してください').required('必須項目です'),
  newPassword: yup.string().min(6, '6文字以上で入力してください').required('必須項目です'),
});

const router = useRouter();
const route = useRoute();
const nuxtApp = useNuxtApp();
const { payload, $api } = nuxtApp;

const user = ref<IUser>({});

const form = ref({
  name: user.value.name,
  email: user.value.email,
  profileLink: user.value.profileLink || null,
  oldPassword: '',
  newPassword: '',
});

interface IUser {
  id: number;
  name: string;
  role: string;
  email: string;
  avatarUrl: string;
  profileLink: string;
  createdAt: string;
  updatedAt: string;
  threads: IListResource<IThread>;
  comments: IListResource<IThreadComment>;
}

onMounted(async () => {
  await fetchUser();
});

async function fetchUser() {
  const { data } = await $api.get<IUser>(`/users/${payload?.user?.id}`, {});
  user.value = data;
  form.value.name = data.name;
  form.value.email = data.email;
  form.value.profileLink = data.profileLink;
}

async function update() {
  if (confirm('ユーザー情報を更新しますか？')) {
    try {
      // 空文字列をnullに変換
      if (form.value.profileLink === '') {
        form.value.profileLink = null;
      }

      await $api.put('/users/me', {
        name: form.value.name,
        email: form.value.email,
        profileLink: form.value.profileLink,
      });
      router.go(0);
    } catch (err) {
      alert(err.response.data.error);
    }
  }
}
async function updatePassword() {
  if (confirm('パスワードを更新しますか？')) {
    try {
      const response = await $api.patch('/users/me/password', {
        oldPassword: form.value.oldPassword,
        newPassword: form.value.newPassword,
      });
      console.log(response);
      // router.go(0);
    } catch (err) {
      console.log(err);
      alert(err.response.data.error);
    }
  }
}
</script>
