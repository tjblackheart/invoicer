<template>
  <div class="field">
    <input
      :id="id"
      type="checkbox"
      class="switch is-rounded is-outlined is-small is-info"
      :checked="value"
      @change="emit($event)"
    >
    <label :for="id">
      {{ label }}
    </label>
  </div>
</template>

<script>
export default {
  props: {
    value: {
      type: Boolean,
      required: true
    },
    id: {
      type: String,
      required: true,
    },
    label: {
      type: String,
      required: true
    },
  },

  methods: {
    emit (e) {
      e.target.blur()
      this.$emit('input', e.target.checked)
    }
  }
}
</script>

<style lang="scss" scoped>
label {
  font-size: 14px !important;
  padding-top: 0 !important;
}
</style>
