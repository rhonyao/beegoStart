import Vue from 'vue';
import Vuex from 'vuex';
import * as types from '../constants/vuex-types';
import user from './modules/user';

Vue.use(Vuex);

const state = {
  priceCat: 'compute'
};

const getters = {
  [types.GETTERS_PRICE_CAT] : state => state.priceCat
}

const mutations = {
  [types.MUTATION_PROVIDER_ID] (state, providerId) {
    state.providerId = providerId;
  }
}

const actions = {
  updateProviderId ({ commit }, { providerId }) {
    commit(types.MUTATION_PROVIDER_ID, providerId);
  }
}

export default new Vuex.Store({
  state,
  getters,
  actions,
  mutations,
  modules: {
    user
  }
});
