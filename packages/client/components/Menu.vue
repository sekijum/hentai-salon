<template>
  <v-sheet class="outer-border" v-if="rows.length">
    <v-row align="center" justify="center" class="m-0" v-for="(row, rowIndex) in rows" :key="rowIndex">
      <v-col
        v-for="(item, idx) in row"
        :key="idx"
        :cols="columnWidth(row.length)"
        class="p-0"
        :class="{
          'no-right-border': shouldRemoveRightBorder(idx, row.length),
          'no-bottom-border': rowIndex === rows.length - 1,
        }"
      >
        <v-sheet class="menu-item" @click="item.clicked">
          <v-icon class="menu-icon">{{ item.icon }}</v-icon>
          <span class="menu-title">{{ item.title }}</span>
        </v-sheet>
      </v-col>
    </v-row>
  </v-sheet>
</template>

<script setup lang="ts">
const props = defineProps<{ items: { title: string; icon: string; clicked: () => void }[] }>();

const rows = computed(() => {
  const result = [];
  for (let i = 0; i < props.items.length; i += 3) {
    result.push(props.items.slice(i, i + 3));
  }
  return result;
});

function columnWidth(itemsPerRow: number) {
  if (itemsPerRow === 1) return 12;
  if (itemsPerRow === 2) return 6;
  return 4;
}

function shouldRemoveRightBorder(idx: number, itemsPerRow: number) {
  return (idx + 1) % itemsPerRow === 0;
}
</script>

<style scoped>
.outer-border {
  border: 1px solid #ccc;
}

.menu-item {
  align-items: center;
  box-shadow: none;
  width: 100%;
  margin: 0;
  padding: 16px;
  display: flex;
  justify-content: center;
  cursor: pointer;
  text-align: center;
}

.menu-icon {
  margin-right: 8px;
}

.menu-title {
  font-size: 12px;
  max-width: calc(100% - 32px);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.v-col {
  padding: 0 !important;
  border-bottom: 1px solid #ccc;
  border-right: 1px solid #ccc;
}

.v-col.no-right-border {
  border-right: none;
}

.v-col.no-bottom-border {
  border-bottom: none;
}

.v-row {
  margin: 0 !important;
}

.v-row:last-child .v-col {
  border-bottom: none;
}
</style>
