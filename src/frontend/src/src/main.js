import './assets/app.scss'

import Vue from 'vue'
import Vuelidate from 'vuelidate'
import Showdown from 'showdown'
import App from './App.vue'
import router from './modules/router'
import store from './modules/store'
import './modules/filters'

Vue.config.productionTip = process.env.NODE_ENV !== 'production'
Vue.use(Vuelidate)

Vue.prototype.$md = new Showdown.Converter({
  simpleLineBreaks: true,
  openLinksInNewWindow: true,
})

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
