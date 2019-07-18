import Vue from 'vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faCircleNotch } from '@fortawesome/free-solid-svg-icons'

library.add(faCircleNotch)

Vue.component('font-awesome-icon', FontAwesomeIcon)
