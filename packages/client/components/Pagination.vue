<template>
  <div class="text-center">
    <v-pagination
      v-model="page"
      :length="totalPages"
      :total-visible="7"
      density="compact"
      rounded="lg"
      prev-icon="mdi-menu-left"
      next-icon="mdi-menu-right"
      show-select
      class="elevation-1"
      @update:modelValue="onPageChange"
    />
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  limit: number;
  totalCount: number;
}>();

const route = useRoute();
const router = useRouter();

const limit = ref(parseInt(route.query.limit as string) || props.limit);
const offset = ref(Number.isNaN(parseInt(route.query.offset as string)) ? 0 : parseInt(route.query.offset as string));
const page = ref(Math.floor(offset.value / limit.value) + 1);

const totalPages = computed(() => Math.ceil(props.totalCount / limit.value));

const onPageChange = (page: number) => {
  const newOffset = (page - 1) * limit.value;
  router.push({ query: { ...route.query, offset: newOffset, limit: limit.value } });
};
</script>
