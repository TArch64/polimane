<template>
  <Card title="Двофакторна автентифікація">
    <AuthFactorsList v-if="authFactorsStore.list.length" />
    <AuthFactorsEmpty v-else />
  </Card>
</template>

<script setup lang="ts">
import { Card } from '@/components/card';
import { useAuthFactorsStore } from '../../stores';
import AuthFactorsEmpty from './AuthFactorsEmpty.vue';
import AuthFactorsList from './AuthFactorsList.vue';

const authFactorsStore = useAuthFactorsStore();
</script>
