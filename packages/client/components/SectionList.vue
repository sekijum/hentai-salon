<template>
  <div>
    <div class="section-title">
      <h2>{{ title }}</h2>
    </div>

    <v-data-table :headers="headers" :items="items" hide-default-footer hide-default-header>
      <template v-slot:item="{ item, index }">
        <div :class="{ alternate: index % 2 === 0 }" @click="navigateTo(item.link)" class="d-flex align-center p-2">
          <div class="fixed-image mr-4">
            <v-img :src="item.image || 'https://via.placeholder.com/80'" class="image"></v-img>
          </div>
          <div class="flex-grow-1">
            <p class="item-title mb-0">
              <strong>{{ item.title }}</strong>
            </p>
          </div>
          <div class="text-right">
            <small>
              {{ item.comments }}
              <v-icon small>mdi-comment</v-icon>
              {{ item.board }}
            </small>
          </div>
        </div>
      </template>
    </v-data-table>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();

const navigateTo = link => {
  router.push(link);
};

defineProps({
  title: String,
  items: Array,
});

const headers = [
  { text: '', value: 'image', width: '20%' },
  { text: 'Title', value: 'title', width: '50%' },
  { text: 'Comments/Board', value: 'commentsBoard', width: '30%' },
];
</script>

<style scoped>
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
}

.fixed-image .image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
</style>
