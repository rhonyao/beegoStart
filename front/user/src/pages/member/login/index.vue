<template lang="pug">
div
	.login
		.panel
			.logo
				img(src='/static/user/images/logo.png')
			.md-input-horizon
				md-autocomplete(v-model='username', @md-changed='usernameChanged', :md-options='[]', md-layout='box', md-dense)
					label 请输入您的用户名
			.md-input-horizon
				md-autocomplete(type='password', v-model='password', @md-changed='passwordChanged', :md-options='[]', md-layout='box', md-dense)
					label 请输入您的密码
			.md-input-horizon
				md-autocomplete(v-model='captcha', @md-changed='captchaChanged', :md-options='[]', md-layout='box', md-dense)
					label 请输入图形验证码
					img.captcha(v-if='captcha_src', @click='refreshCaptcha', :src='captcha_src')
			.md-button-horizon
				md-button.md-raised.md-primary(@click='login') 立即登录
				md-button.md-raised.md-block(@click="$router.push('/create')") 注册账号
			.md-button-horizon
				p(align='center')
					router-link(to='/forgot') 登录遇到问题？找回密码
		.bottom
			p(align='center') 使用以下社交账号登录或绑定
			.iconGroups
				md-button.md-icon-button.md-raised(@click='github')
					md-icon.md-size-1x(md-src='/static/user/images/app/github.svg')
</template>
<script lang="babel">
  import * as types from '../../../constants/vuex-types';
  import { mapState, mapGetters } from 'vuex';
  import api from '../../../api/user';
  export default {
    name: 'login',
    computed: {
      ...mapGetters({
      }),
    },
    data(){
        return{
          redirect: this.$route.query.redirect,
          username: '',
          password: '',
          captcha: '',
          captcha_src: ''
        }
    },
    mounted(){
      this.$nextTick(async function () {
          this.refreshCaptcha();
          console.log("跳转Url:",decodeURI(this.redirect));
      });
    },
    methods: {
      async login(){
        let {username, password, captcha, redirect} = this;
        let query = {
          username,
          password,
          captcha
        }
        if(!username || !password || !captcha){
          this.$toasted.show("请完善登录信息后再进行登录", { 
            theme: "primary", 
            position: "bottom-center", 
            duration : 2000
          });
          return;
        }
        let rs = await api.login(query);
        if(rs){
          if(rs.code == "0"){
            this.$toasted.show(rs.message, { 
              theme: "primary", 
              position: "bottom-center", 
              duration : 2000
            });
          }else{
            /*登录成功，返回回调页面*/
            if(redirect == ""){
              location.href = `/`;
            }else{
              location.href = decodeURI(redirect);
            }
          }
        }
      },
      github(){
        location.href = "/api/v1/auth/github/login";
      },
      refreshCaptcha(){
        let ver = new Date().getTime();
        this.captcha_src = `/api/v1/user/captcha/?ver=${ver}`;
      },
      usernameChanged(val){
        this.username = val;
      },
      captchaChanged(val){
        this.captcha = val;
      },
      passwordChanged(val){
        this.password = val;
      },
    },
    watch: {
    }
  }
</script>
<style lang="less" src="./index.less"></style>
