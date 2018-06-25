import moment from 'moment';
import numeral from 'numeral';
import store from '../store';
import _ from 'lodash';

export default {
  isMobile(){
    return navigator.userAgent.match(/mobile/i);
  },
  findSelectText(val, options){
    let op = options.find(it=>it.value.toLowerCase() == val.toLowerCase());
    return op ? op.label : '';
  },
  getIpCountOfCIDR(str){
    let rs = /\d+\.\d+\.\d+\.\d+\/(\d+)/.exec(str);
    if(rs && rs.length >= 2){
      let mask = rs[1];
      let count = Math.pow(2,32-mask) - 4;
      return count;
    }
    return 0;
  },
  getRegionName(regionId){
    let regions = store.getters["getters/regions"];
    let region = regions && regions.find(it=>it.id == regionId)
    return region ? region.name : '';
  },
  getToday(){
    return moment().subtract('0', 'days').format('YYYY-MM-DD');
  },
  getYesterday(){
    return moment().subtract('1', 'days').format('YYYY-MM-DD');
  },
  get7DaysAgo(){
    return moment().subtract('7', 'days').format('YYYY-MM-DD');
  },
  getNDaysAgo(n){
    return moment().subtract(n, 'days').format('YYYY-MM-DD');
  },
  formatMonth(m){
    let time = moment(m).format('YYYY-MM');
    if(time == 'Invalid date'){
      return '-';
    }
    return time;
  },
  formatDate(m){
    let time = moment(m).format('YYYY-MM-DD');
    if(time == 'Invalid date'){
      return '-';
    }
    return time;
  },
  formatDateTime(m){
    let time = moment(m).format('YYYY-MM-DD HH:mm:ss');
    if(time == 'Invalid date'){
      return '-';
    }
    return time;
  },
  formatMoney(m){
    return numeral(m).format('0.00');
  },
  //13800001234 => 138xxxx1234
  formatMobile(m){
    if(!m || m.length != 11){
      return m;
    }
    return m.substr(0, 3) + '****' + m.substr(7,4);
  },
  isInt(num){
    if (!isNaN(num)){
      num = parseFloat(num, 10);
      return parseInt(num, 10) === num;
    } 
    return false;
  },
  isFloat(num){
    return !isNaN(num) && num.toString().indexOf('.') >= 0;
  },
  getURLParam(name){
    var rs = '';
    var query = location.href.substr(location.href.indexOf('?')+1);
    if(query){
      var params = query.split('&');
      params.find(it=>{
        var param = it.split('=');
        if(param.length == 2 && param[0] == name){
          rs = param[1];
          return true;
        }
      });
    }
    return rs;
  },
};
