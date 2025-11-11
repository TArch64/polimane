<template>
  <SelectField variant="control" v-model="model">
    <option
      v-for="option of accessOptions"
      :key="option.level"
      :value="option.level"
    >
      {{ option.title }}
    </option>
  </SelectField>
</template>

<script setup lang="ts">
import { SelectField } from '@/components/form';
import { AccessLevel, AccessLeveList, getAccessLevelTitle } from '@/enums';

const model = defineModel<AccessLevel>({ required: true });

const accessOptions = AccessLeveList.map((level) => ({
  level,
  title: getAccessLevelTitle(level),
}));
</script>
