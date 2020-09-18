import Vue from 'vue';
import Vuetify from 'vuetify/lib';

Vue.use(Vuetify);

const theme = {
  primary: "#3DDB85",
  primaryDark: "0324a6",
  secondary: "#0033e6",
  textColor: "#1c2055",
  textColorLight: "#A8B0C9",
  accent: "#4698a4",
  error: "#FF5252",
  info: "#2196F3",
  success: "#4CAF50",
  warning: "#FFC107"
};

export default new Vuetify({
  theme: {
    themes: {
      light: theme
    }
  }
});