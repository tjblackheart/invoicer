<template>
  <span>
    <a
      v-if="!busy || button"
      href
      :class="{ 'button is-primary': button, 'is-loading': busy }"
      :disabled="busy"
      @click.prevent="printInvoice"
    >
      {{ text }}
    </a>
    <icon
      v-if="busy && !button"
      name="circle-notch"
      spin
    />
  </span>
</template>

<script>
import http from '@/modules/http'

import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/circle-notch'

export default {
  components: {
    Icon,
  },

  props: {
    id: {
      type: Number,
      required: true,
    },
    text: {
      type: String,
      default: 'PDF',
    },
    button: {
      type: Boolean,
      default: false,
    },
  },

  data () {
    return {
      busy: false,
    }
  },

  methods: {
    async printInvoice () {
      this.$store.commit('clearMessage')
      this.busy = true

      try {
        const r = await http.printInvoice(this.id)
        let a = document.createElement('a')

        a.href = `data:application/octet-stream;base64,${r.content}`
        a.download = r.filename
        a.click()
      } catch (err) {
        this.$store.commit('setMessage', {
          text: err,
          style: 'is-danger'
        })
      } finally {
        this.busy = false
      }
    },
  },
}
</script>
