<template lang="pug">
div
	.create
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
				md-autocomplete(type='email', v-model='email', @md-changed='emailChanged', :md-options='[]', md-layout='box', md-dense)
					label 请输入您常用的电子邮箱账号
			.md-input-horizon
				md-autocomplete(v-model='phone', @md-changed='phoneChanged', :md-options='[]', md-layout='box', md-dense)
					label 请输入您常用的手机号
					md-button.md-primary(@click='getMessage', :disabled='messageButtonDisabled') {{getMessageButton}}
			.md-input-horizon
				md-autocomplete(type='number', v-model='messageCode', @md-changed='codeChanged', :md-options='[]', :max='6', md-layout='box', md-dense)
					label 请输入收到的短信验证码
			.md-button-horizon
				md-button.md-raised.md-primary(@click='create') 立即创建
				md-button.md-raised.md-block(@click="$router.push('/login?redirect=')") 已有账号快速登录
			.md-button-horizon
				p(align='center')
					router-link(to='/forgot') 登录遇到问题？找回密码
		.bottom
			p(align='center') 使用以下社交账号登录或绑定
			.iconGroups
				md-button.md-icon-button.md-raised(@click='github')
					md-icon.md-size-1x(md-src='/static/user/images/app/github.svg')
	// agree modal
	md-dialog-confirm(:md-active.sync='activeAgree', md-title='如您确认代表您同意协议', md-content="如您同意注册，代表您已经同意<a href='http://www.golang.pw/public/rules' target='_blank'>《用户使用条约》</a>. <br> 这意味着 <strong>您的账户收到保护与法律约束</strong> 如您不同意，则不能继续创建账户.", md-confirm-text='同意并继续', md-cancel-text='放弃', @md-confirm='onConfirm')
	md-dialog.captchaProptDialogue(:md-active.sync='proptCaptcha')
		md-dialog-title 请输入图片验证码
		md-tabs(md-dynamic-height)
			md-tab(md-label='验证码请求安全校验')
				.md-input-horizon
					md-autocomplete(v-model='captcha', @md-changed='captchaChanged', :md-options='[]', md-layout='box', md-dense)
						label 请输入图形验证码
						img.captcha(v-if='captcha_src', @click='refreshCaptcha', :src='captcha_src')
		md-dialog-actions
			md-button.md-primary(@click='proptCaptcha = false') 取消
			md-button.md-primary(@click='getCode') 立即发送短信验证码
	// agree modal end
	md-dialog-alert(:md-active.sync='registySuccess', md-title='注册成功，前往登录吧', md-content='现在您可以点击立即登录去登录您的账号了')

</template>
<script lang="babel">
  import * as types from '../../../constants/vuex-types';
  import { mapState, mapGetters } from 'vuex';
  import api from '../../../api/user';
  
  export default {
    name: 'create',
    computed: {
      ...mapGetters({
      }),
    },
    data(){
        return{
          username: '',
          password: '',
          email: '',
          phone: '',
          captcha: '',
          captcha_src: '',
          value: '',
          messageCode: '',
          getMessageButton: '获取验证码',
          position: 'center',
          isInfinity: false,
          activeAgree: false,
          messageButtonDisabled: false,
          proptCaptcha: false,
          registySuccess: false,
          timer: null,
          count: 0,
          duration: 3000,
        }
    },
    mounted(){
      this.$nextTick(async function () {
          this.refreshCaptcha();
      });
    },
    methods: {
      async onConfirm (){
        let {username, phone, password, email, messageCode} = this;
        let query = {
          username,
          messageCode,
          phone,
          password,
          email          
        }
        let rs = await api.create(query);
        if(rs){
          if(rs.code == "0"){
            this.$toasted.show(rs.message, { 
              theme: "primary", 
              position: "bottom-center", 
              duration : 2000
            });
          }else{
            /*注册成功，提示登录*/
            this.registySuccess = true
          }
        }
      },
      async create(){
        let {username, phone, password, email, messageCode} = this;
        if(!username || !phone || !password || !email || !messageCode){
          this.$toasted.show("请您完整输入注册所需参数在进行创建", { 
            theme: "primary", 
            position: "bottom-center", 
            duration : 2000
          });
          return;
        }
        this.activeAgree = true;
      },
      async getMessage(){
        let {messageButtonDisabled, phone} = this;
        if(messageButtonDisabled){
          return;
        }
        if(!phone){
          this.$toasted.show("请输入您的手机号码", { 
            theme: "primary", 
            position: "bottom-center", 
            duration : 2000
          });
          return;
        }
        this.proptCaptcha = true;
      },
      async getCode(){
        //获得短信验证码
        let {captcha, phone} = this;
        if(!captcha){
           this.$toasted.show("若要取得短信验证码，请输入图片验证码", { 
            theme: "primary", 
            position: "bottom-center", 
            duration : 2000
          });
          return;
        }
        const TIME_COUNT = 60;
        if (!this.timer) {
          this.count = TIME_COUNT;
          this.messageButtonDisabled = true;
          let query = {
            captcha,
            phone
          }
          let rs = await api.getMessageCode(query);
          if(rs){
            if(rs.code == "0"){
              this.$toasted.show(rs.message, { 
                theme: "primary", 
                position: "bottom-center", 
                duration : 2000
              });
              return;
            }else{
              this.proptCaptcha = false;
            }
          }
          this.timer = setInterval(() => {
            if (this.count > 0 && this.count <= TIME_COUNT) {
              this.count--;
              this.getMessageButton = `等待${this.count}秒`;
            } else {
              this.getMessageButton = '重新获取';
              this.messageButtonDisabled = false;
              clearInterval(this.timer);
              this.timer = null;
            }
          }, 1000)
        }
      },
      phoneChanged(val){
        this.phone = val;
      },
      captchaChanged(val){
        this.captcha = val;
      },
      passwordChanged(val){
        this.password = val;
      },
      emailChanged(val){
        this.email = val;
      },
      usernameChanged(val){
        this.username = val;
      },
      codeChanged(val){
        this.messageCode = val;
      },
      github(){
        location.href = "/api/v1/auth/github/login";
      },
      refreshCaptcha(){
        let ver = new Date().getTime();
        this.captcha_src = `/api/v1/user/captcha/?ver=${ver}`;
      }
    }
  }
</script>
<style lang="less" src="./index.less"></style>
