import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    UserDetails: {
      Login_id: 0,
      Role: "",
      User_id: "",
      History_id: "",
    },
    // BillStockDetails:[],
  },
  mutations: {
    setUserValue(state, value) {

      state.UserDetails.Login_id = value.Login_id;
      state.UserDetails.Role = value.Role;
      state.UserDetails.User_id = value.User_id;
      state.UserDetails.History_id = value.History_id;
    },
    SetEmptyUserDetails(state) {
      (state.UserDetails.Login_id = ""),
        (state.UserDetails.Role = ""),
        (state.UserDetails.User_id = ""),
        (state.UserDetails.History_id = "");
    },
  },
  actions: {},
  modules: {},
});
