<template>
	<div>
		{{ message }}
    <router-link to="/hello">Go to Hello</router-link>
	</div>
</template>
<script lang="ts">
import Vue from 'vue';
import Component from 'vue-class-component'

@Component
export default class Home extends Vue {
	message = 'Hello,World!'
}

</script>
