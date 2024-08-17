<template>
  <template v-if="commentAttachments?.length">
    <swiper
      id="gallery"
      :slidesPerView="2"
      :spaceBetween="5"
      :grid="getGridConfig()"
      :modules="modules"
      :autoplay="{
        delay: 3000,
        disableOnInteraction: false,
      }"
      class="swiper"
      :style="{ height: swiperHeight }"
    >
      <swiper-slide class="swiper-slide" v-for="(attachment, idx) in commentAttachments">
        <nuxt-link
          class="gallery-item"
          :href="attachment.url"
          :data-pswp-width="getWidthFromUrl(attachment.url)"
          :data-pswp-height="getHeightFromUrl(attachment.url)"
          data-caption="画像のキャプション"
          data-author="画像の作者"
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
          <span class="hidden-caption-content">
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
</template>

<script setup lang="ts">
import { Swiper, SwiperSlide } from 'swiper/vue';
import 'swiper/css';
import type { IThreadCommentAttachmentForThread } from '~/types/thread-comment-attachment';
import { Grid, Autoplay } from 'swiper/modules';
import 'swiper/css/grid';
import PhotoSwipeLightbox from 'photoswipe/lightbox';
import 'photoswipe/style.css';
import 'swiper/css/navigation';

const props = defineProps<{ commentAttachments: IThreadCommentAttachmentForThread[] }>();

const nuxtApp = useNuxtApp();
const modules = [Grid, Autoplay];

function getGridConfig() {
  return { rows: Math.min(props.commentAttachments?.length, 3) };
}

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

const lightbox = ref<PhotoSwipeLightbox | null>();

function calculateSwiperHeight(rows: number): string {
  const rowHeight = 200;
  return `${rows * rowHeight}px`;
}

const swiperHeight = computed(() => {
  const rows = Math.min(props.commentAttachments.length, 3);
  return calculateSwiperHeight(rows);
});

onMounted(async () => {
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

onUnmounted(() => {
  if (lightbox.value) {
    lightbox.value.destroy();
    lightbox.value = null;
  }
});
</script>

<style>
.pswp__custom-caption {
  font-size: 16px;
  width: calc(100% - 32px);
  max-width: 400px;
  padding: 2px 8px;
  border-radius: 4px;
  position: absolute;
  left: 50%;
  bottom: 16px;
  transform: translateX(-50%);
}

.hidden-caption-content {
  display: none;
}
</style>
