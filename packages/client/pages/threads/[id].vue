<template>
  <div>
    <PageTitle :title="item.title" />

    <v-divider></v-divider>

    <v-chip-group active-class="primary--text" column>
      <v-chip
        small
        v-for="tag in [
          'Work',
          'Home Improvement',
          'Vacation',
          'Food',
          'Drawers',
          'Shopping',
          'Art',
          'Tech',
          'Creative Writing',
        ]"
        :key="tag"
      >
        {{ tag }}
      </v-chip>
    </v-chip-group>

    <v-divider></v-divider>

    <MenuSection :items="menuItems" />

    <br />

    <template v-if="tab === 'comments'">
      <CommentForm />

      <br />

      <v-divider></v-divider>

      <div id="comment-top" />
      <CommentList />
      <div id="comment-bottom" />

      <Pagination />

      <br />

      <CommentForm />

      <!-- 上にスクロールするFAB -->
      <v-btn icon large color="primary" class="fab fab-top" @click="scrollToCommentTop">
        <v-icon>mdi-arrow-up</v-icon>
      </v-btn>

      <!-- 下にスクロールするFAB -->
      <v-btn icon large color="primary" class="fab fab-bottom" @click="scrollToCommentBottom">
        <v-icon>mdi-arrow-down</v-icon>
      </v-btn>
    </template>
    <template v-else-if="tab === 'media'">
      <MediaGallery />
    </template>
  </div>
</template>

<script setup>
import { useRoute } from 'vue-router';
import Header from '~/components/Header.vue';
import ThreadTable from '~/components/thread/ThreadTable.vue';
import CommentList from '~/components/comment/CommentList.vue';
import CommentForm from '~/components/comment/CommentForm.vue';
import MenuSection from '~/components/MenuSection.vue';
import PageTitle from '~/components/PageTitle.vue';
import Pagination from '~/components/Pagination.vue';
import MediaGallery from '~/components/MediaGallery.vue';

const route = useRoute();
const keyword = ref(route.query.keyword ?? '');
const tab = ref(route.query.tab ?? 'comments');

const menuItems = [
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
];
const item = ref({
  title: 'ラーメン店主異例の訴え',
  subtitle: '食事中にイヤホンやめて',
  link: '/news/1',
  comments: 12,
  board: 'ニュース',
});

const scrollToCommentTop = () => {
  const commentTop = document.getElementById('comment-top');
  if (commentTop) {
    commentTop.scrollIntoView({ behavior: 'smooth' });
  }
};

const scrollToCommentBottom = () => {
  const commentBottom = document.getElementById('comment-bottom');
  if (commentBottom) {
    commentBottom.scrollIntoView({ behavior: 'smooth' });
  }
};
</script>

<style scoped>
.fab {
  position: fixed;
  right: 16px;
}

.fab-top {
  bottom: 72px; /* 下のFABとの間にスペースを確保 */
}

.fab-bottom {
  bottom: 16px;
}
</style>
