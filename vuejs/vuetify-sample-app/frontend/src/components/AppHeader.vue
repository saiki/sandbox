<template>
    <v-toolbar>
      <v-toolbar-side-icon @click.stop="onClickSideIcon"></v-toolbar-side-icon>
      <v-btn icon @click.stop="onClickResize" v-show="drawer">
        <v-icon v-if="mini">chevron_right</v-icon>
        <v-icon v-else>chevron_left</v-icon>
      </v-btn>
    </v-toolbar>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'

@Component({
    props: {
        "drawer": Boolean,
        "mini": Boolean
    }
})
export default class AppHeader extends Vue {
    onClickSideIcon() {
        this.$emit("drawer")
    }
    onClickResize() {
        this.$emit("mini")
    }
}
</script>
