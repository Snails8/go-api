const proxy = require('http-proxy-middleware');

module.exports = function(app) {
//   app.use(proxy("/api/auth", { target: "http://localhost:8080/" }));
    app.use(
        proxy("/api/users", 
        { 
            target: "http://localhost:8081/" 
        }
    ));
};