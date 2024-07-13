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
const props = defineProps<{ items: [] }>();

const router = useRouter();

const navigateTo = link => {
  router.push(link);
};

const rows = computed(() => {
  const result = [];
  for (let i = 0; i < props.items.length; i += 3) {
    result.push(props.items.slice(i, i + 3));
  }
  return result;
});

const columnWidth = itemsPerRow => {
  if (itemsPerRow === 1) return 12;
  if (itemsPerRow === 2) return 6;
  return 4; // default for 3 items
};

const shouldRemoveRightBorder = (idx, itemsPerRow) => {
  return (idx + 1) % itemsPerRow === 0;
};
</script>

<style scoped>
.outer-border {
  border: 1px solid #ccc; /* Outer border */
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
  border-bottom: 1px solid #ccc; /* Add bottom border to each column */
  border-right: 1px solid #ccc; /* Add right border to each column */
}

.v-col.no-right-border {
  border-right: none; /* Remove right border for specific columns */
}

.v-col.no-bottom-border {
  border-bottom: none; /* Remove bottom border for the last row */
}

.v-row {
  margin: 0 !important;
}

.v-row:last-child .v-col {
  border-bottom: none; /* Remove bottom border for columns in the last row */
}
</style>
