<script lang="ts">
import playerState from "../../core/player";

export default {
  name: "NavBar",
  data() {
    return {
      playerState,
    };
  },
};
</script>

<template>
  <nav class="navbar navbar-expand-xl bg-body-tertiary">
    <div class="container">
      <a class="navbar-brand" href="#">
        <img src="/logo.svg" width="30" />
      </a>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <div class="navbar-nav me-auto mb-2 mb-lg-0">
          <RouterLink active-class="active" class="nav-link" to="/sequences">
            Sequences
          </RouterLink>
          <RouterLink active-class="active" class="nav-link" to="/sources">
            Sources
          </RouterLink>
        </div>
        <div class="btn-group float-end">
          <button
            type="button"
            class="btn dropdown-toggle"
            data-bs-toggle="dropdown"
            aria-expanded="false"
          >
            <font-awesome-icon
              :icon="['fas', 'circle-notch']"
              :style="{ color: playerState.connected ? 'green' : 'red' }"
              :spin="playerState.connected && playerState.running"
            />
            Server
          </button>
          <ul class="dropdown-menu dropdown-menu-end">
            <div class="px-2">Status: {{ playerState.status }}</div>
          </ul>
        </div>
      </div>
    </div>
  </nav>
</template>
