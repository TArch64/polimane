<template>
  <div />
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router';
import { authChannel, AuthChannelEvent } from '../channel';

const route = useRoute();
const isDeleted = 'deleted' in route.query;
const event = isDeleted ? AuthChannelEvent.DELETED_USER : AuthChannelEvent.COMPLETE;

authChannel.postMessage({ type: event });
window.close();
</script>
