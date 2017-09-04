<template>
  <v-app id="app" toolbar footer>
      <side-menu :drawer="drawer" :mini="mini"></side-menu>
    <v-toolbar>
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-btn icon @click.native.stop="mini = !mini">
        <v-icon>chevron_left</v-icon>
      </v-btn>
    </v-toolbar>
    <main>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </main>
    <v-footer></v-footer>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'

import Menu from '@/components/Menu.vue'

@Component(
  {
    components: {
      'side-menu': Menu
    }
  }
)
export default class App extends Vue {
  name = 'app';
  drawer = true;
  mini = false;
}
</script>
