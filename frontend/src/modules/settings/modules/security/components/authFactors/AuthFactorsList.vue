<template>
  <CardList>
    <AuthFactorsListItem
      v-for="factor of authFactorsStore.list"
      :key="factor.id"
      :factor
    />

    <Button
      class="auth-factors-list__add"
      :loading="createAuthFactor.isActive"
      @click="createAuthFactor"
    >
      Додати ще
    </Button>
  </CardList>
</template>

<script setup lang="ts">
import { CardList } from '@/components/card';
import { Button } from '@/components/button';
import { useAuthFactorsStore } from '../../stores';
import { useCreateAuthFactor } from '../../composables';
import AuthFactorsListItem from './AuthFactorsListItem.vue';

const authFactorsStore = useAuthFactorsStore();
const createAuthFactor = useCreateAuthFactor();
</script>

<style>
@layer page {
  .auth-factors-list__add {
    margin: 0 auto;
  }
}
</style>
