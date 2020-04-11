<template>
  <div>
    <div class="is-clearfix">
      <div class="is-pulled-left">
        <h1
          id="top"
          class="title is-4"
        >
          Settings
        </h1>
        <h2 class="subtitle is-6">
          Application config
        </h2>
      </div>

      <div class="is-pulled-right">
        <div class="field">
          <div class="control">
            <button
              :class="['button is-primary', { 'is-loading': busy }]"
              :disabled="hasErrors || (!hasErrors && dirty === false)"
              @click="submit()"
            >
              Save changes
            </button>
          </div>
          <!-- <p
            v-if="hasErrors"
            class="help has-text-danger"
          >
            There are errors in the form.
          </p> -->
        </div>
      </div>
    </div>

    <hr>

    <message />

    <div class="columns">
      <div class="column is-3">
        <settings-menu
          :items="items"
          :errors="errors"
          :active="view"
          @select="toggle($event)"
        />
      </div>
      <div class="column is-9">
        <keep-alive>
          <component
            :is="view"
            v-if="inItems(view)"
            v-model="user"
            @error="handleError($event)"
          />
        </keep-alive>
      </div>
    </div>
  </div>
</template>

<script>
import http from '@/modules/http'
import Message from '@/components/misc/Message'

import SettingsMenu from './Menu'
import User from './User'
import Password from './Password'
import Banking from './Banking'
import Numbers from './Numbers'
import Company from './Company'
import Contacts from './Contacts'

export default {
  components: {
    Message,
    SettingsMenu,
    User,
    Password,
    Banking,
    Numbers,
    Company,
    Contacts,
  },

  props: {
    view: {
      type: String,
      default: 'user',
    },
  },

  data () {
    return {
      user: {},
      defaults: {}, // clone of 'user'
      items: [
        { view: 'user', title: 'Profile' },
        { view: 'password', title: 'Password' },
        { view: 'banking', title: 'Banking' },
        { view: 'numbers', title: 'Numbers' },
        { view: 'company', title: 'Company' },
        { view: 'contacts', title: 'Contacts' },
      ],
      dirty: false,
      busy: false,
      errors: [],
    }
  },

  computed: {
    hasErrors () {
      return this.errors.some(e => e.errors)
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

    view () {
      this.$store.commit('clearMessage')
    },
  },

  created () {
    this.load()
  },

  methods: {
    inItems (view) {
      return this.items.some(i => i.view === view)
    },

    async load () {
      try {
        this.user = await http.fetchUser(this.$store.getters.uuid)
        this.setDefaults()
      } catch (error) {
        this.$store.commit('setMessage', {
          text: error.message,
          style: 'is-danger',
        })
      }
    },

    async submit () {
      try {
        this.$store.commit('clearMessage')
        if (this.hasErrors) {
          this.$store.commit('setMessage', {
            text: 'There are errors in the submitted data. Please review.',
            style: 'is-danger',
          })
          return
        }

        this.busy = true

        const user = await http.putUser(this.user)
        this.$store.commit('setUser', user)
        this.$store.commit('setMessage', { text: 'Settings saved.' })

        this.dirty = false
        this.setDefaults()
      } catch (error) {
        this.$store.commit('setMessage', {
          text: error.message,
          style: 'is-danger',
        })
      } finally {
        this.busy = false
      }
    },

    setDefaults () {
      this.defaults = JSON.parse(JSON.stringify(this.user))
    },

    handleError (event) {
      const { view, errors, count } = { ...event }
      const err = this.errors.find(e => e.view === view)

      if (!err) {
        this.errors.push({ view, errors, count })
      } else {
        err.errors = errors
        err.count = count
      }
    },
  },

  beforeRouteEnter (to, from, next) {
    next (vm => {
      const target = to.params.view
      if (!vm.items.some(i => i.view === target)) {
        vm.$router.push(`/settings/${vm.items[0].view}`)
      }
    })
  },

  beforeRouteUpdate (to, from, next) {
    if (!this.items.some(i => i.view === to.params.view)) {
      to.params.view = this.items[0].view
    }
    next()
  },

  beforeRouteLeave (to, from, next) {
    if (this.dirty) {
      if (!confirm('There are unsaved changes. Really leave?')) {
        return
      }
    }

    this.$store.commit('clearMessage')
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
  }

  /deep/ .is-bordered {
    border-bottom: 1px solid #f1f1f1;
    margin-bottom: 20px;
  }
</style>
