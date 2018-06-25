
import * as types from '../../constants/vuex-types';
import api from '../../api/user';


const state = {
  user: null,
  hasLogin: false
};

const getters = {
  [types.GETTERS_USER] : state => state.user,
  [types.GETTERS_LOGIN] : state => state.hasLogin
}

const mutations = {
  [types.MUTATION_USER] (state, user) {
    state.user = user;
  },
  [types.MUTATION_LOGIN] (state, hasLogin) {
    state.hasLogin = hasLogin;
  }
}

const actions = {
  resetCurrentUser({ commit }){
    commit(types.MUTATION_USER, {}  );
    commit(types.MUTATION_LOGIN, false);
  },
  async loadCurrentUser ({ commit }) {
    var rs = await api.getUser(f=>f);
    if(rs && rs.data && rs.data.length > 0){
      let user = rs.data[0];
      commit(types.MUTATION_USER, user);
      commit(types.MUTATION_LOGIN, true);
    }
  }
}

export default {
  state,
  getters,
  actions,
  mutations
};
