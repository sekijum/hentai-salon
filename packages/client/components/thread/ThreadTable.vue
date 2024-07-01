<template>
  <div>
    <div v-if="title" class="section-title">
      <h2>{{ title }}</h2>
    </div>

    <v-data-table :headers="headers" :items="items" hide-default-footer hide-default-header class="thread-section">
      <template v-slot:item="{ item, idx }">
        <div
          :class="{ alternate: idx % 2 === 0 }"
          @click="() => router.push(`/threads/${item.id}`)"
          class="d-flex align-center p-2 item-row"
        >
          <div class="fixed-image mr-1">
            <v-img :src="item.image || 'https://via.placeholder.com/80'" class="image"></v-img>
          </div>
          <div class="flex-grow-1">
            <p class="item-title">
              <strong>{{ truncateTitle(item.title) }}</strong>
            </p>
          </div>
          <div class="text-right mr-2">
            <small>
              2024/01/09
              <br />
              {{ item.board }}
              <br />
              話題度
              <v-icon small>mdi-comment</v-icon>
              <br />
              {{ item.comments }}
              <v-icon small>mdi-comment</v-icon>
              <br />
            </small>
          </div>
        </div>
      </template>
    </v-data-table>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';

defineProps({
  title: String,
  items: Array,
});

const router = useRouter();

const truncateTitle = title => {
  return title.length > 50 ? title.slice(0, 50) + '...' : title;
};

const headers = [
  { text: '', value: 'image', width: '20%' },
  { text: 'Title', value: 'title', width: '70%' },
  { text: 'Comments/Board', value: 'commentsBoard', width: '10%' },
];
</script>

<style scoped>
.thread-section {
  cursor: pointer;
}

.section-title h2 {
  color: orange;
  font-weight: bold;
}

.alternate {
  background-color: #f5f5f5;
}

.fixed-image {
  width: 100px;
  height: 100px;
  flex-shrink: 0;
}

.fixed-image .image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
</style>
