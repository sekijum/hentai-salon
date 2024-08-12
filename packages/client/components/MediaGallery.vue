<template>
  <div v-if="attachments.length">
    <v-row class="pa-0 ma-0">
      <v-col v-for="(item, index) in attachments" :key="index" class="d-flex child-flex pa-0 ma-0" cols="6">
        <nuxt-link
          :href="item.url"
          class="glightbox"
          :data-title="`${item?.commentAuthorName ? item.commentAuthorName : '匿名'} ${item?.createdAt}`"
          :data-description="item.commentContent"
          :data-type="item.type"
          data-effect="fade"
          data-zoomable="true"
          data-draggable="true"
          style="height: 100%; width: 100%"
        >
          <v-img :lazy-href="item.url" :src="item.url" alt="Image" aspect-ratio="1" class="bg-grey-lighten-2" cover>
            <template v-slot:placeholder>
              <v-row align="center" class="fill-height ma-0" justify="center">
                <v-progress-circular color="grey-lighten-5" indeterminate></v-progress-circular>
              </v-row>
            </template>
          </v-img>
        </nuxt-link>
      </v-col>
    </v-row>
  </div>
</template>

<script setup lang="ts">
import type { IThreadCommentAttachmentForThread, IThreadCommentAttachment } from '~/types/thread-comment-attachment';
import GLightbox from 'glightbox';

const props = defineProps<{
  attachments: IThreadCommentAttachmentForThread[];
  commentLimit: number;
  threadId: number;
}>();

const route = useRoute();

onMounted(() => {
  nextTick(() => {
    GLightbox({
      touchNavigation: true,
      loop: true,
    });
  });
});
watch(
  () => props.attachments || route.fullPath,
  () => {
    nextTick(() => {
      GLightbox({
        touchNavigation: true,
        loop: true,
      });
    });
  },
);
</script>

<style scoped>
.v-col {
  padding: 2px !important;
  cursor: pointer;
}
</style>
