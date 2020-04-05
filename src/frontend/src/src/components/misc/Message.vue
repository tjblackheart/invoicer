<template>
  <div
    v-if="message.text"
    :class="message.style"
    class="notification"
  >
    {{ message.text }}
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  data () {
    return {}
  },

  computed: {
    ...mapGetters([ 'message' ]),
  },
}
</script>
