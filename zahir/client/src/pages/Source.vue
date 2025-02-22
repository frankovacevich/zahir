<script lang="ts">
import Source from "../core/source";
import Variable from "../core/variable";

export default {
  name: "source",
  computed: {
    id(): number {
      return +this.$route.params.id;
    },
  },
  data() {
    return {
      source: undefined as Source | undefined,
      variables: [] as Variable[],
    };
  },
  async mounted() {
    this.source = await Source.fetchById(this.id);
    this.variables = this.source.getVariables();
  },
};
</script>

<template>
  <div v-if="source !== undefined">
    <div class="title">
      <RouterLink to="/sources">Sources</RouterLink>
      <font-awesome-icon :icon="['fas', 'chevron-right']" />
      <a>{{ source?.name }}</a>
      <button type="submit" class="btn btn-outline-primary float-end">Save</button>
    </div>

    <!---------- Name and description ---------->
    <div class="mb-3">
      <label class="form-label" for="name">Name</label>
      <input type="text" class="form-control" name="name" v-model="source.name" />
    </div>
    <div class="mb-3">
      <label class="form-label" for="description">Description</label>
      <input type="text" class="form-control" name="description" v-model="source.description" />
    </div>
    <div class="mb-3">
      <label class="form-label" for="type">Type</label>
      <select class="form-select" name="type">
        <option value="NONE">None</option>
        <option value="MQTT-JSON">MQTT (JSON)</option>
        <option value="MQTT-SPB">MQTT (Sparkplug B)</option>
      </select>
    </div>

    <!----------------- Message ----------------->
    <h2>Message</h2>
    <div class="mb-3">
      <label class="form-label" for="mqtt-topic">Topic</label>
      <input type="text" class="form-control" name="mqtt-topic" />
    </div>
    <div class="mb-3">
      <label class="form-label" for="template">Message template</label>
      <textarea class="form-control" name="template"></textarea>
    </div>

    <!---------------- Variables ---------------->
    <h2>Variables</h2>
    <table class="table">
      <thead>
        <tr>
          <th>Name</th>
          <th>Description</th>
          <th>Default value</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="variable in variables" :key="variable.id">
          <td><input class="ghost" type="text" v-model="variable.name" /></td>
          <td width="50%"><input class="ghost" type="text" v-model="variable.description" /></td>
          <td><input class="ghost" type="text" v-model="variable.defaultValue" /></td>
          <td>
            <button class="btn btn-light">
              <font-awesome-icon :icon="['fas', 'trash']" />Delete
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
tr > td {
  vertical-align: middle;
}
input.ghost {
  border: none;
  border-radius: 3px;
  max-width: auto;
  width: 100%;
  padding: 5px;
}
</style>
