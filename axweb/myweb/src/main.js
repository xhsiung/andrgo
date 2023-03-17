import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/main.css'
import '@fortawesome/fontawesome-free/css/all.min.css'
import '@fortawesome/fontawesome-free/js/all.min.js'

//custom js 
import "./init.js"

import { createI18n } from 'vue-i18n'
import tw from '@/lang/tw.json'
import en from '@/lang/en.json'
let xi18n = createI18n({
    locale: 'tw',
    messages: {
        'en': en,
        'tw': tw,
    }
})

const app = createApp(App)

app.use(router)
app.use(xi18n)

app.mount('#app')

