import Vue from "vue";
import App from "./App.vue";
import {
  FormInputPlugin,
  NavbarPlugin,
  LayoutPlugin,
  IconsPlugin,
  BCard,
  BInputGroup,
  BButton,
} from "bootstrap-vue";

import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";

Vue.config.productionTip = false;

Vue.use(FormInputPlugin);
Vue.use(NavbarPlugin);
Vue.use(LayoutPlugin);
Vue.component("b-card", BCard);
Vue.component("b-input-group", BInputGroup);
Vue.component("b-button", BButton);
Vue.use(IconsPlugin);

new Vue({
  render: (h) => h(App),
}).$mount("#app");
