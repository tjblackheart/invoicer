import Vue from 'vue'
import VueScrollTo from 'vue-scrollto'
import Vuelidate from 'vuelidate'

import App from './App.vue'
import router from './modules/router'
import store from './modules/store'
import './modules/filters'
import './assets/app.scss'
import './assets/fonts.js'

Vue.config.productionTip = process.env.NODE_ENV !== 'production'
Vue.use(VueScrollTo, { easing: 'ease', offset: -70 })
Vue.use(Vuelidate)

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
