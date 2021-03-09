const pureHttp = require('pure-http');
const cors = require('cors');
const rateLimit = require('express-rate-limit');
const v1Route = require('./v1');

const app = pureHttp();
const limiter = rateLimit({
  windowMs: 15 * 60 * 1000, // 15 minutes
  max: 100, // limit each IP to 100 requests per windowMs
});
app.use(cors());
app.use(limiter());

app.use('/v1', v1Route);

app.listen(3000);
