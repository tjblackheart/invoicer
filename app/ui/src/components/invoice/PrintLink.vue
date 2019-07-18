<template>
  <span>
      &middot;
      <a
        v-if="!busy"
        href
        :class="{ 'button is-primary': button, 'is-loading': button && busy }"
        :disabled="busy"
        @click.prevent="printInvoice">
        {{ text }}
      </a>
      <font-awesome-icon
        v-if="busy"
        icon="circle-notch" spin
      />
  </span>
</template>

<script>
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
    async printInvoice () {
      this.busy = true
      const r = await http.printInvoice(this.id)
      this.busy = false

      let a = document.createElement('a')
      a.href = `data:application/octet-stream;base64,${r.content}`
      a.download = r.filename
      a.click()
    },
  }
}
</script>
