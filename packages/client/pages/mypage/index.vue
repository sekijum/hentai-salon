<template>
  <div>
    <PageTitle title="ユーザー" />

    <Menu :items="menuItems" />

    <br />

    <Form @submit="update" :validation-schema="schema" class="mx-2" v-slot="{ meta }">
      <div class="field mb-2">
        <Field name="name" v-model="updateForm.name" v-slot="{ errors }">
          <v-text-field
            v-model="updateForm.name"
            label="名前(コメントの表示名になります)"
            variant="outlined"
            dense
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <div class="field mb-2">
        <Field name="email" v-model="updateForm.email" v-slot="{ errors }">
          <v-text-field
            v-model="updateForm.email"
            label="メールアドレス"
            type="email"
            variant="outlined"
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <div class="field mb-2">
        <Field name="profileLink" v-model="updateForm.profileLink" v-slot="{ errors }">
          <v-text-field
            v-model="updateForm.profileLink"
            label="プロフィールリンク"
            type="url"
            variant="outlined"
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <v-btn type="submit" color="primary" block :disabled="!meta.valid">編集</v-btn>
    </Form>

    <br />

    <v-divider class="border-opacity-100" />

    <br />

    <Form @submit="updatePassword" :validation-schema="passwordSchema" class="mx-2" v-slot="{ meta }">
      <div class="field mb-2">
        <Field v-model="updatePasswordForm.oldPassword" name="oldPassword" v-slot="{ errors }">
          <v-text-field
            v-model="updatePasswordForm.oldPassword"
            label="現在のパスワード"
            type="password"
            variant="outlined"
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <div class="field mb-2">
        <Field v-model="updatePasswordForm.newPassword" name="newPassword" v-slot="{ errors }">
          <v-text-field
            v-model="updatePasswordForm.newPassword"
            label="新しいパスワード"
            type="password"
            variant="outlined"
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <v-btn type="submit" color="primary" block :disabled="!meta.valid">パスワード編集</v-btn>
    </Form>

    <br />

    <v-divider class="border-opacity-100" />

    <br />

    <div class="mx-2">
      <v-btn block variant="outlined" color="red" @click="deleteUser">アカウント削除</v-btn>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Form, Field } from 'vee-validate';
import PageTitle from '~/components/PageTitle.vue';
import * as yup from 'yup';
import type { IUser } from '~/types/user';

definePageMeta({ middleware: ['logged-in-access-only'] });

const menuItems = [
  { title: 'マイスレ', clicked: () => router.push('/mypage/threads'), icon: 'mdi-file-document-multiple-outline' },
  { title: 'ユーザー情報', clicked: () => router.push('/mypage'), icon: 'mdi-account-cog-outline' },
  { title: 'マイレス', clicked: () => router.push('/mypage/comments'), icon: 'mdi-message-text-outline' },
  {
    title: 'お気に入りスレ',
    clicked: () => router.push('/mypage/liked-threads'),
    icon: 'mdi-star-box-multiple-outline',
  },
  { title: 'お気に入りレス', clicked: () => router.push('/mypage/liked-comments'), icon: 'mdi-message-star-outline' },
];

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
const { $storage } = useNuxtApp();

const user = ref<IUser>();

const updateForm = ref({
  name: '',
  email: '',
  profileLink: null as string | null,
});

const updatePasswordForm = ref({
  oldPassword: '',
  newPassword: '',
});

onMounted(async () => {
  await fetchUser();
});

async function fetchUser() {
  const { data } = await $api.get<IUser>(`/users/me`, {});
  user.value = data;
  updateForm.value.name = data.name;
  updateForm.value.email = data.email;
  updateForm.value.profileLink = data.profileLink;
}

async function update() {
  if (confirm('ユーザー情報を編集しますか？')) {
    try {
      // 空文字列をnullに変換
      if (updateForm.value.profileLink === '') {
        updateForm.value.profileLink = null;
      }

      await $api.put('/users/me', {
        name: updateForm.value.name,
        email: updateForm.value.email,
        profileLink: updateForm.value.profileLink,
      });
      fetchUser();
    } catch (err) {
      alert(err.response.data.error);
    }
  }
}

async function updatePassword(_, { resetForm }: { resetForm: () => void }) {
  if (confirm('パスワードを編集しますか？')) {
    try {
      await $api.patch('/users/me/password', {
        oldPassword: updatePasswordForm.value.oldPassword,
        newPassword: updatePasswordForm.value.newPassword,
      });
      fetchUser();
      resetForm();
    } catch (err) {
      alert('通信中にエラーが発生しました');
    }
  }
}

async function deleteUser() {
  if (confirm('パスワードを削除しますか？')) {
    try {
      await $api.delete('/users/me');
      $storage.removeItem('access_token');
      router.go(0);
    } catch (err) {
      alert('通信中にエラーが発生しました');
    }
  }
}
</script>
