<template>
  <span>
    <a
      v-if="!busy || button"
      href
      :class="{ 'button is-primary': button, 'is-loading': busy }"
      :disabled="busy"
      @click.prevent="printInvoice">
      {{ text }}
    </a>
    <font-awesome-icon
      v-if="busy && !button"
      icon="circle-notch"
      spin />
  </span>
</template>

<script>
import { mapMutations } from 'vuex'
import http from '@/modules/http'

export default {
  props: {
    id: {
      type: String,
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
    ...mapMutations([ 'setMessage', 'clearMessage' ]),

    async printInvoice () {
      this.clearMessage()
      this.busy = true

      try {
        const r = await http.printInvoice(this.id)

        let a = document.createElement('a')
        a.href = `data:application/octet-stream;base64,${r.content}`
        a.download = r.filename
        a.click()
      } catch (err) {
        this.setMessage({ text: err, style: 'is-danger' })
      } finally {
        this.busy = false
      }
    },
  },
}
</script>
