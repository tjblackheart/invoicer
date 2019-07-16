import axios from 'axios'
import store from '@/modules/store'
import router from '@/modules/router'

const baseURL = process.env.VUE_APP_ENDPOINT || 'http://localhost:3000'

const instance = axios.create({
  baseURL: baseURL,
  timeout: 3000,
  headers: { 'Content-Type': 'application/json' },
})

instance.interceptors.request.use(
  config => {
    const token = store.getters.token
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  error => Promise.reject(error)
)

instance.interceptors.response.use(null, error => {
  if (error.response.status === 401) {
    store.commit('logout')
    store.commit('setMessage', { message: error.response.data, css: 'is-warning' })
    store.commit('setLayout', 'login-view')
    router.push('/login')
  }
  return Promise.reject(error)
})

export default {
  async login (payload) {
    try {
      const r = await instance.post('/auth/login', payload)
      this.addHeader('Authorization', `Bearer ${r.data.token}`)

      return r.data
    } catch (error) {
      this.handleError(error)
    }
  },

  async fetchInvoices () {
    try {
      const r = await instance.get('/api/invoice')
      return r.data
    } catch (error) {
      this.handleError(error)
    }
  },

  async fetchInvoice (id) {
    try {
      const r = await instance.get(`/api/invoice/${id}`)
      return r.data
    } catch (error) {
      this.handleError(error)
    }
  },

  async postInvoice (invoice) {
    try {
      const r = await instance.post('/api/invoice', invoice)
      return r.data
    } catch (e) {
      this.handleError(e)
    }
  },

  async printInvoice (id) {
    try {
      const r = await instance.get(`/api/invoice/pdf/${id}`)
      return baseURL + `/file/${r.data.file}`
    } catch (e) {
      this.handleError(e)
    }
  },

  async togglePayment (payload) {
    try {
      const r = await instance.post('/api/invoice/payment', payload)
      return r.data
    } catch (e) {
      this.handleError(e)
    }
  },

  async fetchCustomers () {
    try {
      const r = await instance.get('/api/customer')
      return r.data
    } catch (error) {
      this.handleError(error)
    }
  },

  async fetchCustomer (id) {
    try {
      const r = await instance.get(`/api/customer/${id}`)
      return r.data
    } catch (error) {
      this.handleError(error)
    }
  },

  async putCustomer (customer) {
    try {
      const r = await instance.put(`/api/customer/${customer.id}`, customer)
      return r.data
    } catch (error) {
      this.handleError(error)
    }
  },

  async removeCustomer (id) {
    try {
      const r = await instance.delete(`/api/customer/${id}`)
      return r.data
    } catch (error) {
      this.handleError(error)
    }
  },

  async postCustomer (customer) {
    try {
      const r = await instance.post('/api/customer', customer)
      return r.data
    } catch (e) {
      this.handleError(e)
    }
  },

  async fetchUser (uuid) {
    try {
      const r = await instance.get(`/api/user/${uuid}`)
      return r.data
    } catch (e) {
      this.handleError(e)
    }
  },

  async createUser (user) {
    try {
      const r = await instance.post('/auth/register', user)
      return r.data
    } catch (e) {
      this.handleError(e)
    }
  },

  async putUser (user) {
    try {
      const r = await instance.put(`/api/user/${user.uuid}`, user)
      return r.data
    } catch (e) {
      this.handleError(e)
    }
  },

  handleError (error) {
    let msg = 'unspecified'

    if (error.response) {
      msg = `${error.response.data}`
    } else if (error.message) {
      msg = `${error.message}`
    } else {
      msg = `${error.text}`
    }

    throw new Error(msg)
  },

  addHeader (key, value) {
    instance.defaults.headers.common[key] = value
  },
}
