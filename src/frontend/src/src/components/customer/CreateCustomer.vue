<template>
  <div>
    <h1 class="title is-4">
      Customers
    </h1>
    <h2 class="subtitle is-6">
      <span v-if="!edit"> Create new customer </span>
      <span v-else> Edit customer </span>
    </h2>

    <message />

    <hr>

    <form @submit.prevent="submit">
      <b-text :helplink="{ to: 'settings', text: 'Edit ...' }">
        Number: <strong> {{ customer.number }} </strong>
      </b-text>

      <hr>

      <div class="columns">
        <div class="column">
          <h3 class="title is-6">
            Address
          </h3>

          <b-input
            id="c.company"
            v-model.trim="customer.address.company"
            type="text"
            placeholder="Company"
          />

          <div class="columns nm">
            <div class="column">
              <b-input
                id="c.firstname"
                v-model.trim="customer.address.first_name"
                placeholder="First name"
              />
            </div>
            <div class="column">
              <b-input
                id="c.lastname"
                v-model.trim="customer.address.last_name"
                placeholder="Last name"
              />
            </div>
          </div>

          <div class="columns nm">
            <div class="column is-9">
              <b-input
                id="c.street"
                v-model.trim="customer.address.street"
                placeholder="Street"
              />
            </div>
            <div class="column">
              <b-input
                id="c.number"
                v-model.trim="customer.address.number"
                placeholder="Number"
              />
            </div>
          </div>

          <b-input
            id="c.zip"
            v-model.trim="customer.address.zip"
            placeholder="Zipcode"
          />

          <b-input
            id="c.city"
            v-model.trim="customer.address.city"
            placeholder="City"
          />

          <b-input
            id="c.country"
            v-model.trim="customer.address.country"
            placeholder="Country"
          />
        </div>

        <div class="column">
          <h3 class="title is-6">
            VAT ID
          </h3>

          <b-input
            id="c.tax_id"
            v-model.trim="customer.tax_number"
            placeholder="VATID"
          />

          <hr>

          <h3 class="title is-6">
            Contacts
          </h3>

          <table class="table is-fullwidth">
            <tbody>
              <tr
                v-for="(c, index) in customer.contacts"
                :key="index"
              >
                <td>{{ c.type }}</td>
                <td>{{ c.value }}</td>
                <td class="has-text-right">
                  <a @click.prevent="editContact(index)"> Edit </a> &middot;
                  <a @click.prevent="removeContact(index)"> Remove </a>
                </td>
              </tr>
            </tbody>
          </table>

          <button
            class="button is-small"
            @click.prevent="showModal = true"
          >
            Add contact
          </button>
        </div>
      </div>

      <hr>

      <div class="has-text-right">
        <button
          :class="{'is-loading': busy}"
          :disabled="busy"
          class="button is-primary"
          @click.prevent="submit"
        >
          Save
        </button> &nbsp;
        <button
          class="button"
          @click.prevent="$router.go(-1)"
        >
          Cancel
        </button>
      </div>

      <modal
        v-if="showModal"
        :title="title"
        :error="addContactError"
        @close="close"
      >
        <div slot="content">
          <contact-form
            :contact="contact"
            :edit="contactEdit"
          />
        </div>

        <button
          slot="action"
          class="button is-primary"
          @click.prevent="addContact"
        >
          <span v-if="!contactEdit"> Add </span>
          <span v-else> Save </span>
        </button>
      </modal>
    </form>
  </div>
</template>

<script>
import { mapMutations } from 'vuex'
import http from '@/modules/http'
import Modal from '@/components/modals/Modal.vue'
import ContactForm from '@/components/modals/Contact.vue'
import Message from '@/components/misc/Message'
import BText from '@/components/fields/Text'
import BInput from '@/components/fields/Input'

const { required } = require('vuelidate/lib/validators')

export default {
  components: {
    Modal,
    ContactForm,
    Message,
    BText,
    BInput,
  },

  data () {
    return {
      user: null,
      customer: {
        number: null,
        address: {
          country: 'Germany',
        },
        contacts: [],
        user_id: null,
      },
      showModal: false,
      contact: {
        type: 'Email',
        value: '',
      },
      edit: false,
      contactEdit: false,
      busy: false,
      addContactError: '',
    }
  },

  computed: {
    title () {
      return !this.contactEdit ? 'Add contact' : 'Edit contact'
    },
  },

  created () {
    this.clearMessage()
    this.load()
  },

  methods: {
    ...mapMutations([ 'setMessage', 'clearMessage' ]),

    async load () {
      try {
        this.user = await http.fetchUser(this.$store.getters.uuid)

        if (this.user.settings.user_id === 0) {
          this.setMessage({
            text: 'Please review your application settings.',
            style: 'is-warning',
          })
          this.$router.push('/settings')
        }

        this.setCustomerNumber()
        this.$store.commit('setUser', this.user)
      } catch (error) {
        this.setMessage({
          text: error.message,
          style: 'is-danger',
        })
      }
    },

    setCustomerNumber () {
      const prefix = this.user.settings.customer_number_prefix
      const nr = this.user.settings.next_customer_number

      this.customer.number = `${prefix}${nr}`
      this.customer.user_id = this.user.id
    },

    async submit () {
      this.busy = true
      try {
        if (!this.edit) {
          await http.postCustomer(this.customer)
        } else {
          this.customer = await http.putCustomer(this.customer)
        }
        this.$router.go(-1)
      } catch (error) {
        this.setMessage({
          text: error.message,
          style: 'is-danger',
        })
      } finally {
        this.busy = false
      }
    },

    close () {
      this.addContactError = ''
      this.contact = { type: 'Email', value: '' }
      this.showModal = false
      this.contactEdit = false
    },

    addContact () {
      this.$v.$touch()

      if (!this.$v.$error) {
        if (!this.contactEdit) {
          this.customer.contacts.push(this.contact)
        } else {
          this.contactEdit = false
        }

        this.contact = { type: 'Email', value: '' }
        this.showModal = false
      } else {
        this.addContactError = 'Please fill in all the data.'
      }
    },

    editContact (index) {
      this.addContactError = ''
      this.contactEdit = true
      this.contact = this.customer.contacts[index]
      this.showModal = true
    },

    removeContact (index) {
      this.customer.contacts.splice(index, 1)
    },
  },

  beforeRouteEnter (to, from, next) {
    next(async vm => {
      if (to.name === 'customer_details') {
        try {
          vm.edit = true
          vm.customer = await http.fetchCustomer(to.params.id)
        } catch (error) {
          vm.setMessage({
            text: error.message,
            style: 'is-danger',
          })
        }
      }
    })
  },

  validations: {
    contact: {
      value: { required },
    },
  },
}
</script>

<style lang="scss" scoped>
  .nm { margin-bottom: 0 !important; }
</style>
