<template>
  <div />
</template>

<script lang="ts" setup>
import { authChannel, AuthChannelComplete } from '../channel';

authChannel.postMessage({ type: AuthChannelComplete });
window.close();
</script>
