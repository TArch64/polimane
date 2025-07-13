<template>
  <div />
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router';
import { useSessionStore } from '@/stores';
import { authChannel, AuthChannelComplete } from '../channel';

const sessionStore = useSessionStore();

const route = useRoute('authComplete');
const accessToken = route.query['access-token'] as string;
const refreshToken = route.query['refresh-token'] as string;

if (accessToken && refreshToken) {
  sessionStore.setTokens(accessToken, refreshToken);
  authChannel.postMessage(AuthChannelComplete);
}

window.close();
</script>
