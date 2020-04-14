<template>
  <aside class="menu">
    <ul
      v-for="(item, index) in items"
      :key="index"
      class="menu-list is-hidden-mobile"
    >
      <li>
        <a
          href
          :class="['is-clearfix', { 'is-active': active === item.view }]"
          @click.prevent="goTo(item.view)"
        >
          <span class="is-pulled-left">
            {{ item.title }}
          </span>
          <span
            v-if="errorCount(item.view) > 0"
            class="tag is-pulled-right is-danger"
          >
            {{ errorCount(item.view) }}
          </span>
        </a>
      </li>
    </ul>

    <div :class="['select is-hidden-tablet is-fullwidth', { 'is-danger': hasErrors }]">
      <select
        id="settings_menu"
        ref="settings_menu"
        v-model="selected"
      >
        <option
          v-for="(item, index) in items"
          :key="index"
          :value="item.view"
        >
          {{ item.title }}
          <span v-if="errorCount(item.view)">
            (!)
          </span>
        </option>
      </select>
    </div>
  </aside>
</template>

<script>
export default {
  props: {
    items: {
      type: Array,
      required: true,
    },
    errors: {
      type: Array,
      default: () => [],
    },
    active: {
      type: String,
      default: '',
    },
  },

  data () {
    return {
      selected: '',
    }
  },

  computed: {
    hasErrors () {
      return this.errors.filter(e => e.count > 0).length
    }
  },

  watch: {
    selected () {
      this.goTo(this.selected)
      this.$refs.settings_menu.blur()
    }
  },

  created () {
    this.selected = this.active
  },

  methods: {
    hasError (view) {
      return this.errors.some(e => e.view === view && e.errors)
    },

    errorCount (view) {
      const errors = this.errors.find(e => e.view === view)
      if (errors) {
        return errors.count
      }

      return 0
    },

    goTo (view) {
      if (view !== this.active) {
        this.$router.push(`/settings/${view}`)
      }
    }
  },
}
</script>
