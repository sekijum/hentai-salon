<template>
  <video ref="video" class="video-js"></video>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, watch } from 'vue';
import videojs from 'video.js';
import 'video.js/dist/video-js.css';

const props = defineProps({
  src: {
    type: String,
    required: true,
  },
});

const video = ref(null);
let player = null;

onMounted(() => {
  player = videojs(
    video.value,
    {
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
    },
    function onPlayerReady() {
      console.log('Player is ready!');
    },
  );
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
