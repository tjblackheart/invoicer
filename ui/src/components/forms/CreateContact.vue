<template>
  <div>
    <message />

    <div class="field">
      <label class="label">
        Type
      </label>
      <div class="control">
        <div class="select is-fullwidth">
          <select v-model="contact.type">
            <option
              v-for="(t, i) in types"
              :key="i">
              {{ t }}
            </option>
          </select>
        </div>
      </div>
    </div>

    <div class="field">
      <label class="label">
        Value
      </label>
      <div class="control">
        <input
          v-model="contact.value"
          class="input"
          type="text">
      </div>
    </div>
  </div>
</template>

<script>
import { mapMutations } from 'vuex'
import Message from '@/components/misc/Message'

export default {
  components: {
    Message,
  },

  props: {
    contact: {
      type: Object,
      default: () => {
        return {
          type: 'Email',
          value: '',
        }
      },
    },
  },

  data () {
    return {
      types: [ 'Email', 'Phone', 'Other' ],
      error: null,
    }
  },

  methods: {
    ...mapMutations([ 'setMessage', 'clearMessage' ]),
  },
}
</script>
