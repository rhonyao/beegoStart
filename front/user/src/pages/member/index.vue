<template lang="pug">
  md-app
    md-app-content
      router-view
</template>

<script lang="babel">
  import * as types from '../../constants/vuex-types';
  import { mapState, mapGetters } from 'vuex';
  export default {
    name: 'PersistentMini',
  }
</script>

<style lang="less" src="./index.less"></style>

