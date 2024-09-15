<template>
  <v-container fluid>
    <v-breadcrumbs
      :items="[
        {
          title: 'ダッシュボード',
          disabled: false,
          href: '/admin',
        },
        {
          title: '広告一覧',
          disabled: false,
          href: '/admin/ads',
        },
        {
          title: '広告編集',
          disabled: true,
        },
      ]"
    />

    <Form @submit="submit" :validation-schema="schema" v-slot="{ meta, errors }">
      <div class="field mb-2">
        <Field name="content" v-model="form.content" v-slot="{ errors }">
          <v-textarea
            v-model="form.content"
            label="内容"
            variant="outlined"
            density="compact"
            :error-messages="errors"
          />
        </Field>
      </div>

      <div class="field mb-2">
        <Field v-model="form.isActive" name="isActive" v-slot="{ errors }">
          <v-select
            v-model="form.isActive"
            :items="statusList"
            :item-props="statusProps"
            label="アクティブ"
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

    <template v-if="form.content">
      <v-sheet class="mt-6" elevation="2" outlined rounded>
        <h3>プレビュー</h3>

        <Ad :content="form.content" />
      </v-sheet>
    </template>
  </v-container>
</template>

<script setup lang="ts">
import { Form, Field } from 'vee-validate';
import * as yup from 'yup';
import type { IAd } from '~/types/ad';
import Ad from '~/components/Ad.vue';

definePageMeta({
  layout: 'admin',
  middleware: ['admin-access-only'],
});

const router = useRouter();
const route = useRoute();
const { $api, $formatDate } = useNuxtApp();
const ad = ref<IAd>();

const form = ref({
  content: '',
  isActive: 1,
});

const statusList = [
  { text: 'アクティブ', value: 1, subtitle: '現在この広告は有効です' },
  { text: '非アクティブ', value: 0, subtitle: '現在この広告は無効です' },
];

const iframeContent = ref<HTMLIFrameElement>();

function statusProps(item: (typeof statusList)[0]) {
  return {
    title: item.text,
    subtitle: item.subtitle,
  };
}

const schema = yup.object({
  content: yup.string().required('必須項目です'),
});

async function submit() {
  try {
    await $api.put(`/admin/ads/${ad!.value!.id}`, form.value);
    router.push('/admin/ads');
  } catch (err) {
    alert('通信中にエラーが発生しました');
  }
}

async function fetchAd() {
  const { data } = await $api.get<IAd>(`/admin/ads/${route.params.adId}`);
  console.log(data);
  ad.value = data;
  form.value = {
    content: data.content,
    isActive: data.isActive,
  };
}

onMounted(async () => {
  await fetchAd();
});
</script>
