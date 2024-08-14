<template>
  <v-row id="gallery" class="pa-0 ma-0">
    <v-col v-for="(attachment, key) in attachments" :key="key" class="d-flex child-flex pa-0 ma-0" cols="3">
      <nuxt-link
        class="gallery-item"
        :href="attachment.url"
        data-pswp-width="1080"
        data-pswp-height="1080"
        style="height: 100%; width: 100%"
      >
        <v-img :lazy-href="attachment.url" :src="attachment.url" aspect-ratio="1" class="bg-grey-lighten-2" cover>
          <template v-slot:placeholder>
            <v-row align="center" class="fill-height ma-0" justify="center">
              <v-progress-circular color="grey-lighten-5" indeterminate></v-progress-circular>
            </v-row>
          </template>
        </v-img>
        <span class="hidden-caption-content" style="display: none">
          <p>{{ attachment.commentContent }}</p>
        </span>
      </nuxt-link>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import type { IThreadCommentAttachmentForThread } from '~/types/thread-comment-attachment';
import PhotoSwipeLightbox from 'photoswipe/lightbox';
import 'photoswipe/style.css';

const props = defineProps<{
  attachments: IThreadCommentAttachmentForThread[];
}>();

const route = useRoute();

const lightbox = ref<PhotoSwipeLightbox | null>();

onMounted(() => {
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
.v-col {
  padding: 1px !important;
  cursor: pointer;
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
