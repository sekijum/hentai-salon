<template>
  <div>
    <PageTitle title="スレ一覧" />

    <v-divider />

    <Menu :items="menuItems" />

    <template v-if="route.query.filter === 'keyword'">
      <v-text-field
        v-model="keyword"
        label="スレを検索"
        variant="outlined"
        hide-details
        counter
        single-line
        clearable
        dense
        density="compact"
      />
      <v-btn type="submit" color="primary" block @click="searchKeyword">検索</v-btn>
    </template>

    <template v-if="route.query.filter === 'tags'">
      <div class="field">
        <v-autocomplete
          v-model="tagNameList"
          no-data-text="一致する結果が見つかりませんでした"
          chips
          small-chips
          label="タグ"
          variant="outlined"
          density="compact"
          hide-details
          clearable
          multiple
          single-line
          dense
          :items="tagSuggestions"
        />
      </div>
      <v-btn type="submit" color="primary" block @click="searchTagNameList">検索</v-btn>
      <v-card-text>
        <v-chip-group v-model="tagNameList" column multiple>
          <template v-for="tag in tagSuggestions" :key="tag">
            <v-chip size="x-small" :text="tag" :value="tag" variant="outlined" filter />
          </template>
        </v-chip-group>
      </v-card-text>
    </template>

    <ThreadItem
      v-if="!route.query.filter || route.query.filter === 'popularity'"
      filter="popularity"
      title="人気"
      :isInfiniteScroll="true"
    />
    <ThreadItem v-if="route.query.filter === 'history'" filter="history" title="閲覧履歴" :isInfiniteScroll="false" />
    <ThreadItem v-if="route.query.filter === 'newest'" filter="newest" title="新着" :isInfiniteScroll="true" />
    <ThreadItem v-if="route.query.filter === 'keyword'" filter="keyword" title="全板検索" :isInfiniteScroll="true" />
    <ThreadItem
      v-if="route.query.filter === 'related-by-history'"
      filter="related-by-history"
      title="関連"
      :isInfiniteScroll="true"
    />
    <ThreadItem v-if="route.query.filter === 'board'" filter="board" title="板" :isInfiniteScroll="true" />
    <ThreadItem v-if="route.query.filter === 'tags'" filter="tags" title="タグ" :isInfiniteScroll="true" />
  </div>
</template>

<script setup lang="ts">
import ThreadItem from '~/components/thread/ThreadItem.vue';
import Menu from '~/components/Menu.vue';
import PageTitle from '~/components/PageTitle.vue';

const route = useRoute();
const router = useRouter();
const nuxtApp = useNuxtApp();

const keyword = ref(route.query.keyword ?? '');
const tagNameList = ref<string[]>(route.query.tagNameList ?? []);
const tagSuggestions = ref<string[]>([]);
const { $api } = nuxtApp;

const menuItems = [
  {
    title: 'トップ',
    clicked: () => router.push({ query: {} }),
    icon: 'mdi-home',
  },
  {
    title: '関連',
    clicked: () => router.push({ query: { filter: 'related-by-history' } }),
    icon: 'mdi-link-variant',
  },
  {
    title: '人気',
    clicked: () => router.push({ query: { filter: 'popularity' } }),
    icon: 'mdi-fire',
  },
  {
    title: '履歴',
    clicked: () => router.push({ query: { filter: 'history' } }),
    icon: 'mdi-history',
  },
  {
    title: '新着',
    clicked: () => router.push({ query: { filter: 'newest' } }),
    icon: 'mdi-new-box',
  },
  {
    title: 'スレッド検索',
    clicked: () => router.push({ query: { filter: 'keyword' } }),
    icon: 'mdi-new-box',
  },
  {
    title: 'タグ検索',
    clicked: () => router.push({ query: { filter: 'tags' } }),
    icon: 'mdi-new-box',
  },
];

async function searchKeyword() {
  await router.push({ ...{ query: { ...{ keyword: keyword.value }, ...{ filter: 'keyword' } } } });
}

async function searchTagNameList() {
  await router.push({ ...{ query: { ...{ tagNameList: tagNameList.value }, ...{ filter: 'tags' } } } });
}

onMounted(async () => {
  const response = await $api.get<string[]>('/tags/name');
  tagSuggestions.value = response.data ? response.data : [];
});

useHead({
  title: '変態サロン | スレ一覧',
  meta: [
    {
      property: 'og:title',
      content: '変態サロン | スレ一覧',
    },
    {
      property: 'og:url',
      content: location.href,
    },
  ],
});
</script>
