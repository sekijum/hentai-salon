<template>
  <div v-html="embedHtml" />
</template>

<script setup lang="ts">
const props = defineProps<{ text: string }>();
const embedHtml = ref('');

function convertTextToHtml(text: string): string {
  let html = text.replace(/(https?:\/\/[^\s]+)/g, '<a href="$1" target="_blank" rel="noopener noreferrer">$1</a>');
  html = html.replace(/\n/g, '<br>');
  return html;
}

onMounted(() => {
  const htmlContent = convertTextToHtml(props.text);
  embedHtml.value = htmlContent;
});
</script>
