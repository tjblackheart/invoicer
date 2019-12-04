import Vue from 'vue'
import dayjs from 'dayjs'

const locale = process.env.VUE_APP_LOCALE || 'de-DE'

Vue.filter('money', (val, cur = 'EUR') => {
  return (val).toLocaleString(locale, {
    style: 'currency',
    currency: cur,
  })
})

Vue.filter('date', val => {
  const d = dayjs(val)
  return (locale === 'de-DE') ? d.format('DD.MM.YYYY') : d.format('MM/DD/YYYY')
})
