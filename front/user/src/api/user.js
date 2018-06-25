import request from '../utils/request';
import URLS from '../constants/URLS';

export default {
  login (query) {
    return request.post({
      url: URLS.USER_LOGIN,
      data: query
    });
  },
  create (query) {
    return request.post({
      url: URLS.USER_CREATE,
      data: query
    });
  },
  reset (query) {
    return request.post({
      url: URLS.USER_RESET,
      data: query
    });
  },
  getMessageCode (query) {
    return request.post({
      url: URLS.USER_GET_MESSAGECODE,
      data: query
    });
  },
};
