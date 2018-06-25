import axios from 'axios';

const requestFilter = {
  requests: [],
  push(opts){
    let req = {
      time: new Date(),
      key: JSON.stringify(opts),
      opts,
    }
    this.requests.push(req);
  },
  pop(opts){
    let key = JSON.stringify(opts);
    let index = this.requests.findIndex(it=> it.key == key);
    if(index >= 0 ){
      this.requests.splice(index, 1);
    }
  },
  hasPending(opts){
    let key = JSON.stringify(opts);
    return this.requests.findIndex(it=> it.key == key) >= 0;
  }
};

export default {
  post (opts) {
    if(requestFilter.hasPending(opts)){ 
      console.log('repeating ajax url:', opts.url);
      return;
    }
    requestFilter.push(opts);
    return axios.post(opts.url, JSON.stringify(opts.data) || opts.body || {},
      {
        timeout: 1000*60,
        headers: {'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'}
      }).then((res) => {
        requestFilter.pop(opts);
        return res.data;
    }, (err) => {
      requestFilter.pop(opts);
      return opts.failed ? opts.failed(err) : 
        console.log(err);
      }).catch(e=>{
        requestFilter.pop(opts);
        console.log('Error happened:'+e)
    })
  },
  get (opts) {
    if(requestFilter.hasPending(opts)){ 
      console.log('repeating ajax url:', opts.url);
      return;
    }
    requestFilter.push(opts);
    return axios.get(opts.url, {
      params: opts.data || opts.params || {},
      timeout: 1000*60,
    }).then((res) => {
      requestFilter.pop(opts);
      if(res.code == 1) {
        return res.data;
      }else{
        this.$toasted.show(res.message, { 
          theme: "primary", 
          position: "top-center", 
          duration : 2000
        });
      }
    }, (err) => {
       requestFilter.pop(opts);
        if (err.response.status == 401) {
          return;
        }
      return opts.failed ? opts.failed(err) : console.log(err);
    });
  }
};
