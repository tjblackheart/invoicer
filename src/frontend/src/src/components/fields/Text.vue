<template>
  <div class="field">
    <div class="control">
      <slot />
      <help
        :class="{'is-danger': hasError}"
        :text="helptext"
        :link="helplink"
      />
    </div>
  </div>
</template>

<script>
import Help from './Help'

export default {
  components: {
    Help
  },

  props: {
    helptext: {
      type: String,
      default: '',
    },
    helplink: {
      type: Object,
      default: () => {}
    },
    hasError: {
      type: Boolean,
      default: false,
    },
  }
}
</script>

<style>

</style>
