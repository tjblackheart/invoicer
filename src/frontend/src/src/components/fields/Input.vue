<template>
  <div class="field">
    <label
      v-if="label"
      class="label"
      :for="id"
    >
      {{ label }}
    </label>

    <div class="control">
      <input
        :class="['input', {'is-danger': hasError}]"
        :id="id"
        :value="value"
        :type="type"
        :placeholder="placeholder"
        :disabled="disabled"
        @input="$emit('input', $event.target.value)"
        @keydown.escape="$emit('escape')"
      >
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
      type: [Number, String],
      default: '',
    },
    label: {
      type: String,
      default: '',
    },
    id: {
      type: String,
      required: true,
    },
    type: {
      type: String,
      default: 'text',
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
    },
    hasError: {
      type: Boolean,
      default: false,
    },
    autofocus: {
      type: Boolean,
      default: false,
    }
  },

  components: {
    Help
  },

  mounted () {
    if (this.autofocus) {
      document.getElementById(this.id).focus()
    }
  }
}
</script>
