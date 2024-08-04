<template>
  <video ref="video" class="video-js vjs-default-skin"></video>
</template>

<script setup lang="ts">
import videojs from 'video.js';
import 'video.js/dist/video-js.css';

const props = defineProps({
  src: {
    type: String,
    required: true,
  },
});

const video = ref<string>();
let player = null;

onMounted(() => {
  player = videojs(video.value, {
    autoplay: false,
    controls: true,
    sources: [
      {
        src: props.src,
        type: 'video/mp4',
      },
    ],
    controlBar: {
      remainingTimeDisplay: {
        displayNegative: false,
      },
    },
    aspectRatio: '16:9',
    fluid: true,
    playbackRates: [0.5, 1, 1.5, 2],
    preferFullWindow: true,
  });

  player.on('fullscreenchange', () => {
    if (player.isFullscreen()) {
      player.exitFullscreen();
    }
  });
});

watch(
  () => props.src,
  newSrc => {
    if (player) {
      player.src({
        src: newSrc,
        type: 'video/mp4',
      });
    }
  },
);

onBeforeUnmount(() => {
  if (player) {
    player.dispose();
  }
});
</script>

<style>
.video-js {
  -webkit-user-select: none;
  -webkit-touch-callout: none;
  -webkit-user-drag: none;
  -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
}

.video-js.vjs-default-skin {
  touch-action: manipulation;
}
</style>
