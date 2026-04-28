export default {
  async fetch(request, env, ctx) {
    const url = new URL(request.url);
    const { pathname, search } = url;

    const proxyMap = {
      '/volces/':     'https://ark.cn-beijing.volces.com',
      '/dashscope/':  'https://coding.dashscope.aliyuncs.com',
      '/deepseek/':   'https://api.deepseek.com',
      '/openrouter/': 'https://openrouter.ai',
      '/opencode/':   'https://opencode.ai',
    };

    let targetUrl = '';
    for (const [prefix, base] of Object.entries(proxyMap)) {
      if (pathname.startsWith(prefix)) {
        const restPath = pathname.slice(prefix.length-1);
        targetUrl = base + restPath + search;
        break;
      }
    }

    if (!targetUrl) {
      return new Response('Not Found', { status: 404 });
    }

    const proxyRequest = new Request(targetUrl, {
      method: request.method,
      headers: request.headers,
      body: request.body,
    });

    return fetch(proxyRequest, { redirect: 'follow' });
  }
};

// export default {
//   async fetch(request, env, ctx) {
//     const url = new URL(request.url);

//     let newUrl = '';
//     if (url.pathname.startsWith('/volces/')) {
//       // https://ark.cn-beijing.volces.com/api/coding/v3
//       // https://ark.cn-beijing.volces.com/api/coding
//       newUrl = 'https://ark.cn-beijing.volces.com' + url.pathname.substring(7) + url.search;
//     }
//     if (url.pathname.startsWith('/dashscope/')) {
//       // https://coding.dashscope.aliyuncs.com/v1
//       // https://coding.dashscope.aliyuncs.com/apps/anthropic
//       newUrl = 'https://coding.dashscope.aliyuncs.com' + url.pathname.substring(10) + url.search;
//     }
//     if (url.pathname.startsWith('/deepseek/')) {
//       // https://api.deepseek.com
//       // https://api.deepseek.com/anthropic      
//       newUrl = 'https://api.deepseek.com' + url.pathname.substring(9) + url.search;
//     }
//     if (url.pathname.startsWith('/openrouter/')) {
//       // https://openrouter.ai/api/v1      
//       newUrl = 'https://openrouter.ai' + url.pathname.substring(11) + url.search;
//     }
//     if (url.pathname.startsWith('/opencode/')) {
//       // https://opencode.ai/zen/go/v1
//       newUrl = 'https://opencode.ai' + url.pathname.substring(9) + url.search;
//     }

//     if (newUrl !== '') {
//       // const newUrl = prefix + url.pathname + url.search;
//       const proxyRequest = new Request(newUrl, {
//         method: request.method,
//         headers: request.headers,
//         body: request.body,
//       });
//       return fetch(proxyRequest);
//     }
//   }
// };
