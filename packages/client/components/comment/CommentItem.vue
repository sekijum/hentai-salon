<template>
  <div class="comment-item">
    <v-sheet class="d-flex">
      <v-sheet class="me-auto">
        <template v-if="comment?.user?.id && comment?.user?.profileLink">
          <nuxt-link :to="comment.user.profileLink" target="_blank" rel="noopener">
            {{ username() }}
          </nuxt-link>
        </template>
        <template v-else>
          {{ username() }}
        </template>
      </v-sheet>
      <v-sheet>{{ $formatDate(comment.createdAt) }}</v-sheet>
    </v-sheet>

    <v-divider class="border-opacity-0"></v-divider>

    <v-sheet class="d-flex">
      <v-sheet class="me-auto">
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
      <v-sheet>
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

    <v-divider class="border-opacity-0"></v-divider>

    <OEmbedContent :text="comment.content" />

    <v-divider class="border-opacity-0"></v-divider>

    <template v-if="comment.attachments && comment.attachments.length">
      <v-row dense>
        <v-col cols="3" v-for="(attachment, idx) in comment.attachments" :key="idx" @click="openLightbox(idx)">
          <div class="media-item-wrapper">
            <v-img
              :src="`${attachment.url}?w=5`"
              :srcset="`${attachment.url}?w=5 5w, ${attachment.url}?w=5 5w`"
              class="media-item"
              contain
              referrerpolicy="no-referrer"
              :alt="comment.content"
            >
              <template v-slot:placeholder>
                <v-row align="center" class="fill-height ma-0" justify="center">
                  <v-progress-circular color="grey-lighten-5" indeterminate></v-progress-circular>
                </v-row>
              </template>
            </v-img>
          </div>
        </v-col>
      </v-row>
    </template>

    <v-divider class="border-opacity-0"></v-divider>

    <v-sheet class="d-flex">
      <v-sheet class="me-auto">
        <v-icon @click="toggleReplyForm">{{ showReplyForm ? 'mdi-close' : 'mdi-reply' }}</v-icon>
      </v-sheet>
      <v-sheet>
        <nuxt-link
          :to="`/threads/${threadId}/comments/${comment.id}`"
          class="text-decoration-none bg-transparent"
          target="_blank"
          rel="noopener"
        >
          <v-icon>mdi-comment</v-icon> {{ comment.replyCount }}
        </nuxt-link>
      </v-sheet>
      <v-sheet>
        <v-icon class="text-decoration-none ml-2" @click="toggleLike">
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

    <nuxt-link
      v-for="(item, index) in comment.attachments"
      :key="index"
      :href="item.url"
      :class="`glightbox-${comment.id}`"
      :data-title="`${comment.user?.name ? comment.user?.name : comment.guestName ? comment.guestName : '匿名'} ${
        comment?.createdAt
      }`"
      :data-description="comment.content"
      :data-type="item.type"
      data-effect="fade"
      data-zoomable="true"
      data-draggable="true"
      style="display: none"
    >
      <v-img :lazy-href="item.url" :src="item.url" alt="Image" aspect-ratio="1" class="bg-grey-lighten-2" cover>
        <template v-slot:placeholder>
          <v-row align="center" class="fill-height ma-0" justify="center">
            <v-progress-circular color="grey-lighten-5" indeterminate></v-progress-circular>
          </v-row>
        </template>
      </v-img>
    </nuxt-link>

    <v-divider></v-divider>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router';
import CommentForm from '~/components/comment/CommentForm.vue';
import OEmbedContent from '~/components/OEmbedContent.vue';
import type { IThreadComment } from '~/types/thread-comment';
import GLightbox from 'glightbox';

const props = defineProps<{ comment: IThreadComment; commentLimit: number; threadId: number }>();
const nuxtApp = useNuxtApp();
const route = useRoute();

const { $formatDate, $api, payload } = nuxtApp;

const showReplyForm = ref(false);
const emit = defineEmits(['replied']);
const isLiked = ref(props.comment.isLiked);

function toggleReplyForm() {
  showReplyForm.value = !showReplyForm.value;
}

function openLightbox(idx = 0) {
  const lightbox = GLightbox({
    selector: `.glightbox-${props.comment.id}`,
    touchNavigation: true,
    loop: true,
  });

  lightbox.openAt(idx);
}

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

<style scoped>
.comment-item {
  font-size: 12px;
}

.media-item-wrapper {
  position: relative;
  width: 100%;
  padding-bottom: 100%;
  background-color: black;
}

.media-item {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}
</style>
