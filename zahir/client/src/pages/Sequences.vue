<script lang="ts">
import Sequence from "../core/sequence";

export default {
  name: "Sequences",
  data() {
    return {
      sequences: [] as Sequence[],
    };
  },
  async mounted() {
    this.sequences = await Sequence.fetchAll();
  },
};
</script>

<template>
  <div class="title">Sequences</div>
  <table class="table">
    <thead>
      <tr>
        <th>Name</th>
        <th>Description</th>
        <th>Length</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="sequence in sequences" :key="sequence.id">
        <td>
          <RouterLink :to="`/sequences/${sequence.id}`">{{ sequence.name }}</RouterLink>
        </td>
        <td>{{ sequence.description }}</td>
        <td>{{ sequence.length }}</td>
      </tr>
    </tbody>
  </table>
</template>
