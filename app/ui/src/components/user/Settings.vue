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

    <form @submit.prevent="submit">
      <div
        v-scroll-to="'#user-settings'"
        class="is-clearfix is-bordered is-pointer"
        @click="userSettings = !userSettings">
        <h3
          id="user-settings"
          class="title is-5 is-pulled-left">
          <icon name="user-cog" /> &nbsp; Username
        </h3>
        <p class="is-pulled-right">
          <span v-if="userSettings">
            <icon name="caret-up" />
          </span>
          <span v-else>
            <icon name="caret-down" />
          </span>
        </p>
      </div>

      <div
        v-if="userSettings"
        class="fieldset">
        <div class="control">
          <input
            v-model="user.username"
            type="text"
            class="input">
        </div>
      </div>

      <div
        v-scroll-to="'#bank-settings'"
        class="is-clearfix is-bordered is-pointer"
        @click="bankSettings = !bankSettings">
        <h3
          id="bank-settings"
          class="title is-5 is-pulled-left">
          <icon name="university" /> &nbsp; Bank
        </h3>
        <p class="is-pulled-right">
          <span v-if="bankSettings">
            <icon name="caret-up" />
          </span>
          <span v-else>
            <icon name="caret-down" />
          </span>
        </p>
      </div>

      <div
        v-if="bankSettings"
        class="fieldset">
        <label class="label">
          Bank
        </label>
        <div class="control">
          <input
            v-model="user.settings.bank"
            type="text"
            class="input">
        </div>
        <label class="label">
          IBAN
        </label>
        <div class="control">
          <input
            v-model="user.settings.iban"
            type="text"
            class="input">
        </div>
        <label class="label">
          BIC
        </label>
        <div class="control">
          <input
            v-model="user.settings.bic"
            type="text"
            class="input">
        </div>
      </div>

      <div
        v-scroll-to="'#number-settings'"
        class="is-clearfix is-bordered is-pointer"
        @click="numberSettings = !numberSettings">
        <h3
          id="number-settings"
          class="title is-5 is-pulled-left">
          <icon name="cogs" /> &nbsp; Numbers
        </h3>
        <p class="is-pulled-right">
          <span v-if="numberSettings">
            <icon name="caret-up" />
          </span>
          <span v-else>
            <icon name="caret-down" />
          </span>
        </p>
      </div>

      <div
        v-if="numberSettings"
        id="numbers"
        class="fieldset">
        <div class="field">
          <label class="label">
            Invoicenumber - Prefix
          </label>
          <div class="control">
            <div class="control">
              <input
                v-model="user.settings.invoice_number_prefix"
                type="text"
                class="input">
            </div>
          </div>
        </div>

        <div class="field">
          <label class="label">
            Next invoice number
          </label>
          <div class="control">
            <div class="control">
              <input
                v-model.number="user.settings.next_invoice_number"
                type="number"
                class="input">
            </div>
          </div>
        </div>

        <div class="field">
          <label class="label">
            Customernumber - Prefix
          </label>
          <div class="control">
            <div class="control">
              <input
                v-model="user.settings.customer_number_prefix"
                type="text"
                class="input">
            </div>
          </div>
        </div>

        <div class="field">
          <label class="label">
            Next customer number
          </label>
          <div class="control">
            <div class="control">
              <input
                v-model.number="user.settings.next_customer_number"
                type="number"
                class="input">
            </div>
          </div>
        </div>

        <div class="field">
          <label class="label">
            VAT ID
          </label>
          <div class="control">
            <div class="control">
              <input
                v-model.number="user.settings.tax_number"
                type="text"
                class="input">
            </div>
          </div>
        </div>
      </div>

      <div
        v-scroll-to="'#company-settings'"
        class="is-clearfix is-bordered is-pointer"
        @click="companySettings = !companySettings">
        <h3
          id="company-settings"
          class="title is-5 is-pulled-left">
          <icon name="address-card" /> &nbsp; Company
        </h3>
        <p class="is-pulled-right">
          <span v-if="companySettings">
            <icon name="caret-up" />
          </span>
          <span v-else>
            <icon name="caret-down" />
          </span>
        </p>
      </div>

      <div
        v-if="companySettings"
        id="company"
        class="fieldset">
        <div class="field">
          <label class="label">
            Name
          </label>
          <div class="control">
            <div class="control">
              <input
                v-model="user.settings.company"
                type="text"
                class="input">
            </div>
          </div>
        </div>

        <div class="columns">
          <div class="column">
            <div class="field">
              <label class="label">
                First name
              </label>
              <div class="control">
                <div class="control">
                  <input
                    v-model="user.settings.first_name"
                    type="text"
                    class="input">
                </div>
              </div>
            </div>
          </div>
          <div class="column">
            <div class="field">
              <label class="label">
                Last name
              </label>
              <div class="control">
                <div class="control">
                  <input
                    v-model="user.settings.last_name"
                    type="text"
                    class="input">
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="columns">
          <div class="column is-9">
            <div class="field">
              <label class="label">
                Street
              </label>
              <div class="control">
                <div class="control">
                  <input
                    v-model="user.settings.street"
                    type="text"
                    class="input">
                </div>
              </div>
            </div>
          </div>
          <div class="column">
            <div class="field">
              <label class="label">
                Number
              </label>
              <div class="control">
                <div class="control">
                  <input
                    v-model="user.settings.number"
                    type="text"
                    class="input">
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="columns">
          <div class="column is-3">
            <div class="field">
              <label class="label">
                Zip
              </label>
              <div class="control">
                <div class="control">
                  <input
                    v-model="user.settings.zip"
                    type="text"
                    class="input">
                </div>
              </div>
            </div>
          </div>

          <div class="column is-9">
            <div class="field">
              <label class="label">
                City
              </label>
              <div class="control">
                <div class="control">
                  <input
                    v-model="user.settings.city"
                    type="text"
                    class="input">
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="field">
          <label class="label">
            Country
          </label>
          <div class="control">
            <div class="control">
              <input
                v-model="user.settings.country"
                type="text"
                class="input">
            </div>
          </div>
        </div>
      </div>

      <div
        v-scroll-to="'#contact-settings'"
        class="is-clearfix is-bordered is-pointer"
        @click="contactSettings = !contactSettings">
        <h3
          id="contact-settings"
          class="title is-5 is-pulled-left">
          <icon name="envelope" /> &nbsp; Contacts
        </h3>
        <p class="is-pulled-right">
          <span v-if="contactSettings">
            <icon name="caret-up" />
          </span>
          <span v-else>
            <icon name="caret-down" />
          </span>
        </p>
      </div>

      <div
        v-if="contactSettings"
        id="contacts"
        class="fieldset">
        <div class="field">
          <label class="label">
            Company email
          </label>
          <div class="control">
            <div class="control">
              <input
                v-model="user.settings.email"
                type="text"
                class="input">
            </div>
          </div>
        </div>

        <div class="field">
          <label class="label">
            Phone
          </label>
          <div class="control">
            <div class="control">
              <input
                v-model="user.settings.phone"
                type="text"
                class="input">
            </div>
          </div>
        </div>
      </div>

      <hr>

      <div class="has-text-right">
        <button
          :class="{ 'is-loading': busy }"
          class="button is-primary"
          @click.prevent="submit">
          Save
        </button>
      </div>
    </form>
  </div>
</template>

<script>
import { mapMutations } from 'vuex'
import http from '@/modules/http'
import Message from '@/components/misc/Message.vue'

import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/user-cog'
import 'vue-awesome/icons/address-card'
import 'vue-awesome/icons/cogs'
import 'vue-awesome/icons/envelope'
import 'vue-awesome/icons/caret-up'
import 'vue-awesome/icons/caret-down'
import 'vue-awesome/icons/university'

export default {
  components: {
    Message,
    Icon,
  },

  data () {
    return {
      user: {
        settings: {},
      },
      busy: false,
      numberSettings: true,
      companySettings: true,
      contactSettings: true,
      userSettings: true,
      bankSettings: true,
    }
  },

  created () {
    this.load()
  },

  methods: {
    ...mapMutations([ 'setMessage', 'clearMessage', 'setUser' ]),

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
      try {
        this.clearMessage()
        this.busy = true

        const user = await http.putUser(this.user)
        this.setUser(user)
        this.setMessage({
          text: 'Settings saved.',
        })

        this.userSettings = false
        this.numberSettings = false
        this.companySettings = false
        this.contactSettings = false
        this.bankSettings = false
      } catch (error) {
        this.setMessage({
          text: error.message,
          style: 'is-danger',
        })
      } finally {
        this.busy = false
        this.$scrollTo('#top')
      }
    },
  },

  beforeRouteLeave (to, from, next) {
    this.clearMessage()
    next()
  },
}
</script>

<style lang="scss">
  .fieldset {
    margin-bottom: 60px;
    padding: 20px;
    border-radius: 2px;
    background: #fdfdfd;
    border: 1px solid #e9e9e9;
    border-radius: 2px;
  }

  .is-bordered {
    border-bottom: 1px solid #f1f1f1;
    margin-bottom: 20px;
  }

  .is-pointer {
    cursor: pointer;
  }
</style>
