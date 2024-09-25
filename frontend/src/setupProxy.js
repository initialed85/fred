const { createProxyMiddleware } = require("http-proxy-middleware");

module.exports = function (app) {
  app.use(
    "/api",
    createProxyMiddleware({
      target:
        process.env.REMOTE === "1"
          ? "https://fred.initialed85.cc/api"
          : "http://localhost:7070/api",
      changeOrigin: true,
    }),
  );
};
