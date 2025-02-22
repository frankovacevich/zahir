import { createRouter, createWebHistory } from "vue-router";
import Sequences from "./pages/Sequences.vue";
import Sequence from "./pages/Sequence.vue";
import Sources from "./pages/Sources.vue";
import Source from "./pages/Source.vue";

const routes = [
  {
    path: "/",
    name: "Sequences",
    component: Sequences,
  },
  {
    path: "/sequences",
    name: "Sequences",
    component: Sequences,
  },
  {
    path: "/sequences/:id",
    name: "Sequence",
    component: Sequence,
  },
  {
    path: "/sources",
    name: "Sources",
    component: Sources,
  },
  {
    path: "/sources/:id",
    name: "Source",
    component: Source,
  },
];

export const router = createRouter({ history: createWebHistory(), routes });
