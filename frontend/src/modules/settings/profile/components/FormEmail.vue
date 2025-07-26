<template>
  <FormEmailEdit v-if="isEdit" />
  <FormEmailVerify v-else />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useProfileStore } from '../stores';
import FormEmailEdit from './FormEmailEdit.vue';
import FormEmailVerify from './FormEmailVerify.vue';

const profileStore = useProfileStore();

const isEdit = computed(() => {
  return profileStore.isChangeVerifyingEmail || profileStore.user.isEmailVerified;
});
</script>
