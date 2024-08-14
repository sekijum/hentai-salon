<template>
  <div class="mt-2">
    <v-row>
      <v-col cols="3" class="d-flex justify-start align-center ml-3">
        <v-sheet class="menu-item" @click="router.back()">
          <v-icon class="menu-icon">mdi-arrow-left</v-icon>
          <span class="menu-title">戻る</span>
        </v-sheet>
      </v-col>
      <v-col class="d-flex justify-center align-center">
        <h1 class="font-weight-regular header-title" @click="router.push('/')">変態サロン</h1>
      </v-col>
      <v-col cols="3" class="d-flex justify-end align-center mr-3">
        <v-sheet class="menu-item" @click="openMenuModal">
          <v-icon class="menu-icon">mdi-cog</v-icon>
          <span class="menu-title">メニュー</span>
        </v-sheet>
      </v-col>
    </v-row>
    <div class="d-flex justify-center">
      <span class="subtitle">変態に特化したサロン。</span>
    </div>
    <v-divider />
  </div>
</template>

<script setup lang="ts">
const router = useRouter();

const isMenuModal = useState('isMenuModal', () => false);

function openMenuModal() {
  isMenuModal.value = true;
}
</script>

<style scoped>
.menu-item {
  display: flex;
  align-items: center;
  cursor: pointer;
  border: none;
  box-shadow: none;
}

.header-title {
  cursor: pointer;
}

.subtitle {
  font-size: 12px;
}
</style>
