import { createApp } from 'vue'
import App from './App.vue'
import PrimeVue from 'primevue/config'
import 'primevue/resources/themes/aura-light-green/theme.css'
import CustomButton from "primevue/button"
import "/node_modules/primeflex/primeflex.css"
import "primeflex/themes/primeone-light.css"

const app = createApp(App)
app.use(PrimeVue)
app.component('CustomButton', CustomButton)
app.mount('#app')
