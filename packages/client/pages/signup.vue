<template>
  <div>
    <PageTitle title="サインアップ" />

    <v-divider />

    <Menu :items="menuItems" />

    <br />

    <Form @submit="submit" :validation-schema="schema" class="mx-2 mb-2" v-slot="{ meta, errors }">
      <div class="field">
        <Field name="name" v-slot="{ field, errorMessage }">
          <v-text-field
            v-model="form.name"
            v-bind="field"
            label="名前(コメントの表示名になります)"
            type="text"
            variant="outlined"
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="email" v-slot="{ field, errorMessage }">
          <v-text-field
            v-model="form.email"
            v-bind="field"
            label="メールアドレス"
            type="email"
            variant="outlined"
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="password" v-slot="{ field, errorMessage }">
          <v-text-field
            v-model="form.password"
            v-bind="field"
            label="パスワード"
            type="password"
            variant="outlined"
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="confirmPassword" v-slot="{ field, errorMessage }">
          <v-text-field
            v-bind="field"
            label="パスワード確認"
            type="password"
            variant="outlined"
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <div class="field">
        <Field name="profileLink" v-slot="{ field, errorMessage }">
          <v-text-field
            v-model="form.profileLink"
            v-bind="field"
            label="プロフィールリンク"
            type="url"
            variant="outlined"
            density="compact"
            :error-messages="errorMessage ? [errorMessage] : []"
          />
        </Field>
      </div>

      <!-- <div class="field">
        <v-file-input
          show-size
          truncate-length="25"
          prepend-icon=""
          label="アバターを選択"
          variant="outlined"
          density="compact"
          hide-details
          accept="image/*"
          @change="handleAvatarChange"
        />
      </div> -->

      <v-btn type="submit" color="primary" block :disabled="!meta.valid" class="mt-5">サインアップ</v-btn>
    </Form>
  </div>
</template>

<script setup lang="ts">
import { Form, Field } from 'vee-validate';
import PageTitle from '~/components/PageTitle.vue';
import Menu from '~/components/Menu.vue';
import * as yup from 'yup';

const nuxtApp = useNuxtApp();
const router = useRouter();
const { fetchListPresignedUrl, uploadFilesToS3 } = useActions();

const { $storage, $api } = nuxtApp;

const avatarFile = ref<File>();

const form = ref({
  name: '',
  email: '',
  password: '',
  avatarUrl: null as null | string,
  profileLink: '' as null | string,
});

const menuItems = [
  { title: 'サインイン', clicked: () => router.push('/signin'), icon: 'mdi-login' },
  { title: 'サインアップ', clicked: () => router.push('/signup'), icon: 'mdi-account-plus' },
];

const schema = yup.object({
  name: yup.string().required('必須項目です'),
  email: yup.string().email('有効なメールアドレスを入力してください').required('必須項目です'),
  password: yup.string().min(6, '6文字以上で入力してください').required('必須項目です'),
  confirmPassword: yup
    .string()
    .oneOf([yup.ref('password')], 'パスワードが一致しません')
    .required('必須項目です'),
  profileLink: yup.string().url('有効なURLを入力してください').nullable(),
});

function handleAvatarChange(event: Event) {
  const input = event.target as HTMLInputElement;
  if (input.files && input.files[0]) {
    avatarFile.value = input.files[0];
  }
}

async function submit() {
  try {
    // 空文字列をnullに変換
    if (form.value.profileLink === '') {
      form.value.profileLink = null;
    }

    if (avatarFile.value) {
      const presignedUrls = await fetchListPresignedUrl([avatarFile.value.name]);
      const thumbnailUrl = await uploadFilesToS3(presignedUrls[0], avatarFile.value);
      form.value.avatarUrl = thumbnailUrl;
    } else {
      form.value.avatarUrl = null;
    }
    const response = await $api.post('/signup', form.value);
    const authHeader = response.headers.authorization;
    const token = authHeader.split(' ')[1];
    $storage.setItem('access_token', token);
    alert('サインアップしました。');
    router.push('/');
  } catch (err) {
    alert(err.response.data.error);
  }
}
</script>
