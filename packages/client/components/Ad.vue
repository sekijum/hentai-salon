<template>
  <iframe
    ref="iframeContent"
    referrerpolicy="strict-origin-when-cross-origin"
    frameborder="0"
    scrolling="no"
    loading="lazy"
    style="vertical-align: bottom"
    @load="resizeIframe"
  />
</template>

<script setup lang="ts">
const props = defineProps<{ content: string }>();
const iframeContent = ref<HTMLIFrameElement>();

onMounted(async () => {
  renderIframeContent();
  resizeIframe();
});

function renderIframeContent() {
  if (iframeContent.value) {
    const iframe = iframeContent.value;
    const doc = iframe.contentDocument || iframe.contentWindow?.document;
    if (doc) {
      doc.open();
      doc.write(`
        <html>
          <head>
            <style>
              html, body {
                margin: 0;
                padding: 0;
                width: 100%;
                height: 100%;
              }
            </style>
          </head>
          <body>
            ${props.content}
          </body>
        </html>
      `);
      doc.close();
    }
  }
}

async function resizeIframe() {
  await nextTick();

  await new Promise(resolve => setTimeout(resolve, 1000));
  const iframe = iframeContent.value;
  if (iframe && iframe.contentWindow) {
    iframe.style.width = iframe.contentWindow.document.body.scrollWidth + 'px';
    iframe.style.height = iframe.contentWindow.document.body.scrollHeight + 'px';
  }
}

watch(
  () => props.content,
  async () => {
    renderIframeContent();
    resizeIframe();
  },
);
</script>

<style scoped></style>
