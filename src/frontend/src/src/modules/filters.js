import Vue from 'vue'
import dayjs from 'dayjs'

const locale = process.env.VUE_APP_LOCALE || 'en-US'

Vue.filter('money', (val, cur = 'EUR') => {
  return val.toLocaleString(locale, {
    style: 'currency',
    currency: cur,
  })
})

Vue.filter('date', val => {
  const d = dayjs(val)

  if (locale === 'de-DE') {
    return d.format('DD.MM.YYYY')
  }

  return d.format('MM/DD/YYYY')
})
