<script lang="ts">
import NavBar from "./components/app/NavBar.vue";

export default {
  name: "App",
  components: {
    NavBar,
  },
};
</script>

<template>
  <NavBar></NavBar>
  <div class="container py-5">
    <router-view></router-view>
  </div>
</template>
