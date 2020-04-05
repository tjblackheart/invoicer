<template>
  <!-- eslint-disable-next-line vue/no-v-html -->
  <div v-html="markdownText" />
</template>

<script>
export default {
  props: {
    md: {
      type: String,
      default: '',
    }
  },

  computed: {
    markdownText () {
      return this.$md.makeHtml(this.md)
    }
  },
}
</script>
