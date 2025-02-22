<script lang="ts">
import Source from "../core/source";

export default {
  name: "sources",
  data() {
    return {
      sources: [] as Source[],
    };
  },
  async mounted() {
    await this.loadSources();
  },
  methods: {
    async loadSources() {
      this.sources = await Source.fetchAll();
    },
    async createSource() {
      await Source.createNew();
      await this.loadSources();
    },
  },
};
</script>

<template>
  <div class="title">
    Sources
    <button class="btn btn-outline-primary float-end" @click="createSource()">
      <font-awesome-icon :icon="['fas', 'plus']" />New Source
    </button>
  </div>
  <table class="table">
    <thead>
      <tr>
        <th>Name</th>
        <th>Description</th>
        <th>Type</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="source in sources" :key="source.id">
        <td>
          <RouterLink :to="`/sources/${source.id}`">{{ source.name }}</RouterLink>
        </td>
        <td>{{ source.description }}</td>
        <td>{{ source.type }}</td>
      </tr>
    </tbody>
  </table>
</template>
