<template>
  <div>
    <Menu :items="menuItems" />

    <ThreadTable title="閲覧履歴" :items="historyItems" moreLink="/history" link="/hoge" :maxItems="3" />
    <ThreadTable title="人気" :items="newsItems" moreLink="/news" :maxItems="5" />
    <ThreadTable title="新着" :items="popularItems" moreLink="/popular" :maxItems="3" />
  </div>
</template>

<script setup lang="ts">
import Menu from '~/components/Menu.vue';
import Header from '~/components/Header.vue';
import ThreadTable from '~/components/thread/ThreadTable.vue';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const menuItems = [
  { title: 'お知らせ', navigate: () => router.push('/'), icon: 'mdi-update' },
  { title: 'スレ一覧', navigate: () => router.push('/threads'), icon: 'mdi-new-box' },
  { title: '板一覧', navigate: () => router.push('/boards'), icon: 'mdi-format-list-bulleted' },
  { title: '設定', navigate: () => router.push('/setting'), icon: 'mdi-cog' },
  { title: 'サインイン', navigate: () => router.push('/signin'), icon: 'mdi-login' },
  { title: 'サインアップ', navigate: () => router.push('/signup'), icon: 'mdi-account-plus' },
  { title: 'スレ作成', navigate: () => router.push('/threads/new'), icon: 'mdi-forum' },
  { title: '板作成', navigate: () => router.push('/boards/new'), icon: 'mdi-forum' },
  { title: '管理画面', navigate: () => router.push('/admin'), icon: 'mdi-forum' },
];
const newsItems = ref([
  {
    id: 1,
    title:
      'ラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴えラーメン店主異例の訴え',
    subtitle: '食事中にイヤホンやめて',
    link: '/news/1',
    comments: 12,
    board: 'ニュース',
  },
  {
    id: 1,
    title: '都知事選',
    subtitle: '蓮舫氏、神谷外務次官の是非問う',
    link: '/news/2',
    comments: 8,
    board: 'ニュース',
  },
  {
    id: 1,
    title: 'ニコニコ動画',
    subtitle: '流出情報「ダウンロードや拡散はお控えください」',
    link: '/news/3',
    comments: 5,
    board: 'ニュース',
  },
]);

const popularItems = ref([
  { id: 1, title: 'こいせん7 全レス転載禁止', link: '/popular/1', comments: 30, board: '人気' },
  { id: 1, title: 'とらせん8', link: '/popular/2', comments: 25, board: '人気' },
  { id: 1, title: '巨専9', link: '/popular/3', comments: 20, board: '人気' },
  { id: 1, title: 'こいせん祝勝会 全レス転載禁止', link: '/popular/4', comments: 15, board: '人気' },
  { id: 1, title: 'こいせん4 全レス転載禁止', link: '/popular/5', comments: 10, board: '人気' },
]);

const historyItems = ref([
  { id: 0, title: 'スレッドA', comments: 3, board: '閲覧履歴' },
  { id: 1, title: 'スレッドB', comments: 4, board: '閲覧履歴' },
  { id: 2, title: 'スレッドC', comments: 2, board: '閲覧履歴' },
]);
</script>
