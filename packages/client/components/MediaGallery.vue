<template>
  <div>
    <v-infinite-scroll height="100%" :items="items" :onLoad="load">
      <template #empty>
        <!-- Empty slot to override the default "No more" message -->
        <div></div>
      </template>
      <v-row class="pa-0 ma-0">
        <v-col v-for="item in items" :key="item.commentId" class="d-flex child-flex pa-0 ma-0" cols="6">
          <v-img
            :lazy-src="item.url"
            :src="item.url"
            aspect-ratio="1"
            class="bg-grey-lighten-2"
            cover
            @click="openModalMedia(item)"
          >
            <template v-slot:placeholder>
              <v-row align="center" class="fill-height ma-0" justify="center">
                <v-progress-circular color="grey-lighten-5" indeterminate></v-progress-circular>
              </v-row>
            </template>
            <v-icon v-if="item.type === 'Video'" size="60" class="play-icon">mdi-play-circle</v-icon>
            <div v-if="item.type === 'Video'" class="time-label">{{ formatTime(item.displayOrder) }}</div>
          </v-img>
        </v-col>
      </v-row>
    </v-infinite-scroll>

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
import type { IThreadCommentAttachmentForThread, IThreadCommentAttachment } from '~/types/thread-comment-attachment';

const props = defineProps<{
  attachments: IThreadCommentAttachmentForThread[];
  commentLimit: number;
  threadId: number;
}>();

const route = useRoute();
const items = ref<IThreadCommentAttachmentForThread[]>(props.attachments.slice(0, 10));
const currentIndex = ref(10);

const selectedAttachment = ref<IThreadCommentAttachment | null>();
const selectedAttachmentMeta = ref<{ to: string } | null>();

watch(
  () => props.attachments,
  newAttachments => {
    items.value = newAttachments.slice(0, 10);
    currentIndex.value = 10;
  },
);

function openModalMedia(attachment: IThreadCommentAttachmentForThread) {
  const limit = route.query.limit ? parseInt(route.query.limit as string, 10) : props.commentLimit;
  const newOffset = Math.floor((attachment.idx - 1) / limit) * limit;
  selectedAttachmentMeta.value = {
    to: `/threads/${props.threadId}?offset=${newOffset}&limit=${limit}#comments-${attachment.idx}`,
  };
  selectedAttachment.value = { url: attachment.url, type: attachment.type, displayOrder: attachment.displayOrder };
}

function closeModalMedia() {
  selectedAttachmentMeta.value = null;
  selectedAttachment.value = null;
}

async function load({ done }: { done: (status: string) => void }) {
  if (currentIndex.value >= props.attachments.length) {
    done('empty');
    return;
  }

  const nextItems = props.attachments.slice(currentIndex.value, currentIndex.value + 10);
  items.value.push(...nextItems);
  currentIndex.value += 10;
  done('ok');
}

function navigateToExternalPath(path: string) {
  const domain = window.location.origin;
  window.open(`${domain}${path}`, '_blank');
}

function formatTime(seconds: number): string {
  const minutes = Math.floor(seconds / 60);
  const remainingSeconds = seconds % 60;
  return `${minutes}:${remainingSeconds.toString().padStart(2, '0')}`;
}
</script>

<style scoped>
.v-col {
  padding: 2px !important; /* 画像同士の間隔を狭める */
  cursor: pointer; /* マウスオーバー時にカーソルをポインターに変更 */
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
