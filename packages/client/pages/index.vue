<template>
  <div>
    <v-alert
      v-if="payload.isLoggedIn"
      :text="`${payload?.user?.name}さん、あなたのアカウントにログインしています。`"
      prominent
      class="small-text text-center"
    />

    <Menu :items="guestMenuItems" />

    <template v-if="commentAttachments?.length">
      <swiper
        id="gallery"
        :slidesPerView="2"
        :spaceBetween="10"
        :grid="{ rows: 3 }"
        :modules="modules"
        :autoplay="{
          delay: 2500,
          disableOnInteraction: false,
        }"
        class="swiper"
      >
        <swiper-slide class="swiper-slide" v-for="(attachment, idx) in commentAttachments">
          <nuxt-link
            class="gallery-item"
            :href="attachment.url"
            :data-pswp-width="getWidthFromUrl(attachment.url)"
            :data-pswp-height="getHeightFromUrl(attachment.url)"
            :alt="attachment.commentId"
          >
            <v-img
              :lazy-href="attachment.url"
              :src="attachment.url"
              aspect-ratio="1"
              class="bg-grey-lighten-2"
              cover
              :alt="attachment.commentId.toString()"
            >
              <template v-slot:placeholder>
                <v-row align="center" class="fill-height ma-0" justify="center">
                  <v-progress-circular color="grey-lighten-5" indeterminate></v-progress-circular>
                </v-row>
              </template>
            </v-img>
            <span class="hidden-caption-content" style="display: none">
              <nuxt-link
                :to="`/threads/${attachment.threadId}/comments/${attachment.commentId}`"
                target="_blank"
                rel="noopener"
              >
                {{ attachment.commentId }}
              </nuxt-link>
              <p>{{ attachment.commentContent }}</p>
            </span>
          </nuxt-link>
        </swiper-slide>
      </swiper>
    </template>

    <ThreadItem
      filter="history"
      title="スレッド閲覧履歴"
      :clicked="() => router.push({ path: '/threads', query: { filter: 'history' } })"
      :isInfiniteScroll="false"
      :limit="5"
    />
    <ThreadItem
      filter="popularity"
      title="人気"
      :clicked="() => router.push({ path: '/threads' })"
      :isInfiniteScroll="false"
    />
    <ThreadItem filter="newest" title="新着" :isInfiniteScroll="true" />
  </div>
</template>

<script setup lang="ts">
import Menu from '~/components/Menu.vue';
import ThreadItem from '~/components/thread/ThreadItem.vue';
import { Swiper, SwiperSlide } from 'swiper/vue';
import 'swiper/css';
import type { IThreadCommentAttachmentForThread } from '~/types/thread-comment-attachment';
import { Grid, Autoplay } from 'swiper/modules';
import 'swiper/css/grid';
import PhotoSwipeLightbox from 'photoswipe/lightbox';
import 'photoswipe/style.css';
import 'swiper/css/navigation';

const config = useRuntimeConfig();
const route = useRoute();
const router = useRouter();
const nuxtApp = useNuxtApp();
const { $api } = nuxtApp;
const { getThreadViewHistory } = useStorage();
const commentAttachments = ref<IThreadCommentAttachmentForThread[]>([]);
const modules = [Grid, Autoplay];

const { payload } = nuxtApp;

const guestMenuItems = computed(() => {
  let items: { title: string; clicked: () => any; icon: string }[] = [
    { title: 'スレ一覧', clicked: () => router.push('/threads'), icon: 'mdi-format-list-bulleted' },
    { title: '板一覧', clicked: () => router.push('/boards'), icon: 'mdi-view-list' },
    { title: '設定', clicked: () => router.push('/setting'), icon: 'mdi-cog' },
  ];

  if (!payload.isLoggedIn) {
    items = items.concat([
      { title: 'サインイン', clicked: () => router.push('/signin'), icon: 'mdi-login' },
      { title: 'サインアップ', clicked: () => router.push('/signup'), icon: 'mdi-account-plus' },
    ]);
  }

  if (payload.isMember || payload.isAdmin) {
    items = items.concat([
      { title: 'スレ作成', clicked: () => router.push('/threads/new'), icon: 'mdi-pencil' },
      { title: 'マイページ', clicked: () => router.push('/mypage'), icon: 'mdi-account' },
    ]);
  }

  if (payload.isAdmin) {
    items = items.concat([
      { title: '板作成', clicked: () => router.push('/boards/new'), icon: 'mdi-plus-box' },
      { title: '管理画面', clicked: () => router.push('/admin'), icon: 'mdi-shield-account' },
    ]);
  }

  return items;
});

function getWidthFromUrl(url: string): number {
  const parts = url.split('/');
  const width = parseInt(parts[parts.length - 2], 10);
  return isNaN(width) ? 1080 : width;
}

function getHeightFromUrl(url: string): number {
  const parts = url.split('/');
  const height = parseInt(parts[parts.length - 1], 10);
  return isNaN(height) ? 1080 : height;
}

async function fetchRelatedByHistoryComments() {
  const response = await $api.get<IThreadCommentAttachmentForThread[]>('/attachments', {
    params: {
      filter: 'related-by-history',
      threadIds: getThreadViewHistory(),
    },
  });
  commentAttachments.value = response.data;
}

const lightbox = ref<PhotoSwipeLightbox | null>();

onMounted(async () => {
  await fetchRelatedByHistoryComments();

  lightbox.value = new PhotoSwipeLightbox({
    gallery: '#gallery',
    children: '.gallery-item',
    pswpModule: () => import('photoswipe'),
    bgOpacity: 1,
    showHideAnimationType: 'zoom',
    spacing: 0.5,
  });

  lightbox.value.on('uiRegister', function () {
    lightbox.value!.pswp!.ui!.registerElement({
      name: 'custom-caption',
      order: 9,
      isButton: false,
      appendTo: 'root',
      html: 'Caption text',
      onInit: (el, pswp) => {
        lightbox.value!.pswp!.on('change', () => {
          const currSlideElement = lightbox.value!.pswp!.currSlide!.data.element;
          let captionHTML = '';
          if (currSlideElement) {
            const hiddenCaption = currSlideElement.querySelector('.hidden-caption-content');
            if (hiddenCaption) {
              captionHTML = hiddenCaption.innerHTML;
            }
          }
          el.innerHTML = captionHTML || '';
        });
      },
    });
  });
  lightbox.value.init();
});

useHead({
  title: '変態サロン',
  meta: [
    { name: 'description', content: '変態に特化したサロン。' },
    {
      property: 'og:title',
      content: '変態サロン',
    },
    {
      property: 'og:description',
      content: '変態に特化したサロン。',
    },
    {
      property: 'og:image',
      content: '/hentai-salon-logo/logo_transparent.png',
    },
    {
      property: 'og:url',
      content: location.href,
    },
  ],
});
</script>

<style scoped>
.small-text {
  font-size: 0.75rem;
}

.swiper {
  height: 100vw;
}

.swiper-slide {
  height: 50%;
}

.pswp__custom-caption {
  background: rgba(75, 150, 75, 0.75);
  font-size: 16px;
  color: #fff;
  width: calc(100% - 32px);
  max-width: 400px;
  padding: 2px 8px;
  border-radius: 4px;
  position: absolute;
  left: 50%;
  bottom: 16px;
  transform: translateX(-50%);
}
.pswp__custom-caption a {
  color: #fff;
  text-decoration: underline;
}
.hidden-caption-content {
  display: none;
}
</style>
