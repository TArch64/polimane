<template>
  <div />
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router';
import { authChannel, AuthChannelComplete } from '../channel';

const route = useRoute('auth-complete');
const accessToken = route.query['access-token'] as string;
const refreshToken = route.query['refresh-token'] as string;

if (accessToken && refreshToken) {
  authChannel.postMessage({
    type: AuthChannelComplete,
    accessToken,
    refreshToken,
  });
}

window.close();
</script>
