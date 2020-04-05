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
      <div class="select is-fullwidth">
        <select
          :id="id"
          :value="value"
          @change="$emit('input', $event.target.value)"
        >
          <option
            v-for="(o, i) in options"
            :key="i"
            :value="o.value"
          >
            {{ o.text }}
          </option>
        </select>

        <help
          :class="{'is-danger': hasError}"
          :text="helptext"
          :link="helplink"
        />
      </div>
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
    value: {
      type: [Number, String],
      default: ''
    },
    label: {
      type: String,
      default: ''
    },
    id: {
      type: String,
      required: true,
    },
    options: {
      type: Array,
      required: true
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
      default: false
    }
  }
}
</script>
