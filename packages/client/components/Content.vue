<template>
  <p v-html="sanitizedText" />
</template>

<script setup lang="ts">
import linkifyHtml from 'linkify-html';
import DOMPurify from 'dompurify';

const props = defineProps<{ text: string }>();

const sanitizedText = computed(() => {
  const linkified = linkifyHtml(props.text);
  const linkifiedWithBreaks = linkified.replace(/\n/g, '<br>');
  return DOMPurify.sanitize(linkifiedWithBreaks);
});
</script>
