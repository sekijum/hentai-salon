<template>
  <div>
    <v-row class="pa-0 ma-0">
      <v-col v-for="item in attachments" :key="item.commentId" class="d-flex child-flex pa-0 ma-0" cols="6">
        <v-img
          v-if="item.type !== 'Video'"
          :lazy-src="item.url"
          :src="item.url"
          aspect-ratio="1"
          class="bg-grey-lighten-2"
          cover
          @click="openModalMedia(item)"
          referrerpolicy="no-referrer"
        >
          <template v-slot:placeholder>
            <v-row align="center" class="fill-height ma-0" justify="center">
              <v-progress-circular color="grey-lighten-5" indeterminate></v-progress-circular>
            </v-row>
          </template>
        </v-img>
        <div v-else class="video-container" @click="openModalMedia(item)">
          <video
            ref="video"
            :src="item.url"
            @loadedmetadata="updateVideoMeta(item)"
            muted
            class="video-thumbnail"
            referrerpolicy="no-referrer"
          />
          <v-icon class="play-icon">mdi-play-circle</v-icon>
          <div v-if="item.duration" class="time-label">{{ formatTime(item.duration) }}</div>
        </div>
      </v-col>
    </v-row>

    <MediaModal
      v-if="selectedAttachment"
      :to="selectedAttachmentMeta?.to"
      :type="selectedAttachment.type"
      :url="selectedAttachment.url"
      @close="closeModalMedia"
    />
  </div>
</template>

<script setup lang="ts">
import MediaModal from '~/components/MediaModal.vue';
import type { IThreadCommentAttachmentForThread, IThreadCommentAttachment } from '~/types/thread-comment-attachment';

const props = defineProps<{
  attachments: IThreadCommentAttachmentForThread[];
  commentLimit: number;
  threadId: number;
}>();

const route = useRoute();

const selectedAttachment = ref<IThreadCommentAttachment | null>();
const selectedAttachmentMeta = ref<{ to: string } | null>();

function openModalMedia(attachment: IThreadCommentAttachmentForThread) {
  const limit = route.query.limit ? parseInt(route.query.limit as string, 10) : props.commentLimit;
  const newOffset = Math.floor((attachment.idx - 1) / limit) * limit;
  selectedAttachmentMeta.value = {
    to: `/threads/${props.threadId}/comments/${attachment.commentId}`,
  };
  selectedAttachment.value = { url: attachment.url, type: attachment.type, displayOrder: attachment.displayOrder };
}

function closeModalMedia() {
  selectedAttachmentMeta.value = null;
  selectedAttachment.value = null;
}

function formatTime(seconds: number): string {
  const minutes = Math.floor(seconds / 60);
  const remainingSeconds = Math.floor(seconds % 60);
  return `${minutes.toString().padStart(2, '0')}:${remainingSeconds.toString().padStart(2, '0')}`;
}

function updateVideoMeta(item: IThreadCommentAttachmentForThread) {
  const videoElement = document.createElement('video');
  videoElement.src = item.url;
  videoElement.onloadedmetadata = () => {
    item.duration = videoElement.duration;
  };
}
</script>

<style scoped>
.v-col {
  padding: 2px !important;
  cursor: pointer;
}

.video-container {
  position: relative;
  width: 100%;
  padding-bottom: 100%;
}

.video-thumbnail {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.play-icon {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  color: white;
  font-size: 60px;
}

.time-label {
  position: absolute;
  bottom: 8px;
  right: 8px;
  background-color: rgba(0, 0, 0, 0.7);
  color: white;
  padding: 2px 4px;
  border-radius: 3px;
  font-size: 14px;
}
</style>
