const TARGET = "https://coding.dashscope.aliyuncs.com";

export default {
  async fetch(request, env, ctx) {
    const url = new URL(request.url);
    const newUrl = TARGET + url.pathname + url.search;
    return fetch(newUrl, request);
  }
}
