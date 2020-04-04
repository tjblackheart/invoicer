<template>
  <div class="field">
    <label
      v-if="label"
      class="label"
      :for="id"
    > {{ label }}
    </label>

    <div class="control">
      <textarea
        :id="id"
        class="textarea"
        :value="value"
        :placeholder="placeholder"
        :disabled="disabled"
        @input="$emit('input', $event.target.value)"
      />

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
    value: {
      type: String,
      default: '',
    },
    label: {
      type: String,
      default: ''
    },
    id: {
      type: String,
      required: true,
    },
    placeholder: {
      type: String,
      default: '',
    },
    disabled: {
      type: Boolean,
      default: false,
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
