<template>
  <BackgroundPattern :scale="1">
    <BackgroundCell x="0" y="0" />
  </BackgroundPattern>
</template>

<script setup lang="ts">
import BackgroundPattern from './BackgroundPattern.vue';
import BackgroundCell from './BackgroundCell.vue';
</script>
