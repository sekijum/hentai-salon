<template>
  <div>
    <span v-for="(chunk, index) in htmlChunks" :key="index">
      <template v-if="chunk.type === 'text'">{{ chunk.content }}</template>
      <template v-else-if="chunk.type === 'link'">
        <a :href="chunk.content" target="_blank" rel="noopener noreferrer">{{ chunk.content }}</a>
      </template>
      <template v-else-if="chunk.type === 'newline'"><br /></template>
    </span>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{ text: string }>();
const htmlChunks = ref<{ type: string; content: string }[]>([]);

function convertTextToChunks(text: string) {
  const chunks = [];
  const urlRegex = /(https?:\/\/[^\s]+)/g;
  const newlineRegex = /\n/g;
  let match;
  let lastIndex = 0;

  while ((match = urlRegex.exec(text)) !== null) {
    if (match.index > lastIndex) {
      const part = text.slice(lastIndex, match.index);
      const lines = part.split(newlineRegex);
      lines.forEach((line, index) => {
        chunks.push({ type: 'text', content: line });
        if (index < lines.length - 1) {
          chunks.push({ type: 'newline', content: '\n' });
        }
      });
    }

    chunks.push({ type: 'link', content: match[0] });
    lastIndex = match.index + match[0].length;
  }

  if (lastIndex < text.length) {
    const part = text.slice(lastIndex);
    const lines = part.split(newlineRegex);
    lines.forEach((line, index) => {
      chunks.push({ type: 'text', content: line });
      if (index < lines.length - 1) {
        chunks.push({ type: 'newline', content: '\n' });
      }
    });
  }

  return chunks;
}

htmlChunks.value = convertTextToChunks(props.text);
</script>
