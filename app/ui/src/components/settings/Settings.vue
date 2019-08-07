<template>
  <div>
    <div class="is-clearfix">
      <div class="is-pulled-left">
        <h1
          id="top"
          class="title is-4">
          Settings
        </h1>
        <h2 class="subtitle is-6">
          Application config
        </h2>
      </div>

      <div class="is-pulled-right">
        <button
          :class="['button is-primary', { 'is-loading': busy }]"
          :disabled="dirty === false"
          @click="submit()">
          Save changes
        </button>
      </div>
    </div>

    <hr>

    <message />

    <div class="columns">
      <div class="column is-3">
        <settings-menu
          :items="items"
          :errors="errors"
          @select="toggle($event)" />
      </div>
      <div class="column is-9">
        <keep-alive>
          <component
            :is="activeView"
            v-model="user"
            @error="handleError($event)" />
        </keep-alive>
      </div>
    </div>
  </div>
</template>

<script>
import { mapMutations } from 'vuex'
import http from '@/modules/http'
import Message from '@/components/misc/Message'

import SettingsMenu from './Menu'
import User from './User'
import Banking from './Banking'
import Numbers from './Numbers'
import Company from './Company'
import Contacts from './Contacts'

export default {
  components: {
    Message,
    SettingsMenu,
    User,
    Banking,
    Numbers,
    Company,
    Contacts,
  },

  data () {
    return {
      user: {},
      defaults: {}, // clone of 'user'
      items: [
        { view: 'user', title: 'Profile', active: false },
        { view: 'banking', title: 'Banking', active: false },
        { view: 'numbers', title: 'Numbers', active: false },
        { view: 'company', title: 'Company', active: false },
        { view: 'contacts', title: 'Contacts', active: false },
      ],
      dirty: false,
      busy: false,
      errors: [],
    }
  },

  computed: {
    activeView () {
      const item = this.items.find(i => i.active === true)
      return (item) ? item.view : this.items[0].view
    },
  },

  watch: {
    user: {
      handler (user) {
        if (JSON.stringify(user) !== JSON.stringify(this.defaults)) {
          this.dirty = true
        }
      },
      deep: true,
    },

    activeView () {
      this.clearMessage()
    },
  },

  created () {
    this.clearMessage()
    this.load()
    this.toggle(this.activeView)
  },

  methods: {
    ...mapMutations([ 'setMessage', 'clearMessage', 'setUser' ]),

    async load () {
      this.clearMessage()

      try {
        this.user = await http.fetchUser(this.$store.getters.uuid)
        this.setDefaults()
      } catch (error) {
        this.setMessage({
          text: error.message,
          style: 'is-danger',
        })
      }
    },

    async submit () {
      try {
        this.clearMessage()
        if (this.checkErrors()) {
          this.setMessage({
            text: 'There are errors in your data.',
            style: 'is-danger',
          })
          return
        }

        this.busy = true

        const user = await http.putUser(this.user)
        this.setUser(user)
        this.setMessage({ text: 'Settings saved.' })

        this.dirty = false
        this.setDefaults()
      } catch (error) {
        this.setMessage({
          text: error.message,
          style: 'is-danger',
        })
      } finally {
        this.busy = false
      }
    },

    toggle (view) {
      this.items.map(item => {
        item.active = item.view === view || false
      })
    },

    setDefaults () {
      this.defaults = JSON.parse(JSON.stringify(this.user))
    },

    handleError (event) {
      const { key, errors } = { ...event }
      const err = this.errors.find(e => e.key === key)

      if (!err) {
        this.errors.push({ key, errors })
      } else {
        err.errors = errors
      }
    },

    checkErrors () {
      return this.errors.some(e => e.errors)
    },
  },

  beforeRouteLeave (to, from, next) {
    if (this.dirty) {
      if (!confirm('There are unsaved changes. Really leave?')) {
        return
      }
    }

    this.clearMessage()
    next()
  },
}
</script>

<style lang="scss" scoped>
  /deep/ .fieldset {
    margin-bottom: 60px;
    padding: 20px;
    border-radius: 2px;
    background: #fdfdfd;
    border: 1px solid #e9e9e9;
    border-radius: 2px;
  }

  /deep/ .is-bordered {
    border-bottom: 1px solid #f1f1f1;
    margin-bottom: 20px;
  }
</style>
