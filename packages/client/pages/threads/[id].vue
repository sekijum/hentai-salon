<template>
  <div>
    <PageTitle :title="item.title" />

    <v-divider></v-divider>

    <v-chip-group active-class="primary--text" column>
      <v-chip small v-for="tag in ['Work', 'Home Improvement', 'Vacation', 'Food', 'Drawers', 'Shopping', 'Art', 'Tech', 'Creative Writing']" :key="tag">
        {{ tag }}
      </v-chip>
    </v-chip-group>

    <v-divider></v-divider>

    <MenuSection
      :items="[
        {
          title: 'コメント一覧',
          to: '',
          icon: 'mdi-fire',
        },
        {
          title: 'メディア',
          to: '',
          icon: 'mdi-update',
        },
      ]"
    />

    <CommentForm />

    <v-divider></v-divider>

    <CommentList />

    <Pagination />

    <CommentForm />
  </div>
</template>

<script setup>
import Header from '~/components/Header.vue';
import ThreadTable from '~/components/thread/ThreadTable.vue';
import CommentList from '~/components/comment/CommentList.vue';
import CommentForm from '~/components/comment/CommentForm.vue';
import MenuSection from '~/components/MenuSection.vue';
import PageTitle from '~/components/PageTitle.vue';
import Pagination from '~/components/Pagination.vue';

const route = useRoute();

const keyword = ref(route.query.keyword ?? '');

const item = ref({
  title:
    'ラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴え',
  subtitle: '食事中にイヤホンやめて',
  link: '/news/1',
  comments: 12,
  board: 'ニュース',
});
</script>

<style scoped>
.thread-title {
  font-size: 1rem;
  text-align: center;
}
</style>
