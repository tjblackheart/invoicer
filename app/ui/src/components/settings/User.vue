<template>
  <div>
    <h3 class="title is-5">
      <icon name="user-cog" /> User
    </h3>

    <div class="fieldset">
      <div class="field">
        <label class="label">
          Username
        </label>
        <div class="control">
          <input
            v-model="value.username"
            type="text"
            class="input">
        </div>
      </div>

      <div class="field">
        <label class="label">
          Email
        </label>
        <div class="control">
          <input
            v-model="value.email"
            type="email"
            class="input">
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/user-cog'

export default {
  components: {
    Icon,
  },

  props: {
    value: {
      type: Object,
      required: true,
    },
  },

  watch: {
    value () {
      this.$emit('input', this.value)
    },
  },
}
</script>
