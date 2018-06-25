<template lang="pug">
div
	.beta
		h1 暂不开放
		p 很抱歉我们暂时不开放这个页面的浏览
		md-button.md-raised.md-primary(@click="$router.push('/login')") 返回登录页面
</template>
<script lang="babel">
  import * as types from '../../constants/vuex-types';
  import { mapState, mapGetters } from 'vuex';
  export default {
    computed: {
      ...mapGetters({
      })
    },
  }
</script>
<style lang="less" src="./index.less"></style>
