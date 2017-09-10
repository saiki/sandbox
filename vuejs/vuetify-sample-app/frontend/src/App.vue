<template>
  <v-app id="app" toolbar footer>
    <side-menu :drawer="drawer" :mini="mini"></side-menu>
    <app-header :drawer="drawer" :mini="mini" @mini="onResize" @drawer="onVisible"></app-header>
    <main>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </main>
    <v-footer>
      Copyright Aaa
    </v-footer>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'

import Menu from '@/components/Menu.vue'
import AppHeader from '@/components/AppHeader.vue'

@Component(
  {
    components: {
      'side-menu': Menu,
      'app-header': AppHeader
    }
  }
)
export default class App extends Vue {
  drawer = true;
  mini = false;
  onVisible() {
    if (!this.drawer) {
      this.mini = false
    }
    this.drawer = !this.drawer
  }
  onResize() {
    this.mini = !this.mini
  }
}
</script>
