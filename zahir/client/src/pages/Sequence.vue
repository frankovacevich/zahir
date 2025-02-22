<script lang="ts">
import playerState from "../core/player";
import GraphCollection from "../core/graphs";
import Sequence from "../core/sequence";

export default {
  name: "Sequence",
  computed: {
    id(): number {
      return +this.$route.params.id;
    },
  },
  data() {
    return {
      sequence: undefined as Sequence | undefined,
      playerState,
    };
  },
  methods: {
    drawGraphs() {
      const variables = this.sequence?.getVariables() || [];
      const graphCollection = new GraphCollection(this.sequence as Sequence);
      variables.forEach((variable) => {
        graphCollection.addGraph(variable);
      });

      playerState.subscribe(this, () => {
        if (
          playerState.connected &&
          playerState.running &&
          playerState.currentSequenceID === this.id
        ) {
          graphCollection.placePlayLine(playerState.step);
        } else {
          graphCollection.placePlayLine(null);
        }
      });
    },
  },
  async mounted() {
    this.sequence = await Sequence.fetchById(this.id);
    this.$nextTick(this.drawGraphs); // Wait for the DOM to be updated
  },
  beforeUnmount() {
    playerState.unsubscribe(this);
  },
};
</script>

<template>
  <!----------------- Top bar ----------------->

  <div class="title">
    <RouterLink to="/sequences">Sequences</RouterLink>
    <font-awesome-icon :icon="['fas', 'chevron-right']" />
    <a>{{ sequence?.name }}</a>

    <template v-if="!playerState.running">
      <button type="button" class="btn btn-outline-success float-end" @click="sequence?.run()">
        <font-awesome-icon :icon="['fas', 'play']" />
        Start
      </button>
    </template>
    <template v-if="playerState.running">
      <button type="button" class="btn btn-outline-danger float-end" @click="playerState.stop()">
        <font-awesome-icon :icon="['fas', 'stop']" />
        Stop
      </button>
    </template>
  </div>

  <!----------------- Content ----------------->
  <template v-for="source in sequence?.getSources()">
    <div class="row" v-for="variable in source.getVariables()" :key="variable.id">
      <div class="col-3 align-self-center">
        <a :id="`title-${variable.id}`">{{ variable.name }}</a
        ><br />
        <small>{{ source.name }}</small>
      </div>
      <div class="col-9 p-0">
        <svg :id="`svg-${variable.id}`"></svg>
      </div>
    </div>
  </template>

  <!----------------- Popover ----------------->
  <div id="popover">
    <input id="popover-input" type="text" autocomplete="off" class="form-control" />
    <a id="popover-button-one" class="btn btn-primary mt-2 mr-2">Set one</a>
    <a id="popover-button-all" class="btn btn-primary mt-2 ms-2">Set all</a>
    <div id="popover-arrow"></div>
  </div>
</template>

<style scoped>
#popover {
  position: absolute;
  display: none;
  background-color: var(--bs-light);
  box-shadow: 0 0.1rem 0.4rem rgba(0, 0, 0, 0.5);
  border-radius: 0.375rem;
  padding: 10px;
  z-index: 1000;
  width: 181px;
}
</style>
