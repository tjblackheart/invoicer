<template>
  <label class="checkbox">
    <input
      type="checkbox"
      :checked="value"
      @change="$emit('input', $event.target.checked)"
    > {{ label }}
  </label>
</template>

<script>
export default {
  props: {
    value: {
      type: Boolean,
      required: true
    },
    label: {
      type: String,
      required: true
    },
  },
}
</script>
