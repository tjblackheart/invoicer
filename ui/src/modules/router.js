import Vue from 'vue'
import Router from 'vue-router'
import store from '@/modules/store'

import InvoiceList from '@/components/invoice/InvoiceList.vue'
import InvoiceCreate from '@/components/invoice/InvoiceCreate.vue'
import InvoiceDetails from '@/components/invoice/InvoiceDetails.vue'
import CustomerList from '@/components/customer/ListCustomers.vue'
import CustomerCreate from '@/components/customer/CreateCustomer.vue'
import Settings from '@/components/user/Settings.vue'
import Login from '@/components/user/Login.vue'
import Register from '@/components/user/Register.vue'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      name: 'index',
      path: '/',
      redirect: '/invoices',
    },
    {
      name: 'invoice_list',
      path: '/invoices',
      component: InvoiceList,
      meta: {
        requiresAuth: true,
      },
    },
    {
      name: 'invoice_create',
      path: '/invoices/create',
      component: InvoiceCreate,
      meta: {
        requiresAuth: true,
      },
    },
    {
      name: 'invoice_details',
      path: '/invoices/view/:id',
      component: InvoiceDetails,
      meta: {
        requiresAuth: true,
      },
    },
    {
      name: 'customer_list',
      path: '/customers',
      component: CustomerList,
      meta: {
        requiresAuth: true,
      },
    },
    {
      name: 'customer_create',
      path: '/customers/create',
      component: CustomerCreate,
      meta: {
        requiresAuth: true,
      },
    },
    {
      name: 'customer_details',
      path: '/customers/edit/:id',
      component: CustomerCreate,
      meta: {
        requiresAuth: true,
      },
    },
    {
      name: 'settings',
      path: '/settings',
      component: Settings,
      meta: {
        requiresAuth: true,
      },
    },
    {
      name: 'login',
      path: '/login',
      component: Login,
    },
    {
      name: 'register',
      path: '/register',
      component: Register,
    },
  ],

  scrollBehavior () {
    return { x: 0, y: 0 }
  },

  linkActiveClass: 'is-active',
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (store.getters.loggedIn === false) {
      store.commit('setLayout', 'login-view')
      next({ path: '/login' })
    } else {
      if (store.getters.layout !== 'main-view') {
        store.commit('setLayout', 'main-view')
      }
      next()
    }
  } else {
    next()
  }
})

export default router
