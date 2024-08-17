<template>
  <div class="comment-item mx-2">
    <v-sheet class="d-flex bg-transparent">
      <v-sheet class="me-auto bg-transparent">
        <template v-if="comment?.user?.id && comment?.user?.profileLink">
          <nuxt-link :to="comment.user.profileLink" target="_blank" rel="noopener">
            {{ username() }}
          </nuxt-link>
        </template>
        <template v-else>
          {{ username() }}
        </template>
      </v-sheet>
      <v-sheet class="bg-transparent">{{ $formatDate(comment.createdAt) }}</v-sheet>
    </v-sheet>

    <v-sheet class="d-flex bg-transparent">
      <v-sheet class="me-auto bg-transparent">
        <template v-if="comment.parentCommentId">
          <nuxt-link
            class="text-decoration-none bg-transparent"
            target="_blank"
            rel="noopener"
            :to="`/threads/${threadId}/comments/${comment.parentCommentId}`"
          >
            <v-chip size="x-small">>> {{ comment.parentCommentId }}</v-chip>
          </nuxt-link>
        </template>
      </v-sheet>
      <v-sheet class="bg-transparent">
        <nuxt-link
          :to="`/threads/${threadId}/comments/${comment.id}`"
          class="text-decoration-none bg-transparent"
          target="_blank"
          rel="noopener"
        >
          ID: {{ comment.id }}
        </nuxt-link>
      </v-sheet>
    </v-sheet>

    <CommentContent :text="comment.content" />

    <template v-if="comment.attachments && comment.attachments.length">
      <v-row :id="`gallery-${comment.id}`" class="pa-0 ma-0">
        <v-col v-for="(attachment, key) in comment.attachments" :key="key" class="d-flex child-flex pa-0 ma-0" cols="3">
          <nuxt-link
            :class="`gallery-item-${comment.id}`"
            :href="attachment.url"
            :data-pswp-width="getWidthFromUrl(attachment.url)"
            :data-pswp-height="getHeightFromUrl(attachment.url)"
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
              <p>{{ comment.content }}</p>
            </span>
          </nuxt-link>
        </v-col>
      </v-row>
    </template>

    <v-sheet class="d-flex bg-transparent">
      <v-sheet class="me-auto bg-transparent">
        <v-icon size="large" @click="toggleReplyForm">{{ showReplyForm ? 'mdi-close' : 'mdi-reply' }}</v-icon>
      </v-sheet>
      <v-sheet class="bg-transparent">
        <nuxt-link
          :to="`/threads/${threadId}/comments/${comment.id}`"
          class="text-decoration-none bg-transparent"
          target="_blank"
          rel="noopener"
        >
          <v-icon size="large">mdi-comment</v-icon> {{ comment.replyCount }}
        </nuxt-link>
      </v-sheet>
      <v-sheet class="bg-transparent">
        <v-icon size="large" class="text-decoration-none ml-2" @click="toggleLike">
          {{ isLiked ? 'mdi-thumb-up' : 'mdi-thumb-up-outline' }}
        </v-icon>
      </v-sheet>
    </v-sheet>

    <template v-if="showReplyForm">
      <CommentForm
        :threadId="threadId"
        :title="'返信 >> ' + comment.id"
        :parentCommentId="comment.id"
        @submit="
          () => {
            toggleReplyForm();
            emit('replied');
          }
        "
        :showReplyForm="showReplyForm"
      />
    </template>
  </div>
  <v-divider class="border-opacity-75"></v-divider>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router';
import CommentForm from '~/components/comment/CommentForm.vue';
import CommentContent from '~/components/comment/CommentContent.vue';
import type { IThreadComment } from '~/types/thread-comment';
import PhotoSwipeLightbox from 'photoswipe/lightbox';
import 'photoswipe/style.css';

const props = defineProps<{ comment: IThreadComment; commentLimit: number; threadId: number }>();
const nuxtApp = useNuxtApp();
const route = useRoute();
const lightbox = ref<PhotoSwipeLightbox | null>();

const { $formatDate, $api, payload } = nuxtApp;

const showReplyForm = ref(false);
const emit = defineEmits(['replied']);
const isLiked = ref(props.comment.isLiked);

function toggleReplyForm() {
  showReplyForm.value = !showReplyForm.value;
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

onMounted(() => {
  lightbox.value = new PhotoSwipeLightbox({
    gallery: `#gallery-${props.comment.id}`,
    children: `.gallery-item-${props.comment.id}`,
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

function username() {
  const name = props.comment?.user?.name || props.comment?.guestName || '匿名';
  return `${name}`;
}

async function toggleLike() {
  if (!payload.isLoggedIn) {
    alert('ログインしてください。');
    return;
  }
  if (isLiked.value) {
    await $api.post(`/threads/${props.threadId}/comments/${props.comment.id}/unlike`);
    isLiked.value = false;
  } else {
    await $api.post(`/threads/${props.threadId}/comments/${props.comment.id}/like`);
    isLiked.value = true;
  }
}
</script>

<style>
.comment-item {
  font-size: 12px;
}

.v-col {
  padding: 1px !important;
  cursor: pointer;
}

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
