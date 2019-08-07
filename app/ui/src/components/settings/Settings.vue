<template>
  <div>
    <h1
      id="top"
      class="title is-4">
      Settings
    </h1>
    <h2 class="subtitle is-6">
      Application config
    </h2>

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
            v-model="user"
            :is="activeView"
          />
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
import Commercial from './Commercial'
import Numbers from './Numbers'
import Company from './Company'
import Contacts from './Contacts'

export default {
  components: {
    Message,
    SettingsMenu,
    User,
    Commercial,
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
        { view: 'user', title: 'User', active: false },
        { view: 'commercial', title: 'Commercial', active: false },
        { view: 'numbers', title: 'Numbers', active: false },
        { view: 'company', title: 'Company', active: false },
        { view: 'contacts', title: 'Contacts', active: false },
      ],
    }
  },

  computed: {
    ...mapMutations([ 'setMessage', 'clearMessage', 'setUser' ]),

    activeView () {
      return this.items.find(i => i.active === true).view
    },
  },

  created () {
    this.load()
    this.toggle(this.items[0].view)
  },

  methods: {
    async load () {
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
      this.items.forEach(i => i.active = false)
      this.items[index].active = true
    },
  },

  beforeRouteLeave (to, from, next) {
    // this.clearMessage()
    next()
  },
}
</script>
