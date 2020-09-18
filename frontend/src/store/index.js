import Vue from "vue";
import Vuex from "vuex";
import http from "@/plugins/axios";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    loading: false,
    snackbar: false,
    errMessage: '',
    files: [],
  },
  getters: {
    sortFiles(state) {
      let prop = "upload_time";
      if (state.files) {
        return state.files.sort((a, b) => (a[prop] < b[prop] ? 1 : -1));
      } else {
        return null;
      }
    },
  },
  mutations: {
    loading(state) {
      state.loading = true;
    },
    loaded(state) {
      state.loading = false;
    },
    fetchFileList(state, result) {
      state.files = result;
    },
  },
  actions: {
    async fetchFileList({ commit }) {
      await http()
        .get("/apks")
        .then((response) => {
          commit("fetchFileList", response.data.data.items);
        });
    },
    async addFile({ commit }, { file, name, description }) {
      let params = new FormData();
      params.append("upload_by", name);
      params.append("apk_description", description);
      params.append("file", file);
      commit("loading");
      await http()
        .post("/apks", params)
        .then((response) => {
          if (response.data.code === 0) {
            commit("loaded");
            this.dispatch("fetchFileList");
          }
        })
        .catch((error) => {
          console.log(error);
          commit("loaded");
        });
    },
    async deleteFile(context, { id }) {
      const url = "/apks?id=" + id;
      await http()
        .delete(url)
        .then((response) => {
          if (response.data.code === 0) this.dispatch("fetchFileList");
        }).catch(error => {
          console.log(error)
        });
    },
  },
});
