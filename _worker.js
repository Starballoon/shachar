// const TARGET = "https://coding.dashscope.aliyuncs.com";
const TARGET = 'https://ark.cn-beijing.volces.com/api/coding/v3';

// export default {
//   async fetch(request, env, ctx) {
//     const url = new URL(request.url);
//     const newUrl = TARGET + url.pathname + url.search;
//     return fetch(newUrl, request);
//   }
// }

export default {
  async fetch(request, env, ctx) {
    const url = new URL(request.url);
    const newUrl = TARGET + url.pathname + url.search;
    const proxyRequest = new Request(newUrl, {
      method: request.method,
      headers: request.headers,
      body: request.body,
    });
    return fetch(proxyRequest);
  }
};
