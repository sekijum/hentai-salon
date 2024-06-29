<template>
  <div>
    <div class="section-title">
      <h2>{{ title }}</h2>
    </div>
    <div v-for="(item, index) in items" :key="index" :class="['section-item', { alternate: index % 2 === 0 }]" @click="navigateTo(item.link)">
      <v-row>
        <v-col cols="2">
          <v-img src="https://via.placeholder.com/80" aspect-ratio="2"></v-img>
        </v-col>
        <v-col cols="7">
          <p class="item-title">
            <strong>{{ item.title }}</strong>
          </p>
        </v-col>
        <v-col cols="3" class="text-right comment-board">
          <small class="comment-info"> {{ item.comments }}<v-icon small>mdi-comment</v-icon> {{ item.board }} </small>
        </v-col>
      </v-row>
    </div>
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
</script>

<style scoped>
.section-title {
  display: flex;
  flex-direction: column;
}

.section-title h2 {
  color: orange;
  font-weight: bold;
  margin: 0;
  padding: 8px 0;
}

.section-item {
  width: 100%;
  cursor: pointer;
  padding: 8px 0;
}

.section-item.alternate {
  background-color: #f5f5f5; /* グレー */
}

.section-item:not(.alternate) {
  background-color: #ffffff; /* 白 */
}

.comment-board {
  font-size: 0.8em;
  color: gray;
  display: flex;
  justify-content: flex-end;
  align-items: flex-end;
}

.comment-info {
  display: flex;
  align-items: center;
  justify-content: flex-end;
}
</style>
