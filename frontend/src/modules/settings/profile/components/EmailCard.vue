<template>
  <EmailEditForm v-if="isEdit" />
  <EmailVerifyForm v-else />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useProfileStore } from '../stores';
import EmailEditForm from './EmailEditForm.vue';
import EmailVerifyForm from './EmailVerifyForm.vue';

const profileStore = useProfileStore();

const isEdit = computed(() => {
  return profileStore.isChangeVerifyingEmail || profileStore.user.isEmailVerified;
});
</script>
