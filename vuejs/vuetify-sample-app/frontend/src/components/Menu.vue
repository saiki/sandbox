<template>
    <v-navigation-drawer persistent v-model="drawer" clipped :mini-variant.sync="mini">
        <v-list dense>
            <v-list-tile>
                <v-list-tile-action>
                    <v-icon>home</v-icon>
                </v-list-tile-action>
                <v-list-tile-content>
                    <v-list-tile-title>
                        <router-link to="/">Home</router-link>
                    </v-list-tile-title>
                </v-list-tile-content>
            </v-list-tile>
            <v-list-tile>
                <v-list-tile-action>
                    <v-icon>home</v-icon>
                </v-list-tile-action>
                <v-list-tile-content>
                    <v-list-tile-title>
                        <router-link to="/hello">Hello</router-link>
                    </v-list-tile-title>
                </v-list-tile-content>
            </v-list-tile>
        </v-list>
        <v-footer absolute v-show="footer">
            <span>aaa</span>
        </v-footer>
    </v-navigation-drawer>
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
export default class Menu extends Vue {
    mini: Boolean
    get footer() {
        return !this.mini
    }
}
</script>
