<template>
  <div>
    <h3 class="title is-5">
      <icon name="user-cog" /> User
    </h3>

    <hr>

    <div class="fieldset">
      <div class="field">
        <label class="label">
          Username
        </label>
        <div class="control">
          <input
            v-model="value.username"
            type="text"
            class="input"
            required>
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
            class="input"
            required>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/user-cog'
import 'vue-awesome/icons/envelope'

export default {
  props: {
    value: {
      type: Object,
      required: true,
    },
  },

  components: {
    Icon,
  },

  watch: {
    value() {
      this.$emit('input', this.value);
    },
  },
}
</script>

