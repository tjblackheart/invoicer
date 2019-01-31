import Vue from 'vue'
import VueScrollTo from 'vue-scrollto'
import App from './App.vue'
import router from './modules/router'
import store from './modules/store'
import './modules/filters'
import './assets/app.scss'

Vue.config.productionTip = process.env.NODE_ENV !== 'production'
Vue.use(VueScrollTo, { easing: 'ease', offset: -70 })

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
