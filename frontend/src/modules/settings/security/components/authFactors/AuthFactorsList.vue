<template>
  <CardList>
    <AuthFactorsListItem
      v-for="factor of authFactorsStore.list"
      :key="factor.id"
      :factor
    />
  </CardList>
</template>

<script setup lang="ts">
import { CardList } from '@/components/card';
import { useAuthFactorsStore } from '../../stores';
import AuthFactorsListItem from './AuthFactorsListItem.vue';

const authFactorsStore = useAuthFactorsStore();
</script>
