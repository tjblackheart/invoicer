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
          class="button is-primary"
          :disabled="dirty === false"
          @click="submit()">
          Save changes
        </button>
      </div>
    </div>

    <message />
    <hr>

    <div class="columns">
      <div class="column is-3">
        <settings-menu
          :items="items"
          @select="toggle($event)" />
      </div>
      <div class="column is-9">
        <keep-alive>
          <component
            :is="activeView"
            v-model="user" />
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
      user: {
        username: '',
        settings: {},
      },

      items: [
        { view: 'user', title: 'Profile', active: false },
        { view: 'banking', title: 'Banking', active: false },
        { view: 'numbers', title: 'Numbers', active: false },
        { view: 'company', title: 'Company', active: false },
        { view: 'contacts', title: 'Contacts', active: false },
      ],

      dirty: false,
    }
  },

  computed: {
    activeView () {
      return this.items.find(i => i.active === true).view
    },
  },

  created () {
    this.clearMessage()
    this.load()
    this.toggle(this.items[0].view)
  },

  methods: {
    ...mapMutations([ 'setMessage', 'clearMessage', 'setUser' ]),

    async load () {
      this.clearMessage()

      try {
        this.user = await http.fetchUser(this.$store.getters.uuid)
      } catch (error) {
        this.setMessage({
          text: error.message,
          style: 'is-danger',
        })
      }
    },

    async submit () {
      // TODO
      console.log('submit', this.user)
    },

    toggle (view) {
      const index = this.items.findIndex(i => i.view === view)
      this.items.forEach(i => {
        i.active = false
      })
      this.items[index].active = true
    },
  },

  beforeRouteLeave (to, from, next) {
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
