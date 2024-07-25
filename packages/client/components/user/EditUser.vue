<template>
  <div>
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

      <v-btn type="submit" color="primary" block :disabled="!meta.valid">パスワード更新</v-btn>
    </Form>

    <v-divider class="border-opacity-100" />

    <v-btn type="submit" color="error" block class="mt-5">退会</v-btn>
  </div>
</template>

<script setup lang="ts">
import { Form, Field } from 'vee-validate';
import * as yup from 'yup';
import type { IThread } from '~/types/thread';
import type { IThreadComment } from '~/types/thread-comment';
import type { IListResource } from '~/types/list-resource';

interface IUser {
  id: number;
  name: string;
  role: string;
  email: string;
  profileLink: string;
  createdAt: string;
  updatedAt: string;
  threads: IListResource<IThread>;
  comments: IListResource<IThreadComment>;
}

const props = defineProps<{ user: IUser }>();

const schema = yup.object({
  name: yup.string().required('必須項目です'),
  email: yup.string().email('有効なメールアドレスを入力してください').required('必須項目です'),
  profileLink: yup.string().url('有効なURLを入力してください').nullable(),
});

const passwordSchema = yup.object({
  password: yup.string().min(6, '6文字以上で入力してください').required('必須項目です'),
  confirmPassword: yup
    .string()
    .oneOf([yup.ref('password')], 'パスワードが一致しません')
    .required('必須項目です'),
});

const form = ref({
  name: props.user.name,
  email: props.user.email,
  profileLink: props.user.profileLink,
  password: '',
});

interface IUser {
  id: number;
  name: string;
  role: string;
  email: string;
  profileLink: string;
  createdAt: string;
  updatedAt: string;
  threads: IListResource<IThread>;
  comments: IListResource<IThreadComment>;
}

async function update() {}
async function updatePassword() {}
</script>
