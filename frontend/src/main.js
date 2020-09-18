import Vue from "vue";
import App from "./App.vue";
import vuetify from "./plugins/vuetify";
import axios from "axios";
import uploader from "vue-simple-uploader";
import store from './store'
import VueQrcode from '@chenfengyuan/vue-qrcode'

Vue.config.productionTip = false;
Vue.prototype.$axios = axios;
Vue.component(VueQrcode.name, VueQrcode);

Vue.use(uploader);

new Vue({
  vuetify,
  store,
  render: (h) => h(App)
}).$mount("#app");
