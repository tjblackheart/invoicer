<template>
  <div class="field">
    <div class="control">
      <span v-html="text" />
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
  props: {
    text: {
      type: String,
      default: '',
    },
    helptext: {
      type: String,
      default: '',
    },
    helplink: {
      type: Object,
      default: () => {}
    }
  },

  components: {
    Help
  }
}
</script>

<style>

</style>
