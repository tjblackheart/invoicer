<template>
  <p
    v-if="text || link"
    class="help"
  >
    <span v-if="text"> {{ text }} </span>
    <router-link
      v-if="link"
      :to="{ name: link.to }"
    >
      {{ link.text }}
    </router-link>
  </p>
</template>

<script>
export default {
  props: {
    text: {
      type: String,
      default: '',
    },
    link: {
      type: Object,
      default: () => {}
    }
  }
}
</script>

<style>

</style>
