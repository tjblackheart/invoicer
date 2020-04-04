<template>
  <div class="content">
    <b-input
      type="text"
      id="search"
      v-model="filterValue"
      :placeholder="placeholder"
      @escape="filterValue = ''"
    />

    <div class="is-clearfix">
      <p class="is-pulled-left">
        <b-checkbox
          v-if="type === 'invoice'"
          v-model="showCancelled"
          label="Show cancelled invoices"
        />
      </p>
      <p
        v-if="hasFilters"
        class="is-pulled-right"
      >
        <a href @click.prevent="resetFilters"> Reset all filters </a>
      </p>
    </div>
  </div>
</template>

<script>
import BInput from '@/components/fields/Input'
import BCheckbox from '@/components/fields/Checkbox'

export default {
  props: {
    placeholder: {
      type: String,
      required: true,
    },
    type: {
      type: String,
      required: true,
    }
  },

  components: {
    BInput,
    BCheckbox,
  },

  computed: {
    showCancelled: {
      get () {
        return this.$store.getters.showCancelled
      },
      set (val) {
        this.$store.commit('showCancelled', val)
      }
    },

    filterValue: {
      get () {
        return this.$store.getters.filterValue
      },
      set (val) {
        this.$store.commit('filterValue', val)
      }
    },

    filters () {
      return this.$store.getters.filters
    },

    hasFilters () {
      return this.showCancelled === false || this.filters.length
    }
  },

  methods: {
    resetFilters () {
      this.showCancelled = true
      this.filterValue = ''
    }
  }
}
</script>

<style>

</style>
