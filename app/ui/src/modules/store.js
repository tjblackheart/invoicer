import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    uuid: null,
    user: null,
    token: null,
    layout: 'login-view',
    loggedIn: false,
    message: {
      text: '',
      style: 'is-success',
    },
  },

  getters: {
    loggedIn: state => state.loggedIn ? true : !!window.sessionStorage.getItem('token'),
    layout: state => state.layout,
    token: state => state.token || window.sessionStorage.getItem('token'),
    user: state => state.user,
    username: state => state.user ? state.user.username : window.sessionStorage.getItem('username'),
    uuid: state => state.uuid || window.sessionStorage.getItem('uuid'),
    message: state => state.message,
  },

  mutations: {
    setLayout (state, layout) {
      state.layout = layout
    },

    setUser (state, user) {
      state.user = user
      window.sessionStorage.setItem('username', user.username)
    },

    login (state, payload) {
      state.token = payload.token
      state.user = payload.user
      state.username = payload.user.username

      window.sessionStorage.setItem('token', state.token)
      window.sessionStorage.setItem('uuid', state.user.uuid)
      window.sessionStorage.setItem('username', state.user.username)

      state.loggedIn = true
    },

    logout (state) {
      window.sessionStorage.removeItem('token')
      window.sessionStorage.removeItem('uuid')
      window.sessionStorage.removeItem('username')

      state.token = null
      state.user = {}
      state.uuid = null

      state.loggedIn = false
    },

    setMessage (state, { text, style }) {
      state.message = {
        text,
        style: style || 'is-success',
      }
    },

    clearMessage (state) {
      state.message = {
        text: '',
        style: 'is-success',
      }
    },
  },

  actions: {

  },

  strict: true,
})
